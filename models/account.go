package models

import (
	"time"

	"gorm.io/gorm"
)

func (Account) TableName() string {
	return "accounts"
}

type Account struct {
	AccountID int64          `gorm:"primaryKey"`
	Balance   float64        `gorm:"type:numeric(19,4);not null"`
	CreatedAt time.Time      `gorm:"not null;default:now()"`
	UpdatedAt time.Time      `gorm:"not null;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // GORM's soft delete field
}
