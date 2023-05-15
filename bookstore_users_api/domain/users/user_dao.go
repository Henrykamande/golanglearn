package users

import (
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/datasource/mysql/users_db"
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/dateutils"
	restErrors "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/errors"
	mysqlutils "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/mysql_utils"
)

// dao  means data access object

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *restErrors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	statment, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer statment.Close()

	getUserResult := statment.QueryRow(user.Id)
	if getErr := getUserResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}

	return nil
}
func (user *User) Save() *restErrors.RestErr {
	statment, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer statment.Close()
	user.DateCreated = dateutils.GetNowString()
	insertResult, saveErr := statment.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)

	}
	userId, err := insertResult.LastInsertId()
	if err != nil {

		return mysqlutils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *restErrors.RestErr {

	statment, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return restErrors.NewInternalServerError(err.Error())
	}
	defer statment.Close()
	_, err = statment.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil

}
