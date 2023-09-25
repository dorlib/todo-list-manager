package domain

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	Users       []User `gorm:"-"`
}
