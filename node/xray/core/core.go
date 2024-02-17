package core

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os/exec"
	"surena/node/env"
	"surena/node/utils"
	"sync"
)

var core *Core

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
	logger := utils.CreateLogger("xray-core")
	logger.Debug("initializing xray core")

	core = &Core{
		Logger:         logger,
		Version:        env.GetXrayVersion(),
		ConfigPath:     env.GetXrayConfigPath(),
		ExecutablePath: env.GetXrayExecutablePath(),
		Started:        false,
		LogListeners:   make(map[uint]func(log string)),
	}

	core.Start()
}

func Get() CoreInterface {
	return core
}

func (c *Core) IsRunning() bool {
	return c.Started
}

func (c *Core) Start() {
	c.Lock()
	defer c.Unlock()

	if c.IsRunning() {
		c.Logger.Warn("xray core is already running")
		return
	}

	cmd := exec.Command(c.ExecutablePath, "-c", c.ConfigPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.Logger.Error("error creating StdoutPipe for xray:", err)
		return
	}

	if err = cmd.Start(); err != nil {
		c.Logger.Error("error starting xray:", err)
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			c.Logger.WithField("module", "Xray").Trace(line)
			for _, listener := range c.LogListeners {
				listener(line)
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
