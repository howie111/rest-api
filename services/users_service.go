package services

import (
	"github.com/howie111/rest-api/domains/users"
	"github.com/howie111/rest-api/utils/crypto_utils"
	"github.com/howie111/rest-api/utils/date_utils"
	"github.com/howie111/rest-api/utils/errors"
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type UsersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestError)
	GetUser(id int64) (*users.User, *errors.RestError)
	UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestError)
	DeleteUser(userId int64) *errors.RestError
	SearchUser(status string) (users.Users, *errors.RestError)
}

type usersService struct {
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetDBNowDBFormat()
	user.Password = crypto_utils.GetMD5(user.Password)

	err = user.Save()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usersService) GetUser(id int64) (*users.User, *errors.RestError) {

	result := users.User{Id: id}

	err := result.Get()
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func (u *usersService) UpdateUser(isPatch bool, user users.User) (*users.User, *errors.RestError) {

	currentUser, err := u.GetUser(user.Id)

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

func (u *usersService) DeleteUser(userId int64) *errors.RestError {

	user := &users.User{Id: userId}
	err := user.Delete()
	if err != nil {
		return err
	}

	return nil

}

func (u *usersService) SearchUser(status string) (users.Users, *errors.RestError) {

	user := users.User{}

	return user.FindByStatus(status)

}
