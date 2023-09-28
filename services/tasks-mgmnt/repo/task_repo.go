package repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"
)

type TaskRepo interface {
	CreateOne(ctx context.Context, db gorm.DB, group *domain.Task) (*domain.Task, error)
	GetOne(ctx context.Context, db gorm.DB, taskID uuid.UUID) (*domain.Task, error)
	DeleteOne(ctx context.Context, db gorm.DB, taskID uuid.UUID) error
	GetAll(ctx context.Context, db gorm.DB) ([]*domain.Task, error)
	GetBy(ctx context.Context, db gorm.DB, info string) ([]*domain.Task, error)
	UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error
}
