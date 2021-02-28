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

func UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestError) {

	currentUser, err := GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if isPatch {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}

	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	updateErr := currentUser.Update()

	if updateErr != nil {
		return nil, updateErr
	}

	return currentUser, nil

}

func DeleteUser(userId int64) *errors.RestError {

	user := &users.User{Id: userId}
	err := user.Delete()
	if err != nil {
		return err
	}

	return nil

}
