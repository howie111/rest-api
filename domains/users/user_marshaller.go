package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	var privateUser PrivateUser
	userJson, _ := json.Marshal(user)

	json.Unmarshal(userJson, &privateUser)

	return privateUser
}

func (users Users) Marshall(isPublic bool) interface{} {
	results := make([]interface{}, len(users))

	for i, user := range users {
		results[i] = user.Marshall(isPublic)
	}
	return results
}
