package users

// dto means data transfer object
import (
	"strings"

	restErrors "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *restErrors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {

		return restErrors.NewBadRequestError("invalid email address")
	}
	return nil
}
