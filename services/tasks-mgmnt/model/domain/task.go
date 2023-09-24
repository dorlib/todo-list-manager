package domain

import (
	"fmt"
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

// BeforeSave is a gorm hook in order to initiate the deadline field.
func (t *Task) BeforeSave(tx *gorm.DB) error {
	t.Deadline = fmt.Sprintf("%s/%s/%s", t.DeadlineDate.DeadlineDay, t.DeadlineDate.DeadlineMonth, t.DeadlineDate.DeadlineYear)

	return nil
}
