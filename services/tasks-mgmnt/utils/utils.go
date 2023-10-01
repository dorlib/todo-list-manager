package utils

import (
	"github.com/google/uuid"
	"tasks-mgmnt/model/domain"
)

func IsTaskEmpty(task *domain.Task) bool {
	return task.ID == uuid.Nil
}

func IsGroupEmpty(group *domain.Group) bool {
	return group.ID == uuid.Nil
}

func IsUserEmpty(user *domain.User) bool {
	return user.ID == uuid.Nil
}
