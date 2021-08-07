package controllers

import (
	"net/http"
	"strings"
	"time"

	"MemberSystem/database"
	"MemberSystem/middlewares"
	"MemberSystem/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "registration information is not complete"})
		return
	}
	user.Password, _ = hashPassword(user.Password)

	dbErr := user.Insert(database.DB)
	if dbErr != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "message": "register is not succes", "content": dbErr.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "member created successfully"})
}

func Login(context *gin.Context) {

	var user models.User
	var pwFromDB string
	var userName string
	var id int

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Login information is not complete"})
		return
	}

	auth := context.Request.Header.Get("Authorization")
	if auth != "" {
		authToken := strings.Split(auth, "Bearer ")[1]

		_, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, e error) {
			return middlewares.SecretKey, nil
		})
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token", "status": http.StatusUnauthorized})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "AutoLogin successfully"})
		return
	}

	row := database.DB.Table("users").Where("account = ?", user.Account).Select("password, name, id").Row()
	row.Scan(&pwFromDB, &userName, &id)
	verify := verifyPassword(user.Password, pwFromDB)
	if verify {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["id"] = id
		claims["account"] = user.Account
		claims["name"] = userName
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(middlewares.SecretKey))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "JWT token failed"})
		}
		context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "login successfully", "token": tokenString})
		return
	}

	context.JSON(http.StatusUnauthorized, gin.H{"message": "login failed", "status": http.StatusUnauthorized})
}

func hashPassword(pw string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	return string(bytes), err
}

func verifyPassword(pw, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
