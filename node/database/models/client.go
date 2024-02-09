package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Enable    bool       `gorm:"default:true" json:"enable"`
	Email     string     `gorm:"unique" json:"email"`
	Bandwidth uint64     `gorm:"default:0" json:"bandwidth"`
	Traffic   Traffic    `gorm:"polymorphic:Ref;polymorphicValue:client" json:"traffic"`
	ExpiresAt *time.Time `gorm:"default:null" json:"expires_at"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type ClientModel struct {
	db *gorm.DB
}

func NewClientModel(db *gorm.DB) *ClientModel {
	if db.AutoMigrate(&Client{}) != nil {
		panic("failed to migrate client model")
	}

	return &ClientModel{
		db: db,
	}
}

func (m *ClientModel) Create(email string) (*Client, error) {
	client := &Client{Email: email}

	err := m.db.Create(client).Error
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *ClientModel) UpdateTraffic(email string, upload uint64, download uint64) (*Client, error) {
	client := &Client{Email: email}
	err := m.db.Preload("Traffic").Where(client).First(client).Error

	if err != nil {
		return nil, err
	}

	client.Traffic.Upload += upload
	client.Traffic.Download += download

	err = m.db.Save(client).Error
	if err != nil {
		return nil, err
	}

	return client, nil
}
