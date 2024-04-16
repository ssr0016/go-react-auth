package users

import (
	"react-auth/backend/utils/errors"
	"strings"
)

// DTO data transfer object

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (user *User) Validate() *errors.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
