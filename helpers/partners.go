package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPartners(c *gin.Context) {
	fmt.Println(c.Param("partnerId"))

	c.JSON(http.StatusOK, gin.H{
		"message": "trong",
	})

	return
}

func PostPartners(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})

	return
}
