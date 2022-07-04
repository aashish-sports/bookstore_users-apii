package mysql_utils

import (
	"strings"

	// "github.com/aashish-sport/bookstore_users-api/utils/errors"
	"github.com/aashish-sport/bookstore_users-apii/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	//valid sql error
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")

	}
	return errors.NewInternalServerError("Error processing request")
}
