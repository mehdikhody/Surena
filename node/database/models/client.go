package models

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Client struct {
	sync.Mutex
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
	gorm *gorm.DB
}

func NewClientModel(gorm *gorm.DB) (*ClientModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := gorm.WithContext(ctx).AutoMigrate(&Client{}); err != nil {
		return nil, err
	}

	return &ClientModel{
		gorm: gorm,
	}, nil
}

func (m *ClientModel) Find(client *Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	find := m.gorm.WithContext(ctx).Preload("Traffic").Where(client)
	if err := find.First(client).Error; err != nil {
		return err
	}

	return nil
}

func (m *ClientModel) Create(email string) (*Client, error) {
	client := &Client{Email: email}
	client.Lock()
	defer client.Unlock()

	if err := m.gorm.Create(client).Error; err != nil {
		return nil, err
	}

	return client, nil
}

func (m *ClientModel) UpdateTraffic(email string, upload uint64, download uint64) (*Client, error) {
	client := &Client{Email: email}
	client.Lock()
	defer client.Unlock()

	if err := m.Find(client); err != nil {
		return nil, err
	}

	client.Traffic.Upload += upload
	client.Traffic.Download += download

	if err := m.gorm.Save(client).Error; err != nil {
		return nil, err
	}

	return client, nil
}
