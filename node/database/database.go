package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"surena/node/database/models"
	"surena/node/env"
	"surena/node/utils"
)

var database *Database

type Database struct {
	DatabaseInterface
	Logger      *logrus.Entry
	Filepath    string
	DB          *gorm.DB
	ClientModel models.ClientModelInterface
}

type DatabaseInterface interface {
	GetClientModel() models.ClientModelInterface
}

func init() {
	logger := utils.CreateLogger("database")
	logger.Debug("Initializing database")

	databasePath := env.GetDatabasePath()
	os.MkdirAll(filepath.Dir(databasePath), os.ModePerm)
	logger.Debug("Database path: %s", databasePath)

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
		Logger:      logger,
		Filepath:    databasePath,
		DB:          db,
		ClientModel: clientModel,
	}
}

func Get() DatabaseInterface {
	if database == nil {
		panic("database is not initialized")
	}

	return database
}

func (d *Database) GetClientModel() models.ClientModelInterface {
	return d.ClientModel
}
