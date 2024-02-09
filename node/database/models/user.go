package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       uint8
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	if db.AutoMigrate(&User{}) != nil {
		panic("failed to migrate user model")
	}

	return &UserModel{
		db: db,
	}
}

func (m *UserModel) Create(name string, age uint8) (*User, error) {
	user := &User{
		Name: name,
		Age:  age,
	}

	err := m.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
