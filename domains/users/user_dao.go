package users

import (
	"fmt"
	"strings"

	"github.com/howie111/rest-api/utils/date_utils"

	"github.com/howie111/rest-api/datasources/postgres/users_db"

	"github.com/howie111/rest-api/utils/errors"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES($1,$2,$3,$4);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=$1;"
	queryUpdateUser = "UPDATE users SET first_name=$1, last_name=$2,email=$3 WHERE id=$4;"

	queryDeleteUser = "DELETE FROM users WHERE id=$1"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {

	user.DateCreated = date_utils.GetNowString()

	_, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "users_email_key") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	return nil
}

func (user *User) Get() *errors.RestError {

	result := users_db.Client.QueryRow(queryGetUser, user.Id)
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.Id))

		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to retrieve user: %s", err.Error()))
	}

	return nil
}

func (user *User) Update() *errors.RestError {
	_, err := users_db.Client.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to update user: %s", err.Error()))
	}

	return nil
}

func (user *User) Delete() *errors.RestError {

	_, err := users_db.Client.Exec(queryDeleteUser, user.Id)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to delete user: %s", err.Error()))
	}
	return nil
}
