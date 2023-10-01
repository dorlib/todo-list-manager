package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"
)

type MySQLUserRepo struct{}

func NewMySQLUserRepo() *MySQLUserRepo {
	return &MySQLUserRepo{}
}

func (m *MySQLUserRepo) CreateOne(ctx context.Context, db gorm.DB, user *domain.User) (*domain.User, error) {
	rows := db.WithContext(ctx).Create(&user).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return user, nil
	}

	return nil, fmt.Errorf("failed to create group")
}

func (m *MySQLUserRepo) GetOne(ctx context.Context, db gorm.DB, userID uuid.UUID) (*domain.User, error) {
	var user *domain.User

	r := db.WithContext(ctx).Where("id = ?", userID).First(&user)
	if r.RowsAffected != 0 {
		return user, nil
	}

	return nil, eris.Wrap(fmt.Errorf("user with the id %v doesnt exists", userID), "DB error")
}

func (m *MySQLUserRepo) DeleteOne(ctx context.Context, db gorm.DB, userID uuid.UUID) error {
	r := db.WithContext(ctx).Delete(&domain.User{}, userID)
	if r.RowsAffected != 0 {
		fmt.Printf("deleted user: %v \n", userID)

		return nil
	}

	return eris.Wrap(fmt.Errorf("there is no user with the given ID: %v", userID), "DB error")
}

func (m *MySQLUserRepo) GetAll(ctx context.Context, db gorm.DB) (*[]domain.User, error) {
	var users *[]domain.User

	r := db.WithContext(ctx).Find(&users)
	if r.RowsAffected != 0 {
		return users, nil
	}

	return nil, eris.Wrap(fmt.Errorf("no users found"), "DB error")
}

func (m *MySQLUserRepo) UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error {
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

	if err := tx.Model(&domain.User{}).Where("id = ?", id).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
