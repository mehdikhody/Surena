package database

import (
	errors "errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"surena/node/database/models"
	"surena/node/utils"
)

var database *db

type db struct {
	logger *zap.SugaredLogger
	dbpath string
	gorm   *gorm.DB
	client *models.ClientModel
}

type DBInterface interface {
	GetClient() (*models.ClientModel, error)
}

func init() {
	logger, err := utils.NewLogger("database")
	if err != nil {
		fmt.Println("Failed to create logger for database")
		return
	}

	logger.Info("Initializing database")

	databasePath := utils.GetDatabasePath()
	logger.Infof("Database path: %s", databasePath)

	databaseUri := fmt.Sprintf("file:%s?cache=shared", databasePath)
	file := sqlite.Open(databaseUri)

	gorm, err := gorm.Open(file, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		logger.Warn("Failed to open database")
		return
	}

	clientModel, err := models.NewClientModel(gorm)
	if err != nil {
		logger.Warn("Failed to create client model")
		return
	}

	database = &db{
		logger: logger,
		dbpath: databasePath,
		gorm:   gorm,
		client: clientModel,
	}
}

func Get() (DBInterface, error) {
	if database == nil {
		return nil, errors.New("database is not initialized")
	}

	return database, nil
}

func (d *db) GetClient() (*models.ClientModel, error) {
	if d.client == nil {
		return nil, errors.New("client model is not initialized")
	}

	return d.client, nil
}
