package services

import (
	"github.com/howie111/rest-api/domains/users"
	"github.com/howie111/rest-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	err = user.Save()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(id int64) (*users.User, *errors.RestError) {

	result := users.User{Id: id}

	err := result.Get()
	if err != nil {
		return nil, err
	}

	return &result, nil

}
