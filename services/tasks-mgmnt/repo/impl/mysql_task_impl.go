package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"
)

type MySQLTaskRepo struct{}

func NewMySQLTaskRepo() *MySQLTaskRepo {
	return &MySQLTaskRepo{}
}

func (m *MySQLTaskRepo) CreateOne(ctx context.Context, db gorm.DB, task *domain.Task) (*domain.Task, error) {
	rows := db.WithContext(ctx).Create(&task).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return task, nil
	}

	return nil, eris.Wrap(fmt.Errorf("failed to create task"), "DB error")
}

func (m *MySQLTaskRepo) GetOne(ctx context.Context, db gorm.DB, taskID uuid.UUID) (*domain.Task, error) {
	var task *domain.Task

	r := db.WithContext(ctx).Where("id = ?", taskID).First(&task)
	if r.RowsAffected != 0 {
		return task, nil
	}

	return nil, eris.Wrap(fmt.Errorf("task with the id %v doesnt exists", taskID), "DB error")
}

func (m *MySQLTaskRepo) DeleteOne(ctx context.Context, db gorm.DB, taskID uuid.UUID) error {
	r := db.WithContext(ctx).Delete(&domain.Task{}, taskID)
	if r.RowsAffected != 0 {
		fmt.Printf("deleted task: %v \n", taskID)

		return nil
	}

	return eris.Wrap(fmt.Errorf("there is no task with the given ID: %v", taskID), "DB error")
}

func (m *MySQLTaskRepo) GetAll(ctx context.Context, db gorm.DB) (*[]domain.Task, error) {
	var tasks *[]domain.Task

	r := db.WithContext(ctx).Find(&tasks)
	if r.RowsAffected != 0 {
		return tasks, nil
	}

	return nil, eris.Wrap(fmt.Errorf("no tasks found"), "DB error")
}

func (m *MySQLTaskRepo) UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	updateFields := map[string]interface{}{}
	if name != nil {
		updateFields["name"] = *name
	}

	if len(info) > 0 {
		for key, value := range info {
			updateFields[key] = value
		}
	}

	if err := tx.Model(&domain.Task{}).Where("id = ?", id).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (m *MySQLTaskRepo) GetBy(ctx context.Context, db gorm.DB, info map[string]string) ([]*domain.Task, error) {
	var tasks []*domain.Task
	var result *gorm.DB
	switch {
	case info["group"] != "":
		result = db.WithContext(ctx).Table("tasks").Where("group_name = ?", info["group"]).Scan(&tasks)
	case info["user"] != "":
		result = db.WithContext(ctx).Table("tasks").Where("user_name = ?", info["user"]).Scan(&tasks)
	case info["priority"] != "":
		result = db.WithContext(ctx).Table("tasks").Where("priority = ?", info["priority"]).Scan(&tasks)
	case info["status"] != "":
		result = db.WithContext(ctx).Table("tasks").Where("priority = ?", info["status"]).Scan(&tasks)
	}

	if result.RowsAffected != 0 {
		return tasks, nil
	}

	return nil, eris.Wrap(fmt.Errorf("no tasks found"), "DB error")
}
