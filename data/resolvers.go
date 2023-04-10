package data

import (
	"fmt"
	"github.com/google/uuid"
)

func CreateTask(title, description, priority string, deadline Date) {
	task := Task{Title: title, Description: description, Priority: priority, Deadline: deadline}
	DB.Create(&task)
}

func DeleteTaskByID(taskID uuid.UUID) {
	if TaskExistByID(taskID) {
		DB.Delete(&Task{}, taskID)
		fmt.Printf("deleted task: %v", taskID)

		return
	}

	fmt.Printf("there is no tasks with the given ID: %v", taskID)
}

func DeleteTaskByTitle(taskTitle string) {
	var tasks []Task

	DB.Where("Title = ?", taskTitle).Find(&tasks)

	if len(tasks) > 1 {
		fmt.Printf("more than one task with the same title %v exists", taskTitle)

		return
	}

	if len(tasks) < 1 {
		fmt.Printf("there is no tasks with the given title: %v", taskTitle)

		return
	}

	DB.Where("name = ?", taskTitle).Delete(&tasks)
	fmt.Printf("deleted task: %v", taskTitle)
}

func PrintTask(taskID uuid.UUID) {
	var task Task

	DB.First(&task, taskID)

	printTask(task)
}

func PrintTaskByName(taskName string) {
	task := Task{}

	DB.Where("name = ?", taskName).First(&task)

	printTask(task)
}

func PrintAllTasks() {
	var tasks []Task

	DB.Find(&tasks)

	printAllTasks(tasks)
}
