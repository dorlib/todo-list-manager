package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"
)

type MySQLGroupRepo struct{}

func NewMySQLGroupRepo() *MySQLGroupRepo {
	return &MySQLGroupRepo{}
}

func (m *MySQLGroupRepo) CreateOne(ctx context.Context, db gorm.DB, group *domain.Group) (*domain.Group, error) {
	rows := db.WithContext(ctx).Create(&group).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return group, nil
	}

	return nil, eris.Wrap(fmt.Errorf("failed to create group"), "DB error")
}

func (m *MySQLGroupRepo) GetOne(ctx context.Context, db gorm.DB, groupID uuid.UUID) (*domain.Group, error) {
	var group *domain.Group

	r := db.WithContext(ctx).Where("id = ?", groupID).First(&group)
	if r.RowsAffected != 0 {
		return group, nil
	}

	return nil, eris.Wrap(fmt.Errorf("group with the id %v doesnt exists", groupID), "DB error")
}

func (m *MySQLGroupRepo) DeleteOne(ctx context.Context, db gorm.DB, groupID uuid.UUID) error {
	r := db.WithContext(ctx).Delete(&domain.Group{}, groupID)
	if r.RowsAffected != 0 {
		fmt.Printf("deleted group: %v \n", groupID)

		return nil
	}

	return eris.Wrap(fmt.Errorf("there is no group with the given ID: %v", groupID), "DB error")
}

func (m *MySQLGroupRepo) GetAll(ctx context.Context, db gorm.DB) (*[]domain.Group, error) {
	var groups *[]domain.Group

	db.WithContext(ctx).Find(&groups)
	if groups != nil {
		return groups, nil
	}

	return nil, eris.Wrap(fmt.Errorf("no groups found"), "DB error")
}

func (m *MySQLGroupRepo) UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error {
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

	if err := tx.Model(&domain.Group{}).Where("id = ?", id).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
