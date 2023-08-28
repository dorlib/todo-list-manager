package request

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
)

type CreateUserRequest struct {
	Email                string            `json:"email" validate:"email,required"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type UpdateUserRequest struct {
	ID                   uuid.UUID         `param:"id"`
	Role                 string            `json:"role" validate:"required,oneof=Admin Editor"`
	AdditionalProperties map[string]string `json:"additionalProperties"`
}

type GetUserRequest struct {
	ID uuid.UUID `param:"id"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `param:"id"`
}

func BindAndValidateCreateUserRequest(c echo.Context, validator *validator.RequestValidator) (*CreateUserRequest, error) {
	var err error

	req := new(CreateUserRequest)
	if err = c.Bind(req); err != nil {
		return nil, eris.Wrap(appError.InvalidRequest(err.Error(), "binding request failed"), "failed to bind request")
	}

	if err = validator.Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func BindAndValidateUpdateUserRequest(c echo.Context) (*UpdateUserRequest, error) {
	var err error

	req := new(UpdateUserRequest)
	if err = c.Bind(req); err != nil {
		return nil, eris.Wrap(appError.InvalidRequest(err.Error(), "binding request failed"), "failed to bind request")
	}

	return req, nil
}

func BindAndValidateDeleteUserRequest(c echo.Context) (*DeleteUserRequest, error) {
	var err error

	req := new(DeleteUserRequest)
	if err = c.Bind(req); err != nil {
		return nil, eris.Wrap(appError.InvalidRequest(err.Error(), "binding request failed"), "failed to bind request")
	}

	return req, nil
}

func CreateUserRequestToClientUserInput(req *CreateUserRequest) clientmodel.CreateUserInput {
	return clientmodel.CreateUserInput{
		Email: req.Email,
		Role:  req.Role,
	}
}

func UpdateUserRequestToClientUserInput(req *UpdateUserRequest) clientmodel.UpdateUserInput {
	return clientmodel.UpdateUserInput{
		ID:   &req.ID,
		Role: req.Role,
	}
}
