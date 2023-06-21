package handlers

import (
	"errors"
	"github.com/shadowshot-x/micro-product-go/authservice/data"
)

// searches the user in the database.
func validateUser(email string, passwordHash string) (bool, error) {
	usr, exists := data.GetUserObject(email)
	if !exists {

		return false, errors.New("user does not exist")
	}

	passwordCheck := usr.ValidatePasswordHash(passwordHash)

	if !passwordCheck {

		return false, nil
	}

	return true, nil
}
