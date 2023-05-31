package routers

import (
	handlers "authorizer/collectors"
	"github.com/dorlib/todo-list-mananger/microservices/authorizer/collectors"
	"github.com/dorlib/todo-list-mananger/microservices/authorizer/models"
	"github.com/dorlib/todo-list-mananger/microservices/authorizer/token"
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
	ginEngine.POST("/login", r.loginHandler.Login)
	ginEngine.POST("/verifyToken", r.loginHandler.VerifyToken)
	return ginEngine
}
