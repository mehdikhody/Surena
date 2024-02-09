package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"surena/node/database/models"
)

var database *Database

type Database struct {
	DB      *gorm.DB
	Traffic *models.TrafficModel
	User    *models.UserModel
	Client  *models.ClientModel
}

func NewDatabase(dbPath string) *Database {
	dbFile := sqlite.Open(dbPath)
	db, err := gorm.Open(dbFile, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	database = &Database{
		DB:      db,
		Traffic: models.NewTrafficModel(db),
		User:    models.NewUserModel(db),
		Client:  models.NewClientModel(db),
	}

	return database
}

func GetDatabase() *Database {
	return database
}
