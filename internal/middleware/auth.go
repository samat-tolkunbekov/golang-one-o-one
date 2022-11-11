package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Query("token")

		fmt.Println("Entered auth")

		if os.Getenv("APP_TOKEN") == token {
			context.Next()

			return
		}

		context.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
			"status": "token not matched",
		})

		return
	}
}
