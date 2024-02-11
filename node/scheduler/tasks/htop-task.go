package tasks

import (
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"go.uber.org/zap"
	"surena/node/utils"
	"time"
)

type HtopTask struct {
	cron   *cron.Cron
	id     cron.EntryID
	logger *zap.SugaredLogger
	Swap   *mem.SwapMemoryStat
	RAM    *mem.VirtualMemoryStat
	CPU    float64
}

func NewHtopTask(cron *cron.Cron) (*HtopTask, error) {
	logger, err := utils.NewLogger("htop-task")
	if err != nil {
		logger.Warn("Failed to create htop task")
		return nil, err
	}

	logger.Info("Creating htop task")
	task := &HtopTask{
		cron:   cron,
		logger: logger,
	}

	err = task.Start()
	if err != nil {
		logger.Warn("Failed to start htop task")
		return nil, err
	}

	task.Run()
	return task, nil
}

func (t *HtopTask) Run() {
	t.logger.Debug("Running htop task")

	t.checkSwap()
	t.checkRAM()
	t.checkCPU()
}

// Start Add the task to the scheduler
// to be executed
func (t *HtopTask) Start() error {
	if t.id != 0 {
		t.logger.Warn("Htop task is already running")
		return nil
	}

	t.logger.Info("Starting htop task")
	taskId, err := t.cron.AddJob("@every 5s", t)
	if err != nil {
		t.logger.Warn("Failed to add htop task to scheduler")
		return err
	}

	t.id = taskId
	return nil
}

// Stop Remove the task from the scheduler
// to stop it from being executed
func (t *HtopTask) Stop() {
	if t.id != 0 {
		t.logger.Info("Stopping htop task")
		t.cron.Remove(t.id)
		t.id = 0
	}
}

func (t *HtopTask) checkSwap() {
	swap, err := mem.SwapMemory()
	if err != nil {
		t.logger.Warn("Failed to get swap memory")
		return
	}

	t.Swap = swap
	t.logger.Debug("Swap usage: ", t.Swap.UsedPercent)
}

func (t *HtopTask) checkRAM() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		t.logger.Warn("Failed to get virtual memory")
		return
	}

	t.RAM = memory
	t.logger.Debug("RAM usage: ", t.RAM.UsedPercent)
}

func (t *HtopTask) checkCPU() {
	percent, err := cpu.Percent(time.Second*1, false)
	if err != nil {
		t.logger.Warn("Failed to get CPU usage")
		return
	}

	t.CPU = percent[0]
	t.logger.Debug("CPU usage: ", t.CPU)
}
