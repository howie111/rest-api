package users

import (
	"strings"

	"github.com/howie111/rest-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() *errors.RestError {

	email := strings.TrimSpace(strings.ToLower(u.Email))

	if email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
