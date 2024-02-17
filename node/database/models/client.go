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
	ClientModelInterface
	DB *gorm.DB
}

type ClientModelInterface interface {
	Find(client *Client) error
	FindByEmail(email string) (*Client, error)
	Create(email string) (*Client, error)
	UpdateTraffic(email string, upload uint64, download uint64) (*Client, error)
}

func NewClientModel(db *gorm.DB) (ClientModelInterface, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := db.WithContext(ctx).AutoMigrate(&Client{}); err != nil {
		return nil, err
	}

	return &ClientModel{
		DB: db,
	}, nil
}

func (m *ClientModel) Find(client *Client) error {
	client.Lock()
	defer client.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	find := m.DB.WithContext(ctx).Preload("Traffic").Where(client)
	if err := find.First(client).Error; err != nil {
		return err
	}

	return nil
}

func (m *ClientModel) FindByEmail(email string) (*Client, error) {
	client := &Client{Email: email}
	if err := m.Find(client); err != nil {
		return nil, err
	}

	return client, nil
}

func (m *ClientModel) Create(email string) (*Client, error) {
	client := &Client{Email: email}
	client.Lock()
	defer client.Unlock()

	if err := m.DB.Create(client).Error; err != nil {
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

	if err := m.DB.Save(client).Error; err != nil {
		return nil, err
	}

	return client, nil
}
