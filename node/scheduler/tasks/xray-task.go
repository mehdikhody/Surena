package tasks

import (
	"github.com/robfig/cron/v3"
)

type XrayTask struct {
	cron.Job
}

func NewXrayTask() *XrayTask {
	return &XrayTask{}
}

func (t *XrayTask) Run() {
	// Do something
}
