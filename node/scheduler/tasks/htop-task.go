package tasks

import (
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

type HtopTask struct {
	cron.Job
	Swap *mem.SwapMemoryStat
	RAM  *mem.VirtualMemoryStat
	CPU  float64
}

func NewHtopTask() *HtopTask {
	task := &HtopTask{}

	task.Run()
	return task
}

func (t *HtopTask) Run() {
	t.checkSwap()
	t.checkRAM()
	t.checkCPU()
}

func (t *HtopTask) checkSwap() {
	swap, err := mem.SwapMemory()
	if err != nil {
		return
	}

	t.Swap = swap
}

func (t *HtopTask) checkRAM() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	t.RAM = memory
}

func (t *HtopTask) checkCPU() {
	percent, err := cpu.Percent(time.Second*1, false)
	if err != nil {
		return
	}

	t.CPU = percent[0]
}
