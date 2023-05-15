package users // this is the users_controllers.go file inside the controllers folder
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/domain/users"
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/services"
	restErrors "github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := restErrors.NewBadRequestError("invalid json body")
		c.JSON(restError.StatusCode, restError)
		return
	}
	result, saveerr := services.CreateUser(user)
	if saveerr != nil {
		fmt.Println(" an error ocurred when creating userr", saveerr.Message)
		return
	}
	c.JSON(http.StatusCreated, result)
	fmt.Println((result))
	//c.String(http.StatusNotImplemented, "implement me!")

}

func GetUser(c *gin.Context) {
	userid, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := restErrors.NewBadRequestError("invalid user id")
		c.JSON(err.StatusCode, err)
	}
	user, getErr := services.GetUser(userid)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
	}

	c.JSON(http.StatusOK, user)

}

// func SearchUser(c *gin.Context) {

// 	c.String(http.StatusNotImplemented, "implement me!")

// }

func UpdateUser(c *gin.Context) {
	//id := c.Param("user_id")
	type updatePayload struct {
		UserId string `json:"id" uri:"id" binding:"required" form:"id" query:"id"` // `uri:"user_id" binding:"required"
	}

	var payload updatePayload
	if err := c.ShouldBindUri(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	fmt.Println(" --------->c.Param", payload.UserId)
	userid, userErr := strconv.ParseInt(c.Query("id"), 10, 64)

	if userErr != nil {
		err := restErrors.NewBadRequestError("invalid user id")
		c.JSON(err.StatusCode, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := restErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	user.Id = userid
	result, updateErr := services.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.StatusCode, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
