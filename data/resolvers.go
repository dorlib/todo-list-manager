package data

import (
	"fmt"
)

const (
	DONE   = "done"
	UNDONE = "undone"
)

func CreateTask(title, description, priority string, deadline Date) {
	task := Task{Title: title, Description: description, Priority: priority, DeadlineDate: deadline}
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

	DB.Where("Title = ?", title).Delete(&tasks)
	fmt.Printf("deleted task: %v", title)
}

func PrintTaskByID(taskID uint) {
	var task Task

	DB.First(&task, taskID)

	printTask(task)
}

func PrintTaskByName(taskName string) {
	task := Task{}

	DB.Where("Title = ?", taskName).First(&task)

	printTask(task)
}

func PrintAllTasks(user User, userExist bool, by string, opt string) {
	if userExist {
		printTasks(GetAllTasksOfUser(user, by, opt))

		return
	}

	printTasks(GetAllTasksOfGroup(by, opt))
}

func GetAllTasksOfUser(user User, by string, opt string) []taskSummery {
	var tasks []taskSummery

	switch {
	case by == "" && opt == "":
		DB.Select("id, title, description, priority, created_at, deadline, done").Where(user).Find(&tasks)
	case by == "" && opt != "":
		switch opt {
		case DONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = true", user.ID).Scan(&tasks)
		case UNDONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = false", user.ID).Scan(&tasks)
		default:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND priority = ?", user.ID, opt).Scan(&tasks)
		}
	case by != "" && opt == "":
		DB.Raw("SELECT tasks FROM users WHERE id = ? ORDER BY = ?", user.ID, by).Scan(&tasks)
	default:
		switch opt {
		case DONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = true ORDER BY = ?", user.ID, by).Scan(&tasks)
		case UNDONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = false ORDER BY = ?", user.ID, by).Scan(&tasks)
		default:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND priority = ? ORDER BY = ?", user.ID, opt, by).Scan(&tasks)
		}
	}

	return tasks
}

func GetAllTasksOfGroup(by, opt string) []taskSummery {
	var tasks []taskSummery

	switch {
	case by == "" && opt == "":
		DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Scan(&tasks)
	case by == "" && opt != "":
		switch opt {
		case DONE:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("done = ?", true).Scan(&tasks)
		case UNDONE:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("done = ?", false).Scan(&tasks)
		default:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("priority = ?", opt).Scan(&tasks)
		}
	case by != "" && opt == "":
		DB.Table("tasks").Order(by).Scan(&tasks)
	default:
		switch opt {
		case DONE:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done FROM tasks WHERE done = true ORDER BY = ?", by).Scan(&tasks)
		case UNDONE:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done FROM tasks WHERE done = false ORDER BY = ?", by).Scan(&tasks)
		default:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done  FROM tasks WHERE priority = ? ORDER BY = ?", opt, by).Scan(&tasks)
		}
	}

	return tasks
}

func PrintByDeadLine(user User, userExist bool, opt string) {
	var tasks []taskSummery

	if userExist {
		switch opt {
		case DONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = true ORDER BY deadline", user.ID).Scan(&tasks)
		case UNDONE:
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND done = false ORDER BY deadline", user.ID).Scan(&tasks)
		case "with-priority":
			DB.Raw("SELECT tasks FROM users WHERE id = ? AND priority = ? ORDER BY deadline", user.ID, opt).Scan(&tasks)
		default:
			DB.Raw("SELECT tasks FROM users WHERE id = ? ORDER BY deadline", user.ID).Scan(&tasks)
		}
	} else {
		DB.Raw("SELECT * FROM tasks ORDER BY deadline").Scan(&tasks)
	}

	printTasks(tasks)
}

func PrintByPriority(user User, userExist bool) {
	var tasks []taskSummery

	if userExist {
		DB.Raw("SELECT tasks FROM users WHERE id = ? ORDER BY priority", user.ID).Scan(&tasks)
	} else {
		DB.Raw("SELECT * FROM tasks ORDER BY priority").Scan(&tasks)
	}

	printTasks(tasks)
}

func PrintByCreationDate(user User, userExist bool) {
	var tasks []taskSummery

	if userExist {
		DB.Raw("SELECT tasks FROM users WHERE id = ? ORDER BY created_at", user.ID).Scan(&tasks)
	} else {
		DB.Raw("SELECT * FROM tasks ORDER BY created_at").Scan(&tasks)
	}

	printTasks(tasks)
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

func GetUser(userID uint, username string) (User, bool) {
	var user User

	r := DB.Where("id = ? OR username = ?", userID, username).First(&user)
	if r.RowsAffected != 0 {
		return user, true
	}

	return user, false
}
