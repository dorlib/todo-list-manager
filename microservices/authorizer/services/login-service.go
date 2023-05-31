package services

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"

	"authorizer/models"
	"authorizer/repository"
	"authorizer/token"
)

type Login struct {
	logger          *log.Logger
	flags           *models.Flags
	loginRepository *repository.Login
}

func NewLogin(l *log.Logger, f *models.Flags) *Login {
	return &Login{
		logger:          l,
		flags:           f,
		loginRepository: repository.Init(),
	}
}

func (l *Login) GetToken(loginModel models.LoginRequest, origin string) (string, *models.ErrorDetail) {
	user, err := l.loginRepository.GetUserByUserName(loginModel.UserName, loginModel.Password)
	if err != nil {
		return "", err
	}
	var claims = &models.JwtClaims{
		ComapnyID: strconv.Itoa(user.ID),
		Username:  user.Name,
		Roles:     user.Role,
		StandardClaims: jwt.StandardClaims{
			Audience: origin,
		},
	}
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	return token.GenrateToken(claims, expirationTime)

}

func (*Login) VerifyToken(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}
