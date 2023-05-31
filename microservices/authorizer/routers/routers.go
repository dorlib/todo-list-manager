package routers

import (
	"authorizer/collectors"
	handlers "authorizer/collectors"
	"authorizer/models"
	"authorizer/token"
	"log"

	"github.com/gin-gonic/gin"
)

type Login struct {
	logger       *log.Logger
	loginHandler *collectors.Login
	flags        *models.Flags
}

func NewRoute(l *log.Logger, f *models.Flags) *Login {
	loginHandler := handlers.NewLogin(l, f)
	token.Init()

	return &Login{
		logger:       l,
		loginHandler: loginHandler,
		flags:        f,
	}
}

func (r *Login) RegisterRoutes() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.POST("/sign", r.signHandler.Sign)
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/logout", r.logoutHandler.Logout)
	ginEngine.POST("/reset", r.resetHandler.Reset)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine
}
