package models

import (
	"time"

	"gorm.io/gorm"
)

func (Transaction) TableName() string {
	return "transaction"
}

type Transaction struct {
	ID                   int64          `gorm:"primaryKey"`
	SourceAccountID      int64          `gorm:"not null"`
	DestinationAccountID int64          `gorm:"not null"`
	Amount               float64        `gorm:"type:numeric(19,4);not null;check:amount > 0"`
	Status               string         `gorm:"type:varchar(50);not null"`
	CreatedAt            time.Time      `gorm:"not null;default:now()"`
	UpdatedAt            time.Time      `gorm:"not null;default:now()"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	SourceAccount        Account        `gorm:"foreignKey:SourceAccountID"`
	DestinationAccount   Account        `gorm:"foreignKey:DestinationAccountID"`
}
