package database

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"surena/node/database/models"
	"surena/node/utils"
	"sync"
)

var database *Database
var databaseOnce sync.Once

type Database struct {
	db      *gorm.DB
	logger  *zap.SugaredLogger
	Traffic *models.TrafficModel
	User    *models.UserModel
	Client  *models.ClientModel
}

func initialize() error {
	logger, err := utils.NewLogger("database")
	if err != nil {
		return err
	}

	logger.Info("Initializing database")

	databasePath := GetFilePath()
	logger.Infof("Database path: %s", databasePath)

	databaseUri := fmt.Sprintf("file:%s?cache=shared", databasePath)
	file := sqlite.Open(databaseUri)

	db, err := gorm.Open(file, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		logger.Warn("Failed to open database")
		return err
	}

	clientModel, err := models.NewClientModel(db)
	if err != nil {
		logger.Warn("Failed to create client model")
		return err
	}

	database = &Database{
		db:     db,
		logger: logger,
		Client: clientModel,
	}

	return nil
}

func Get() (*Database, error) {
	var err error
	databaseOnce.Do(func() {
		err = initialize()
	})

	return database, err
}

func GetFilePath() string {
	dbpath := os.Getenv("DATABASE_PATH")
	if dbpath == "" {
		dbpath = "db/node.db"
	}

	absolutePath, _ := filepath.Abs(dbpath)
	return absolutePath
}
