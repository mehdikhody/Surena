package scheduler

import (
	"errors"
	"github.com/robfig/cron/v3"
	"os"
	"surena/node/scheduler/tasks"
	"sync"
	"time"
)

var scheduler *Scheduler
var schedulerOnce sync.Once
var schedulerStarted = false

type Scheduler struct {
	cron     *cron.Cron
	HtopTask *tasks.HtopTask
	XrayTask *tasks.XrayTask
}

func Get() *Scheduler {
	schedulerOnce.Do(func() {
		timezone := GetTimezone()
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

		scheduler.HtopTask = tasks.NewHtopTask()
		scheduler.cron.AddJob("@every 5s", scheduler.HtopTask)

		scheduler.XrayTask = tasks.NewXrayTask()
		scheduler.cron.AddJob("@every 5s", scheduler.XrayTask)
	})

	return scheduler
}

func GetTimezone() string {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone
}

func (s *Scheduler) Start() error {
	if schedulerStarted {
		return errors.New("scheduler is already started")
	}

	schedulerStarted = true
	s.cron.Start()
	return nil
}

func (s *Scheduler) Stop() error {
	if !schedulerStarted {
		return errors.New("scheduler is already stopped")
	}

	schedulerStarted = false
	s.cron.Stop()
	return nil
}
