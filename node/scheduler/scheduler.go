package scheduler

import (
	"github.com/robfig/cron/v3"
	"surena/node/scheduler/tasks"
	"time"
)

var scheduler *Scheduler
var schedulerInit = false
var schedulerStarted = false

type Scheduler struct {
	cron     *cron.Cron
	HtopTask *tasks.HtopTask
	XrayTask *tasks.XrayTask
}

func New(timezone string) *Scheduler {
	if schedulerInit {
		return scheduler
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	scheduler = &Scheduler{
		cron: cron.New(
			cron.WithLocation(location),
			cron.WithSeconds(),
		),
	}

	schedulerInit = true
	return scheduler
}

func Get() *Scheduler {
	if !schedulerInit {
		panic("Scheduler is not initialized")
	}

	return scheduler
}

func (s *Scheduler) Start() {
	if schedulerStarted {
		return
	}

	s.HtopTask = tasks.NewHtopTask()
	s.cron.AddJob("@every 1s", s.HtopTask)

	s.XrayTask = tasks.NewXrayTask()
	s.cron.AddJob("@every 1s", s.XrayTask)

	schedulerStarted = true
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	if !schedulerStarted {
		return
	}

	schedulerStarted = false
	s.cron.Stop()
}
