package data

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	DeadlineDate Date           `gorm:"embedded"` // dd/mm/yyyy
	Deadline     string
	Done         bool
	Title        string `gorm:"not null"`
	Description  string
	Priority     string // enum
	User         User   `gorm:"-;unique;not null"`
}

type Date struct {
	DeadlineYear  string
	DeadlineMonth string
	DeadlineDay   string
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

type taskSummery struct {
	ID          uint
	Title       string
	Description string
	Priority    string
	CreatedAt   time.Time
	Deadline    string
	Done        bool
}

// Priority enum.
const (
	Critical string = "critical"
	VeryHigh string = "very high"
	High     string = "high"
	Medium   string = "medium"
	Low      string = "low"
)
