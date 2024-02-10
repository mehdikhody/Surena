package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"surena/node/database/models"
)

var database *Database
var databaseInitialized = false

type Database struct {
	db      *gorm.DB
	Traffic *models.TrafficModel
	User    *models.UserModel
	Client  *models.ClientModel
}

func Initialize() *Database {
	if databaseInitialized {
		panic("Database already initialized")
	}

	databasePath := GetFilePath()
	file := sqlite.Open(fmt.Sprintf("file:%s?cache=shared", databasePath))
	db, err := gorm.Open(file, &gorm.Config{
		PrepareStmt:          true,
		FullSaveAssociations: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	database = &Database{
		db:      db,
		Traffic: models.NewTrafficModel(db),
		User:    models.NewUserModel(db),
		Client:  models.NewClientModel(db),
	}

	databaseInitialized = true
	return database
}

func Get() *Database {
	if !databaseInitialized {
		panic("Database not initialized")
	}

	return database
}

func GetFilePath() string {
	dbpath := os.Getenv("DATABASE_PATH")
	if dbpath == "" {
		dbpath = "db/node.db"
	}

	absolutePath, _ := filepath.Abs(dbpath)
	return absolutePath
}
