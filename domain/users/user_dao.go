package users

import (
	"fmt"

	"github.com/aashish-sport/bookstore_users-apii/datasources/mysql/users_db"
	"github.com/aashish-sport/bookstore_users-apii/utils/date_utils"
	"github.com/aashish-sport/bookstore_users-apii/utils/errors"
	"github.com/aashish-sport/bookstore_users-apii/utils/mysql_utils"
)

// var (
// 	usersDB = make(map[int64]*User)
// )

const (
	queryInsertUser       = ("Insert Into users(first_name,last_name,email,date_created,password,status)VALUES(?,?,?,?,?,?);")
	queryGetUser          = ("SELECT id, first_name,last_name,email,date_created FROM users WHERE id=?")
	queryUpdateUser       = "UPDATE users SET first_name=?,last_name=?,email=?WHERE id=?; "
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?;"
)

// we need to work with pointer because we need actuall value not a copy
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	results, _ := stmt.Query(user.Id)
	if err != nil {

		return errors.NewInternalServerError(err.Error())
	}
	defer results.Close()
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
		// sqlErr, ok := getErr.(*mysql.MySQLError)
		// if !ok {
		// 	return errors.NewInternalServerError(
		// 		fmt.Sprintf("error when trying to save user:%s", getErr.Error()))
		// }
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// if strings.Contains(err.Error(), "errorNoRows") {
		// 	return errors.NewNotFoundError(fmt.Sprintf("user %d does not exists.", user.Id))
		// }

		// fmt.Println(err)
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d :%s", user.Id, getErr.Error()))
	}
	// for local
	// result := usersDB[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	interResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if saveErr != nil {
		//converting err into mysql error
		return mysql_utils.ParseError(saveErr)
		// sqlErr, ok := saveErr.(*mysql.MySQLError)
		// if !ok {
		// 	return errors.NewInternalServerError(
		// 		fmt.Sprintf("error when trying to save user:%s", saveErr.Error()))
		// }
		// return errors.NewInternalServerError(
		// 	fmt.Sprintf("error when trying to save user:%s", saveErr.Error()))
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// switch sqlErr.Number {
		// case 1062:
		// 	return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user:%s", saveErr.Error()))

		// if strings.Contains(err.Error(), "email_UNIQUE") {
		// 	return errors.NewBadRequestError(fmt.Sprintf("email %s aleady exists", err.Error()))
		// }
		// return errors.NewInternalServerError(
		// 	fmt.Sprintf("error when trying to save user:%s", err.Error()))
	}
	// result, err:=users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.DateCreated)
	userId, err := interResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
		// return errors.NewInternalServerError(err.Error())
	}
	user.Id = userId

	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError("This email already registered")
	// 	}
	// 	return errors.NewBadRequestError("user %d already")
	// }

	// user.DateCreated = date_utils.GetNowString()
	// usersDB[user.Id] = user
	return nil
}
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	_, err = stmt.Exec(user.Id)
	defer stmt.Close()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
		if len(result) == 0 {
			return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
		}
	}
	return result, nil
}
