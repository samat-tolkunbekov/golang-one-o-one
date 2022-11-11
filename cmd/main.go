package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-one-o-one/helpers"
	"golang-one-o-one/internal/middleware"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Let's start coding")
	fmt.Println(os.Getenv("APP_PORT"))

	r := gin.Default()

	partners := r.Group("/partners")
	{
		partners.Use(middleware.CheckToken())

		partners.GET(":partnerId", func(c *gin.Context) {
			helpers.GetPartners(c)
		})

		partners.POST("/form_post", func(c *gin.Context) {
			helpers.PostPartners(c)
		})
	}

	r.Run(":" + os.Getenv("APP_PORT"))
}
