package impl

import (
	"fmt"
	"tasks-mgmnt/model/domain"
	"tasks-mgmnt/repo/impl"
	"tasks-mgmnt/service/utils"
)

func CreateUser(username, role, password string) (domain.User, error) {
	if utils.UserExistByName(username) {
		return domain.User{}, fmt.Errorf("user already exists")
	}
	user := domain.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	rows := impl.DB.Create(&user).RowsAffected
	fmt.Printf("rows affected: %v \n", rows)

	if rows > 0 {
		return user, nil
	}

	return domain.User{}, fmt.Errorf("failed to create user")
}
