package repository

import (
	"authorizer/models"
)

type Login struct {
}

var user []models.User

func Init() *Login {

	user = make([]models.User, 0, 2)

	user = append(user,
		models.User{
			ID:       1,
			Name:     "Steve",
			UserName: "steve@yopmail.com",
			Password: "steve123",
			Role:     "2",
		},
		models.User{
			ID:       2,
			Name:     "Mark",
			UserName: "mark@yopmail.com",
			Password: "mark@123",
			Role:     "1",
		},
	)

	return &Login{}
}

func (l *Login) GetUserByUserName(userName, password string) (models.User, *models.ErrorDetail) {
	for _, value := range user {
		if value.UserName == userName && value.Password == password {
			return value, nil
		}
	}

	return models.User{}, &models.ErrorDetail{
		ErrorType:    models.ErrorTypeError,
		ErrorMessage: "UserName and password not valid",
	}
}
