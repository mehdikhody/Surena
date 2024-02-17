package database

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"surena/node/database/models"
	"surena/node/env"
	"surena/node/utils"
)

var database *Database

type Database struct {
	DatabaseInterface
	Logger   *logrus.Entry
	Filepath string
	DB       *gorm.DB
	Client   models.ClientModelInterface
}

type DatabaseInterface interface {
	GetClient() (models.ClientModelInterface, error)
}

func init() {
	logger := utils.CreateLogger("database")
	logger.Info("Initializing database")

	databasePath := env.GetDatabasePath()
	logger.Infof("Database path: %s", databasePath)

	databaseUri := fmt.Sprintf("file:%s?cache=shared", databasePath)
	file := sqlite.Open(databaseUri)

	db, err := gorm.Open(file, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		logger.Warn("Failed to open database")
		return
	}

	clientModel, err := models.NewClientModel(db)
	if err != nil {
		logger.Warn("Failed to create client model")
		return
	}

	database = &Database{
		Logger:   logger,
		Filepath: databasePath,
		DB:       db,
		Client:   clientModel,
	}
}

func Get() (DatabaseInterface, error) {
	if database == nil {
		return nil, errors.New("database is not initialized")
	}

	return database, nil
}

func (d *Database) GetClient() (models.ClientModelInterface, error) {
	if d.Client == nil {
		return nil, errors.New("client model is not initialized")
	}

	return d.Client, nil
}
