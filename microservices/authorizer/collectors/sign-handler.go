package collectors

import (
	"authorizer/models"
	"authorizer/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type Sign struct {
	logger       *log.Logger
	flags        *models.Flags
	loginService *services.Login
}

func NewSign(l *log.Logger, f *models.Flags) *Sign {
	signService := services.NewSign(l, f)
	return &Sign{
		logger:       l,
		flags:        f,
		loginService: signService,
	}
}

func (s *Sign) Sign(ctx *gin.Context) {
	var body struct {
		UserName string
		Password String
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// create user (consider implement here and move from data).
}
