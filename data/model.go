package data

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   time.Time `gorm:"index"`
	Title       string
	Description string
	Priority    string
	Deadline    string
}
