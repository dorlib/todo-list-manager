package repo

import (
	"context"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"

	"github.com/google/uuid"
)

type GroupRepo interface {
	CreateOne(ctx context.Context, db gorm.DB, group *domain.Group) (*domain.Group, error)
	GetOne(ctx context.Context, db gorm.DB, groupID uuid.UUID) (*domain.Group, error)
	DeleteOne(ctx context.Context, db gorm.DB, groupID uuid.UUID) error
	GetAll(ctx context.Context, db gorm.DB) ([]*domain.Group, error)
	UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error
}
