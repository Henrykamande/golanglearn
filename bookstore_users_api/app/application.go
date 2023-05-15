package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Startapplication() {

	mapUrls(router) // call the function from app\url_mappings.go

	router.Run(":8080")
}
