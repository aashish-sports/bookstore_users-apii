package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	// FirstName   string `json:"first_name"`
	// LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status`
	// Password    string `json:"password"`
}
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status`
	// Password    string `json:"password"`
}

func (user Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(user))
	for index, user := range user {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}
