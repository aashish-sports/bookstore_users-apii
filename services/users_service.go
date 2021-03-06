package services

import (
	"net/http"

	"github.com/aashish-sport/bookstore_users-apii/domain/users"
	"github.com/aashish-sport/bookstore_users-apii/utils/crypto_utils"
	"github.com/aashish-sport/bookstore_users-apii/utils/date_utils"
	"github.com/aashish-sport/bookstore_users-apii/utils/errors"
)

var (
	counter      int
	UsersService userService = userService{}
)

type userService struct {
}
type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	//its implemented on database
	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }
	if !isPartial {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	}

	if err := current.Update(); err != nil {
		return current, err
	}
	return current, nil

}
func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
	// users, err := dao.FindByStatus(status)
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
}
