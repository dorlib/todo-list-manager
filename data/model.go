package data

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Deadline    Date           `gorm:"-"` // dd/mm/yyyy
	Done        bool
	Title       string `gorm:"not null"`
	Description string
	Priority    string // enum
	User        User   `gorm:"-;unique;not null"`
}

type Date struct {
	Year  string
	Month string
	Day   string
}

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;not null"`
	Role      string         // admin/user
	Tasks     []Task         `gorm:"-"`
}

// Priority enum.
const (
	Critical string = "critical"
	VeryHigh string = "very high"
	High     string = "high"
	Medium   string = "medium"
	Low      string = "low"
)
