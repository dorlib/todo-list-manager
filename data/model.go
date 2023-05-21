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
	Status       Status // enum
	Title        string `gorm:"not null"`
	Description  string
	Priority     string // enum
	UserID       uint   `gorm:"column:user_id"`
	UserName     string `gorm:"column:user_name"`
	User         User
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
	Password  string         `gorm:"size:255"`
	Role      string         // admin/user
	Tasks     []Task         `gorm:"-"`
}

type Group struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	Users       []User `gorm:"-"`
}

type taskSummery struct {
	ID          uint
	UserName    string
	Title       string
	Description string
	Priority    string
	CreatedAt   string
	Deadline    string
	Done        bool
	TimeLeft    int
}

// Priority enum.
const (
	Critical string = "critical"
	VeryHigh string = "very high"
	High     string = "high"
	Medium   string = "medium"
	Low      string = "low"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

// returns the corresponding string value to the Card enum.
func (s Status) String() string {
	switch s {
	case Todo:
		return "todo"
	case InProgress:
		return "in progress"
	case Done:
		return "done"
	}

	return "Unknown status"
}
