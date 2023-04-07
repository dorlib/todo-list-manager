package data

import (
	"github.com/google/uuid"
	"time"
)

func CreateTask(title, description, priority string, deadline time.Time) {
	task := Task{Title: title, Description: description, Priority: priority, Deadline: deadline}
	DB.Create(&task)
}

func DeleteTask(taskID uuid.UUID) {
	DB.Delete("ID = ?", taskID)
}

func PrintTask(taskID uuid.UUID) {
	var task Task
	DB.First(&task, taskID)

	printTask(task)
}

func PrintAllTasks() {

}
