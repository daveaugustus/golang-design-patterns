package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userGroup := router.Group("users")

	router.Run(":8000")

	fmt.Println(userGroup)
}
