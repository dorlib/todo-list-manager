package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;not null"`
	Password  string         `gorm:"size:255"`
	PublicKey string
	Role      string // admin/user
	Tasks     []Task `gorm:"-"`
}

// BeforeSave hook to hash the password before saving.
func (u *User) BeforeSave(tx *gorm.DB) error {
	if len(u.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
