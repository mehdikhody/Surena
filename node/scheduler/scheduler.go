package scheduler

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/errors"
	"surena/node/env"
	"surena/node/scheduler/tasks"
	"surena/node/utils"
	"sync"
	"time"
)

var scheduler *Scheduler
var logger = utils.CreateLogger("scheduler")

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
		Started:           false,
		Cron:              cronjob,
		SystemWatcherTask: tasks.NewSystemWatcherTask(cronjob),
	}
}

func Initialize() (SchedulerInterface, error) {
	if scheduler == nil {
		return nil, errors.New("scheduler is not initialized")
	}

	return scheduler, nil
}

func Get() SchedulerInterface {
	if scheduler == nil {
		panic("scheduler is not initialized")
	}

	return scheduler
}

func (s *Scheduler) IsRunning() bool {
	s.Lock()
	defer s.Unlock()
	return s.Started
}

func (s *Scheduler) Start() {
	s.Lock()
	defer s.Unlock()

	if s.Started {
		logger.Warn("scheduler is already running")
		return
	}

	s.Cron.Start()
	s.Started = true
	logger.Info("scheduler started")
}

func (s *Scheduler) Stop() {
	s.Lock()
	defer s.Unlock()

	if !s.Started {
		logger.Warn("Scheduler is already stopped")
		return
	}

	s.Cron.Stop()
	s.Started = false
	logger.Info("scheduler stopped")
}
