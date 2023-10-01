package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          uuid.UUID      `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"unique;not null"`
	Description string
	Users       []User `gorm:"-"`
	Tasks       []Task `gorm:"-"`
}
