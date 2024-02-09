package models

import "time"

type Inbound struct {
	ID        uint   `gorm:"primaryKey"`
	Enable    bool   `gorm:"default:true"`
	Tag       string `gorm:"unique"`
	Title     string
	Listen    string `gorm:"default:127.0.0.1"`
	Port      uint
	Protocol  string
	Traffic   Traffic   `gorm:"polymorphic:Ref;polymorphicValue:inbound"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
