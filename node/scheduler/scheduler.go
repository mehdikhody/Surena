package scheduler

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"surena/node/env"
	"surena/node/scheduler/tasks"
	"surena/node/utils"
	"sync"
	"time"
)

var scheduler *Scheduler

type Scheduler struct {
	SchedulerInterface
	sync.Mutex
	Logger            *logrus.Entry
	Started           bool
	Cron              *cron.Cron
	SystemWatcherTask tasks.SystemWatcherTaskInterface
}

type SchedulerInterface interface {
	IsRunning() bool
	Start()
	Stop()
}

func init() {
	logger := utils.CreateLogger("scheduler")
	logger.Debug("initializing scheduler")

	timezone := env.GetTimezone()
	logger.Debugf("timezone: %s", timezone)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		logger.Warn("failed to load timezone location")
		return
	}

	cronjob := cron.New(
		cron.WithLocation(location),
		cron.WithSeconds(),
	)

	scheduler = &Scheduler{
		Logger:            logger,
		Started:           false,
		Cron:              cronjob,
		SystemWatcherTask: tasks.NewSystemWatcherTask(cronjob),
	}

	scheduler.Start()
}

func Get() SchedulerInterface {
	if scheduler == nil {
		panic("scheduler is not initialized")
	}

	return scheduler
}

func (s *Scheduler) IsRunning() bool {
	return s.Started
}

func (s *Scheduler) Start() {
	s.Lock()
	defer s.Unlock()

	if s.Started {
		s.Logger.Warn("scheduler is already running")
		return
	}

	s.Cron.Start()
	s.Started = true
	s.Logger.Info("scheduler started")
}

func (s *Scheduler) Stop() {
	s.Lock()
	defer s.Unlock()

	if !s.Started {
		s.Logger.Warn("Scheduler is already stopped")
		return
	}

	s.Cron.Stop()
	s.Started = false
	s.Logger.Info("scheduler stopped")
}
