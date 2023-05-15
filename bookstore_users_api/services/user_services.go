package services

import (
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/domain/users"

	restErrors "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *restErrors.RestErr) {
	if err := user.Validate(); err != nil {
		//fmt.Println(err.Error)
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user_Id int64) (*users.User, *restErrors.RestErr) {

	result := &users.User{Id: user_Id}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil

}

func UpdateUser(user users.User) (*users.User, *restErrors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

}
