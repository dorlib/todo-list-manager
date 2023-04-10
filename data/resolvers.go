package data

import (
	"github.com/google/uuid"
)

func CreateTask(title, description, priority string, deadline Date) {
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
	//tasks := DB.Find(&Task{})
	//printAllTasks(tasks)
}
