package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
)

type UserController struct {
	log         *zerolog.Logger
	validator   *validator.RequestValidator
	userService service.UserService
}

func NewUserController(log *zerolog.Logger, validator *validator.RequestValidator, userService service.UserService) *UserController {
	return &UserController{
		log:         log,
		validator:   validator,
		userService: userService,
	}
}

func (u UserController) CreateUser(c echo.Context) error {
	req, err := request.BindAndValidateCreateUserRequest(c, u.validator)
	if err != nil {
		return err
	}

	res, err := u.userService.CreateOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (u UserController) GetUsers(c echo.Context) error {
	res, err := u.userService.GetMany(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u UserController) GetUser(c echo.Context) error {
	req, err := request.BindAndValidateGetUserRequest(c)
	if err != nil {
		return err
	}

	res, err := u.userService.GetOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u UserController) DeleteUser(c echo.Context) error {
	req, err := request.BindAndValidateDeleteUserRequest(c)
	if err != nil {
		return err
	}

	err := u.userService.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (u UserController) UpdateUser(c echo.Context) error {
	req, err := request.BindAndValidateUpdateUserRequest(c)
	if err != nil {
		return err
	}

	res, err := u.userService.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
