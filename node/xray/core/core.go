package core

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/errors"
	"os/exec"
	"path/filepath"
	"surena/node/env"
	"surena/node/utils"
	"sync"
)

var core *Core
var logger = utils.CreateLogger("xray").WithField("module", "core")

type Core struct {
	CoreInterface
	sync.Mutex
	Logger             *logrus.Entry
	Version            string
	ConfigPath         string
	ExecutablePath     string
	Started            bool
	Cmd                *exec.Cmd
	LogListenersLastID uint
	LogListeners       map[uint]func(log string)
}

type CoreInterface interface {
	IsRunning() bool
	AddLogListener(listener func(log string)) uint
	RemoveLogListener(id uint)
	Start()
	Stop()
}

func init() {
	logger.Debug("initializing xray core")

	core = &Core{
		Logger:         logger,
		Version:        env.GetXrayVersion(),
		ConfigPath:     env.GetXrayConfigPath(),
		ExecutablePath: env.GetXrayExecutablePath(),
		Started:        false,
		LogListeners:   make(map[uint]func(log string)),
	}
}

func Initialize() (CoreInterface, error) {
	if core == nil {
		return nil, errors.New("xray core not initialized")
	}

	return core, nil
}

func Get() CoreInterface {
	if core == nil {
		panic("xray core not initialized")
	}

	return core
}

func (c *Core) IsRunning() bool {
	return c.Started
}

func (c *Core) Start() {
	c.Lock()
	defer c.Unlock()

	logger.Info("starting xray core")

	if c.IsRunning() {
		logger.Warn("xray core is already running")
		return
	}

	configDir := filepath.Dir(c.ConfigPath)
	logger.Info(c.ExecutablePath, " -c ./", c.ConfigPath, " -confdir ", configDir)
	cmd := exec.Command(c.ExecutablePath, "-c", c.ConfigPath, "-confdir", configDir)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.Logger.Error("error creating StdoutPipe for xray:", err)
		return
	}

	if err = cmd.Start(); err != nil {
		c.Logger.Error("error starting xray:", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		scanner := bufio.NewScanner(stdout)
		isReady := false

		for scanner.Scan() {
			line := scanner.Text()
			c.Logger.Trace(line)
			for _, listener := range c.LogListeners {
				listener(line)
			}

			// xray core is ready
			if !isReady && len(line) > 5 && line[:5] == "Xray " {
				isReady = true
				wg.Done()
			}
		}

		if err = scanner.Err(); err != nil {
			c.Logger.Error("error reading from xray pipe:", err)
			return
		}

		c.Cmd = cmd
		c.Started = true
		if err = cmd.Wait(); err != nil {
			c.Cmd = nil
			c.Started = false
			c.Logger.Error("error waiting for xray:", err)
		}
	}()

	wg.Wait()
	c.Logger.Info("xray core started")
}

func (c *Core) Stop() {
	c.Lock()
	c.Unlock()

	if !c.IsRunning() {
		c.Logger.Warn("xray core is already stopped")
		return
	}

	if err := c.Cmd.Process.Kill(); err != nil {
		c.Logger.Error("error stopping xray:", err)
		return
	}

	c.Cmd = nil
	c.Started = false
	c.Logger.Info("xray core stopped")
}

func (c *Core) AddLogListener(listener func(log string)) uint {
	c.Lock()
	c.Unlock()

	c.LogListenersLastID++
	c.LogListeners[c.LogListenersLastID] = listener
	return c.LogListenersLastID
}

func (c *Core) RemoveLogListener(id uint) {
	c.Lock()
	c.Unlock()
	delete(c.LogListeners, id)
}
