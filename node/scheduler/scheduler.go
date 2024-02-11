package scheduler

import (
	"errors"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
	"surena/node/scheduler/tasks"
	"surena/node/utils"
	"sync"
	"time"
)

var scheduler *Scheduler
var schedulerOnce sync.Once
var schedulerStarted = false

type Scheduler struct {
	sync.Mutex
	cron     *cron.Cron
	logger   *zap.SugaredLogger
	HtopTask *tasks.HtopTask
}

func initialize() error {
	logger, err := utils.NewLogger("scheduler")
	if err != nil {
		return err
	}

	logger.Info("Initializing scheduler")

	timezone := GetTimezone()
	logger.Infof("Timezone: %s", timezone)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		logger.Warn("Failed to load timezone location")
		return err
	}

	scheduler = &Scheduler{
		logger: logger,
		cron: cron.New(
			cron.WithLocation(location),
			cron.WithSeconds(),
		),
	}

	scheduler.HtopTask, err = tasks.NewHtopTask(scheduler.cron)
	if err != nil {
		logger.Warn("Failed to create htop task")
		return err
	}

	return nil
}

func Get() (*Scheduler, error) {
	var err error
	schedulerOnce.Do(func() {
		err = initialize()
	})

	return scheduler, err
}

func GetTimezone() string {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone
}

func (s *Scheduler) Start() error {
	s.Lock()
	defer s.Unlock()

	if schedulerStarted {
		s.logger.Warn("Scheduler is already started")
		return errors.New("scheduler is already started")
	}

	s.logger.Info("Starting scheduler")
	schedulerStarted = true
	s.cron.Start()
	return nil
}

func (s *Scheduler) Stop() error {
	s.Lock()
	defer s.Unlock()

	if !schedulerStarted {
		s.logger.Warn("Scheduler is already stopped")
		return errors.New("scheduler is already stopped")
	}

	s.logger.Info("Stopping scheduler")
	schedulerStarted = false
	s.cron.Stop()
	return nil
}
