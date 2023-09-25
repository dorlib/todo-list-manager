package service

import (
	"github.com/google/uuid"
	"tasks-mgmnt/model/request"
)

type UserService interface {
	CreateUser(context.Context, *ctx.RequestContext, *request.CreateUserRequest) (*response.UserDto, error)
	GetUsers(context.Context, *ctx.RequestContext) ([]*response.UserDto, error)
	GetUser(context.Context, *ctx.RequestContext, uuid.UUID) (*response.UserDto, error)
	DeleteUser(context.Context, *ctx.RequestContext, uuid.UUID) (*response.UserDto, error)
	UpdateUser(context.Context, *ctx.RequestContext, *request.UpdateUserRequest) (*response.UserDto, error)
}
