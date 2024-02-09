package scheduler

import (
	"github.com/robfig/cron/v3"
	"surena/node/scheduler/tasks"
	"time"
)

var scheduler *Scheduler

type Scheduler struct {
	cron     *cron.Cron
	HtopTask *tasks.HtopTask
	XrayTask *tasks.XrayTask
}

func New(timezone string) *Scheduler {
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

	return scheduler
}

func GetScheduler() *Scheduler {
	return scheduler
}

func (s *Scheduler) Start() {
	s.HtopTask = tasks.NewHtopTask()
	s.cron.AddJob("@every 1s", s.HtopTask)

	s.XrayTask = tasks.NewXrayTask()
	s.cron.AddJob("@every 1s", s.XrayTask)

	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
