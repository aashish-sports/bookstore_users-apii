package users

import (
	"strings"

	"github.com/aashish-sport/bookstore_users-apii/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status`
	Password    string `json:"password"`
}


type Users []User

//function

// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("Invalid email request")
// 	}
// 	return nil
// }

//method

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email request")
	}
	return nil
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewInternalServerError("invalid passwords")
	}
	return nil
}
