package controller

import (
	Ctx "github.com/dorlib/todo-list-manager/libs/microservice-development-kit/ctx"
	"github.com/dorlib/todo-list-manager/libs/microservice-development-kit/validator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
	"tasks-mgmnt/model/request"
	"tasks-mgmnt/service"
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
	req, err := request.BindAndValidateCreateRequest(c, u.validator)
	if err != nil {
		return err
	}

	res, err := u.userService.CreateUser(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (u UserController) GetUsers(c echo.Context) error {
	res, err := u.userService.GetUsers(c.Request().Context(), Ctx.FromEchoContext(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u UserController) GetUser(c echo.Context) error {
	req, err := request.BindAndValidateGetRequest(c, u.validator)
	if err != nil {
		return err
	}

	res, err := u.userService.GetUser(c.Request().Context(), Ctx.FromEchoContext(c), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u UserController) DeleteUser(c echo.Context) error {
	req, err := request.BindAndValidateDeleteRequest(c, u.validator)
	if err != nil {
		return err
	}

	res, err := u.userService.DeleteUser(c.Request().Context(), Ctx.FromEchoContext(c), req.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (u UserController) UpdateUser(c echo.Context) error {
	req, err := request.BindAndValidateUpdateRequest(c, u.validator)
	if err != nil {
		return err
	}

	res, err := u.userService.UpdateUser(c.Request().Context(), Ctx.FromEchoContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
