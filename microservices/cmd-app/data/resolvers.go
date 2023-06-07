package data

import (
	"fmt"
)

const (
	DONE   = "done"
	UNDONE = "undone"
)

func CreateTask(title, description, priority string, deadline Date, user User) (Task, error) {
	task := Task{
		Title:        title,
		Description:  description,
		Priority:     priority,
		DeadlineDate: deadline,
		UserID:       user.ID,
		UserName:     user.Username,
	}

	rows := DB.Create(&task).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return task, nil
	}

	return Task{}, fmt.Errorf("failed to create task")
}

func CreateUser(username, role, password string) (User, error) {
	if UserExistByName(username) {
		return User{}, fmt.Errorf("user already exists")
	}
	user := User{
		Username: username,
		Password: password,
		Role:     role,
	}

	rows := DB.Create(&user).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return user, nil
	}

	return User{}, fmt.Errorf("failed to create user")
}

func CreateGroup(name, description string, users []User) (Group, error) {
	group := Group{
		Name:        name,
		Description: description,
		Users:       users,
	}

	rows := DB.Create(&group).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return group, nil
	}

	return Group{}, fmt.Errorf("failed to create group")
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

	userID := user.ID

	switch {
	case by == "" && opt == "":
		DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("userID = ?", userID).Scan(&tasks)
	case by == "" && opt != "":
		switch opt {
		case DONE:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("done = ? AND userID = ?", true, userID).Scan(&tasks)
		case UNDONE:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("done = ? AND userID = ?", false, userID).Scan(&tasks)
		default:
			DB.Table("tasks").Select("id, user_name, title, description, priority, created_at, deadline, done").Where("priority = ? AND userID = ?", opt, userID).Scan(&tasks)
		}
	case by != "" && opt == "":
		DB.Table("tasks").Where("userID = ?", userID).Order(by).Scan(&tasks)
	default:
		switch opt {
		case DONE:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done FROM tasks WHERE done = true AND userID = ? ORDER BY = ?", userID, by).Scan(&tasks)
		case UNDONE:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done FROM tasks WHERE done = false AND userID = ? ORDER BY = ?", userID, by).Scan(&tasks)
		default:
			DB.Raw("SELECT id, user_name, title, description, priority, created_at, deadline, done  FROM tasks WHERE priority = ? AND userID = ? ORDER BY = ?", userID, opt, by).Scan(&tasks)
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

func GetTask(taskID uint, taskTitle string) Task {
	if taskID != 0 {
		task, exists := getTaskByID(taskID)
		if exists {
			return task
		}
	}

	if taskTitle != "" {
		task, rowsAffected := getTaskByTitle(taskTitle)
		if rowsAffected == 1 {
			return task
		}
	}

	fmt.Printf("task with title %v or id %v doesnt exists", taskTitle, taskID)

	return Task{}
}

func getTaskByID(taskID uint) (Task, bool) {
	var task Task

	if !TaskExistByID(taskID) {
		fmt.Printf("task with the id %v doesnt exists", taskID)
	}

	r := DB.Where("id = ?", taskID).First(&task)
	if r.RowsAffected != 0 {
		return task, true
	}

	return task, false
}

func getTaskByTitle(taskTitle string) (Task, int64) {
	var task Task

	if !TaskExistByName(taskTitle) {
		fmt.Printf("task with the id %v doesnt exists", taskTitle)
	}

	r := DB.Where("title = ?", taskTitle).Find(&task)
	if r.RowsAffected > 1 {
		fmt.Printf("more than one tasks exists with the title: %v, please specify task id", taskTitle)

		return Task{}, r.RowsAffected
	}

	if r.RowsAffected < 1 {
		fmt.Printf("there is no task with the title: %v", taskTitle)

		return Task{}, 0
	}

	return task, 1
}

func UpdateTaskByID(taskID uint, argsMap map[string]interface{}) {
	task, found := getTaskByID(taskID)
	if !found {
		fmt.Printf("task with the id %v doesnt exists", taskID)
	}

	if argsMap["title"] != "" {
		task.Title = argsMap["title"].(string)
	}

	if argsMap["description"] != "" {
		task.Title = argsMap["description"].(string)
	}

	if argsMap["priority"] != "" {
		task.Title = argsMap["priority"].(string)
	}

	if argsMap["deadline"] != "" {
		task.Title = argsMap["deadline"].(string)
	}

	DB.Save(&task)
}

func UpdateTaskByTitle(task Task, argsMap map[string]interface{}) {
	if argsMap["title"] != "" {
		task.Title = argsMap["title"].(string)
	}

	if argsMap["description"] != "" {
		task.Title = argsMap["description"].(string)
	}

	if argsMap["priority"] != "" {
		task.Title = argsMap["priority"].(string)
	}

	if argsMap["deadline"] != "" {
		task.Title = argsMap["deadline"].(string)
	}

	DB.Save(&task)
}
