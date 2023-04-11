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
	Deadline    Date      // dd/mm/yyyy
	Done        bool
	Title       string
	Description string
	Priority    string // enum
}

type Date struct {
	Year  string
	Month string
	Day   string
}

// Priority enum.
const (
	Critical string = "critical"
	VeryHigh string = "very high"
	High     string = "high"
	Medium   string = "medium"
	Low      string = "low"
)
