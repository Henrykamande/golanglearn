package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	restErrors "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *restErrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return restErrors.NewInternalServerError("no record matching given id")
		}
		return restErrors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return restErrors.NewBadRequestError("Duplicated Date")

	}
	return restErrors.NewInternalServerError("error processing request")
}
