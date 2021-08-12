package middlewares

import (
	"MemberSystem/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		if auth == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "token is missing"})
			return
		}

		token := strings.Split(auth, "Bearer ")[1]
		tokenInfo, err := services.ValidateToken(token)
		if err == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "token is invalid"})
			return
		}
		context.Set("id", tokenInfo["id"])
		context.Set("account", tokenInfo["account"])
		context.Next()
	}
}
