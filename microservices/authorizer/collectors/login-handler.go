package collectors

import (
	"authorizer/models"
	"authorizer/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	logger       *log.Logger
	flags        *models.Flags
	loginService *services.Login
}

func NewLogin(l *log.Logger, f *models.Flags) *Login {
	loginService := services.NewLogin(l, f)
	return &Login{
		logger:       l,
		flags:        f,
		loginService: loginService,
	}
}

func (l *Login) Login(ctx *gin.Context) {
	var loginObj models.LoginRequest

	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)

		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})

		badRequest(ctx, http.StatusBadRequest, "invalid request", errors)

		return
	}

	tokeString, err := l.loginService.GetToken(loginObj, ctx.Request.Header.Get("Referer"))

	if err != nil {
		badRequest(ctx, http.StatusBadRequest, "error in gerating token", []models.ErrorDetail{
			*err,
		})
		return
	}

	ok(ctx, http.StatusOK, "token created", tokeString)
}

func (l *Login) VerifyToken(context *gin.Context) {
	tokenString := context.Request.Header.Get("apikey")
	referer := context.Request.Header.Get("Referer")

	valid, claims := l.loginService.VerifyToken(tokenString, referer)
	if !valid {
		returnUnauthorized(context)
		return
	}
	ok(context, http.StatusOK, "token is valid", claims)
}
