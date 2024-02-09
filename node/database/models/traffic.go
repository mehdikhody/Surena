package models

import (
	"gorm.io/gorm"
	"time"
)

type Traffic struct {
	ID        uint `gorm:"primaryKey"`
	RefID     uint
	RefType   string
	Upload    uint64
	Download  uint64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type TrafficModel struct {
	db *gorm.DB
}

func NewTrafficModel(db *gorm.DB) *TrafficModel {
	if db.AutoMigrate(&Traffic{}) != nil {
		panic("failed to migrate traffic model")
	}

	return &TrafficModel{
		db: db,
	}
}
