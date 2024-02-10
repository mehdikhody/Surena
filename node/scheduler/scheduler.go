package scheduler

import (
	"github.com/robfig/cron/v3"
	"os"
	"surena/node/scheduler/tasks"
	"time"
)

var scheduler *Scheduler
var schedulerInitialized = false
var schedulerStarted = false

type Scheduler struct {
	cron     *cron.Cron
	HtopTask *tasks.HtopTask
	XrayTask *tasks.XrayTask
}

func Initialize() *Scheduler {
	if schedulerInitialized {
		panic("Scheduler is already initialized")
	}

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

	schedulerInitialized = true
	return scheduler
}

func Get() *Scheduler {
	if !schedulerInitialized {
		panic("Scheduler is not initialized")
	}

	return scheduler
}

func GetTimezone() string {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone
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
