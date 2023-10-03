package impl

import (
	"fmt"
	"tasks-mgmnt/model/domain"
	"tasks-mgmnt/repo/impl"
)

func CreateTask(title, description, priority string, deadline domain.Date, user domain.User) (domain.Task, error) {
	task := domain.Task{
		Title:        title,
		Description:  description,
		Priority:     priority,
		DeadlineDate: deadline,
		UserID:       user.ID,
		UserName:     user.Username,
	}

	rows := impl.DB.Create(&task).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return task, nil
	}

	return domain.Task{}, fmt.Errorf("failed to create task")
}
