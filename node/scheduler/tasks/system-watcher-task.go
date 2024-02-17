package tasks

import (
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/sirupsen/logrus"
	"surena/node/utils"
	"time"
)

type SystemWatcherTask struct {
	SystemWatcherTaskInterface
	Logger   *logrus.Entry
	Cron     *cron.Cron
	ID       cron.EntryID
	Schedule string
}

type SystemWatcherTaskInterface interface {
	IsRunning() bool
	Start()
	Stop()
}

func NewSystemWatcherTask(cron *cron.Cron) SystemWatcherTaskInterface {
	logger := utils.CreateLogger("system-watcher-task")
	logger.Info("Creating htop task")

	task := &SystemWatcherTask{
		Logger:   logger,
		Cron:     cron,
		Schedule: "@every 10s",
	}

	task.run()
	task.Start()
	return task
}

func (t *SystemWatcherTask) checkSwap() {
	swap, err := mem.SwapMemory()
	if err != nil {
		t.Logger.Warn("Failed to get swap memory")
		return
	}

	if swap.UsedPercent > 90 {
		t.Logger.Warn("Swap usage is too high: ", swap.UsedPercent)
	} else {
		t.Logger.Trace("Swap usage: ", swap.UsedPercent)
	}
}

func (t *SystemWatcherTask) checkRAM() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		t.Logger.Warn("Failed to get virtual memory")
		return
	}

	if memory.UsedPercent > 90 {
		t.Logger.Warn("RAM usage is too high: ", memory.UsedPercent)
	} else {
		t.Logger.Trace("RAM usage: ", memory.UsedPercent)
	}
}

func (t *SystemWatcherTask) checkCPU() {
	percent, err := cpu.Percent(time.Second*1, false)
	if err != nil {
		t.Logger.Warn("Failed to get CPU usage")
		return
	}

	if percent[0] > 90 {
		t.Logger.Warn("CPU usage is too high: ", percent[0])
	} else {
		t.Logger.Trace("CPU usage: ", percent[0])
	}
}

func (t *SystemWatcherTask) run() {
	t.Logger.Trace("Running htop task")

	t.checkSwap()
	t.checkRAM()
	t.checkCPU()
}

func (t *SystemWatcherTask) IsRunning() bool {
	return t.ID != 0
}

func (t *SystemWatcherTask) Start() {
	if t.ID != 0 {
		t.Logger.Warn("htop task is already running")
		return
	}

	var err error
	t.ID, err = t.Cron.AddFunc(t.Schedule, t.run)
	if err != nil {
		t.Logger.Warn("Failed to start htop task")
		return
	}

	t.Logger.Debug("Starting htop task")
}

func (t *SystemWatcherTask) Stop() {
	if t.ID == 0 {
		t.Logger.Warn("htop task is already stopped")
		return
	}

	t.Cron.Remove(t.ID)
	t.ID = 0
	t.Logger.Debug("Stopping htop task")
}
