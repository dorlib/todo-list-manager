package request

import "github.com/google/uuid"

type CreateTaskRequest struct {
	Email                string            `json:"email" validate:"email,required"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type UpdateTaskRequest struct {
	ID                   uuid.UUID         `param:"id"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type GetTaskRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetTaskByRequest struct {
	ID   uuid.UUID         `param:"id"`
	Info map[string]string `json:"info"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `param:"id"`
}
