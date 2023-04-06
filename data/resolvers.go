package data

import "github.com/google/uuid"

func CreateTask(title, description, priority, deadline string) {
	task := Task{Title: title, Description: description, Priority: priority, Deadline: deadline}
	DB.Create(&task)
}

func DeleteTask(taskID uuid.UUID) {
	DB.Delete("ID = ?", taskID)
}
