package repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"tasks-mgmnt/model/domain"
)

type UserRepo interface {
	CreateOne(ctx context.Context, db gorm.DB, group *domain.User) (*domain.User, error)
	GetOne(ctx context.Context, db gorm.DB, userID uuid.UUID) (*domain.User, error)
	DeleteOne(ctx context.Context, db gorm.DB, userID uuid.UUID) error
	GetAll(ctx context.Context, db gorm.DB) ([]*domain.User, error)
	GetBy(ctx context.Context, db gorm.DB, info string) ([]*domain.User, error)
	UpdateOne(ctx context.Context, db gorm.DB, id uuid.UUID, name *string, info map[string]interface{}) error
}
