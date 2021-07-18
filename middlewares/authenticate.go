package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	SecretKey = []byte("arthur")
)

func ParseToken() gin.HandlerFunc {
	return func(context *gin.Context) {

		auth := context.Request.Header.Get("Authorization")
		if auth == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is not exist", "status": http.StatusUnauthorized})
			return
		}

		authToken := strings.Split(auth, "Bearer ")[1]

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, e error) {
			return SecretKey, nil
		})
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "status": http.StatusUnauthorized})
			context.Abort()
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "parse claim error", "status": http.StatusInternalServerError})
			context.Abort()
			return
		}

		context.Set("id", claim["id"])
		context.Set("account", claim["account"])
		context.Set("name", claim["name"])
		context.Next()
	}
}
