package app

import (
	"github.com/gin-gonic/gin"
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/controllers/ping"
	"github.com/restApi/go/src/github.com/bookstorego/bookstore_users_api/controllers/users"
)

// mapUrls is a function that maps the urls
func mapUrls(router *gin.Engine) {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.POST("/users/update/:id,", users.UpdateUser) // call the function from controllers\ping_controller.go
	router.GET("/users/:user_id", users.GetUser)        // call the function from controllers\users_controllers.go

	//router.GET("users/search", controllers.SearchUser)

}
