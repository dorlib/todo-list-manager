package data

import (
	"fmt"
)

func CreateTask(title, description, priority string, deadline Date) {
	task := Task{Title: title, Description: description, Priority: priority, Deadline: deadline}
	rows := DB.Create(&task).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)
}

func DeleteTaskByID(taskID uint) {
	if TaskExistByID(taskID) {
		DB.Delete(&Task{}, taskID)
		fmt.Printf("deleted task: %v \n", taskID)

		return
	}

	fmt.Printf("there is no tasks with the given ID: %v", taskID)
}

func DeleteTaskByTitle(title string) {
	var tasks []Task

	DB.Where("Title = ?", title).Find(&tasks)

	if len(tasks) > 1 {
		fmt.Printf("more than one task with the same title %v exists", title)

		return
	}

	if len(tasks) < 1 {
		fmt.Printf("there is no tasks with the given title: %v", title)

		return
	}

	DB.Where("name = ?", title).Delete(&tasks)
	fmt.Printf("deleted task: %v", title)
}

func PrintTaskByID(taskID uint) {
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

func ToggleDoneByTitle(title string, isDone bool) {
	var tasks []Task

	DB.Where("Title = ?", title).Find(&tasks)

	if len(tasks) > 1 {
		fmt.Printf("more than one task with the same title %v exists, please mark done by id instead", title)

		return
	}

	if len(tasks) < 1 {
		fmt.Printf("there is no task with the given title: %v", title)

		return
	}

	if isDone {
		DB.Where("name = ?", title).Update("Done", true)
	} else {
		DB.Where("name = ?", title).Update("Done", false)
	}

	fmt.Printf("task: %v is done", title)
}

func ToggleDoneByID(taskID uint, isDone bool) {
	var task Task

	DB.First(&task)

	if TaskExistByID(taskID) {
		if isDone {
			DB.Model(&task).Update("Done", true)
		} else {
			DB.Model(&task).Update("Done", false)
		}

		fmt.Printf("task: %v is been marked as done \n", taskID)

		return
	}

	fmt.Printf("there is no task with the given ID: %v \n", taskID)
}
