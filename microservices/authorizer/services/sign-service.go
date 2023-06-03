package services

import (
	"authorizer/models"
	"authorizer/repository"
	"authorizer/token"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

type Sign struct {
	logger         *log.Logger
	flags          *models.Flags
	signRepository *repository.Sign
}

func NewSign(l *log.Logger, f *models.Flags) *Sign {
	return &Sign{
		logger:         l,
		flags:          f,
		signRepository: repository.Init(),
	}
}

func (s *Sign) GetToken(signModel models.SignRequest, origin string) (string, *models.ErrorDetail) {
	user, err := s.signRepository.GetUserByUserName(loginModel.UserName, loginModel.Password)
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
	return token.GenerateToken(claims, expirationTime)

}

func (s *Sign) VerifyPassword(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}
