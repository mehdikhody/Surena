package scheduler

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"surena/node/scheduler/tasks"
	"surena/node/utils"
	"sync"
	"time"
)

var scheduler *Scheduler

type Scheduler struct {
	sync.Mutex
	logger    *zap.SugaredLogger
	isRunning bool
	cron      *cron.Cron
	HtopTask  *tasks.HtopTask
}

func init() {
	logger, err := utils.NewLogger("scheduler")
	if err != nil {
		fmt.Println("failed to create logger for scheduler")
		return
	}

	logger.Info("initializing scheduler")

	timezone := utils.GetTimezone()
	logger.Infof("timezone: %s", timezone)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		logger.Warn("failed to load timezone location")
		return
	}

	scheduler = &Scheduler{
		logger:    logger,
		isRunning: false,
		cron: cron.New(
			cron.WithLocation(location),
			cron.WithSeconds(),
		),
	}

	scheduler.HtopTask, err = tasks.NewHtopTask(scheduler.cron)
	if err != nil {
		logger.Warn("failed to create htop task")
		return
	}

	scheduler.Start()
}

func Get() (*Scheduler, error) {
	if scheduler == nil {
		return nil, errors.New("scheduler is not initialized")
	}

	return scheduler, nil
}

func (s *Scheduler) IsRunning() bool {
	return s.isRunning
}

func (s *Scheduler) Start() {
	s.Lock()
	defer s.Unlock()

	if s.IsRunning() {
		s.logger.Warn("scheduler is already running")
		return
	}

	s.cron.Start()
	s.isRunning = true
	s.logger.Info("scheduler started")
}

func (s *Scheduler) Stop() {
	s.Lock()
	defer s.Unlock()

	if !s.IsRunning() {
		s.logger.Warn("Scheduler is already stopped")
		return
	}

	s.cron.Stop()
	s.isRunning = false
	s.logger.Info("scheduler stopped")
}
