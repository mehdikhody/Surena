package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"surena/node/database/models"
	"surena/node/env"
	"surena/node/utils"
)

var database *Database
var logger = utils.CreateLogger("database")

type Database struct {
	DatabaseInterface
	Filepath    string
	DB          *gorm.DB
	ClientModel models.ClientModelInterface
}

type DatabaseInterface interface {
	GetClientModel() models.ClientModelInterface
}

func Initialize() (DatabaseInterface, error) {
	logger.Debug("Initializing database")

	databasePath := env.GetDatabasePath()
	os.MkdirAll(filepath.Dir(databasePath), os.ModePerm)
	logger.Debugf("Database path: %s", databasePath)

	databaseUri := fmt.Sprintf("file:%s?cache=shared", databasePath)
	file := sqlite.Open(databaseUri)

	db, err := gorm.Open(file, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		logger.Warn("Failed to open database")
		return nil, err
	}

	clientModel, err := models.NewClientModel(db)
	if err != nil {
		logger.Warn("Failed to create client model")
		return nil, err
	}

	database = &Database{
		Filepath:    databasePath,
		DB:          db,
		ClientModel: clientModel,
	}

	logger.Debug("Database initialized")
	return database, nil
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
