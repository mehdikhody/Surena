package database

import (
	"fmt"
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

func New(path string) *Database {
	uri := fmt.Sprintf("file:%s?cache=shared", path)
	file := sqlite.Open(uri)
	db, err := gorm.Open(file, &gorm.Config{
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

func Get() *Database {
	return database
}
