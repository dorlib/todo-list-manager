package data

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// OpenDataBase should connect to mongoDB to manage users and api-keys.
func OpenDataBase() {

}

// BeforeSave is a gorm hook in order to initiate the deadline field.
func (t *Task) BeforeSave(tx *gorm.DB) error {

}

// BeforeSave hook to hash the password before saving.
func (u *User) BeforeSave(tx *gorm.DB) error {

}
