package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"surena/node/database/models"
	"sync"
)

var database *Database
var databaseOnce sync.Once

type Database struct {
	db      *gorm.DB
	Traffic *models.TrafficModel
	User    *models.UserModel
	Client  *models.ClientModel
}

func Get() (*Database, error) {
	var err error
	databaseOnce.Do(func() {
		databasePath := GetFilePath()
		databaseUri := fmt.Sprintf("file:%s?cache=shared", databasePath)
		file := sqlite.Open(databaseUri)

		database = &Database{}
		database.db, err = gorm.Open(file, &gorm.Config{
			PrepareStmt:          true,
			FullSaveAssociations: true,
		})

		if err != nil {
			return
		}

		database.Client, err = models.NewClientModel(database.db)
		if err != nil {
			return
		}
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
