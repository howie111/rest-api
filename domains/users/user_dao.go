package users

import (
	"fmt"
	"strings"

	"github.com/howie111/rest-api/logger"

	"github.com/howie111/rest-api/datasources/postgres/users_db"

	"github.com/howie111/rest-api/utils/errors"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES($1,$2,$3,$4,$5,$6);"
	queryGetUser    = "SELECT id, first_name, last_name, email,status, date_created FROM users WHERE id=$1;"
	queryUpdateUser = "UPDATE users SET first_name=$1, last_name=$2,email=$3 WHERE id=$4;"

	queryDeleteUser = "DELETE FROM users WHERE id=$1"

	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=$1;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {

	_, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if err != nil {
		logger.Error("error when trying to save user statement", err)
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
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated)

	if err != nil {
		logger.Error("error when trying to get user statement", err)
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.Id))

		}
		return errors.NewInternalServerError("database error")

	}

	return nil
}

func (user *User) Update() *errors.RestError {
	_, err := users_db.Client.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		logger.Error("error when trying to update user statement", err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to update user: %s", err.Error()))
	}

	return nil
}

func (user *User) Delete() *errors.RestError {

	_, err := users_db.Client.Exec(queryDeleteUser, user.Id)

	if err != nil {
		logger.Error("error when trying to delete user statement", err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to delete user: %s", err.Error()))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {

	users := make([]User, 0)

	rows, err := users_db.Client.Query(queryFindByStatus, status)
	if err != nil {
		logger.Error("error when trying to find user by status statement", err)
		return nil, errors.NewInternalServerError(
			fmt.Sprintf("error when trying to find user: %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			logger.Error("error when trying to scan user", err)
			return nil, errors.NewInternalServerError(
				fmt.Sprintf("error when trying to find user: %s", err.Error()))
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewInternalServerError(
			fmt.Sprintf("No users matching status: %s", status))
	}

	return users, nil

}
