package database

import (
	"fmt"
	"github.com/xtls/xray-core/common/errors"
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
	Close()
}

func init() {
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
		return
	}

	clientModel, err := models.NewClientModel(db)
	if err != nil {
		logger.Warn("Failed to create client model")
		return
	}

	database = &Database{
		Filepath:    databasePath,
		DB:          db,
		ClientModel: clientModel,
	}
}

func Initialize() (DatabaseInterface, error) {
	if database == nil {
		return nil, errors.New("database is not initialized")
	}

	return database, nil
}

func Get() DatabaseInterface {
	if database == nil {
		panic("database is not initialized")
	}

	return database
}

func (d *Database) Close() {
	logger.Debug("Closing database")
	sqlDB, err := d.DB.DB()
	if err != nil {
		logger.Warn("Failed to get database connection")
		return
	}

	err = sqlDB.Close()
	if err != nil {
		logger.Warn("Failed to close database connection")
		return
	}
}

func (d *Database) GetClientModel() models.ClientModelInterface {
	return d.ClientModel
}
