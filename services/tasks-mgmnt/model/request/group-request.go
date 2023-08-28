package request

import "github.com/google/uuid"

type CreateGroupRequest struct {
	Email                string            `json:"email" validate:"email,required"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type UpdateGroupRequest struct {
	ID                   uuid.UUID         `param:"id"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type GetGroupRequest struct {
	ID uuid.UUID `param:"id"`
}

type DeleteGroupRequest struct {
	ID uuid.UUID `param:"id"`
}
