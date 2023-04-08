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

func PrintTaskByName(taskName string) {
	var task Task
	DB.Where("name = ?", taskName).First(&task)

	printTask(task)
}

func PrintAllTasks() {

}

func TaskExistByID(taskID uuid.UUID) bool {
	var task Task

	tx := DB.First(&task, taskID)
	if tx != nil {
		return true
	}

	return false
}

func TaskExistByName(taskName string) bool {
	var task Task

	tx := DB.Where("name = ?", taskName).First(&task)
	if tx != nil {
		return true
	}

	return false
}
