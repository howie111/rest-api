package users

import (
	"fmt"

	"github.com/howie111/rest-api/utils/date_utils"

	"github.com/howie111/rest-api/datasources/postgres/users_db"

	"github.com/howie111/rest-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES($1,$2,$3,$4);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {

	user.DateCreated = date_utils.GetNowString()

	_, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	return nil
}

func (user *User) Get() *errors.RestError {

	err := users_db.Client.Ping()
	if err != nil {
		panic(err)
	}

	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
