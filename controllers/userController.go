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

func CreateUser(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "registration information is not complete"})
		return
	}
	user.Password, _ = hashPassword(user.Password)
	result := database.DB.Create(&user)
	if result.Error != nil {
		println("Create failt")
	}
	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "member created successfully"})
}

func UpdateUser(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Login information is not complete"})
		return
	}

	account, flag := context.Get("account")
	if !flag {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "account is not found"})
		return
	}

	database.DB.Table("users").Where("account = ?", account).Updates(map[string]interface{}{"name": user.Name, "birthday": user.Birthday})
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Update successfully"})
}

func GetUser(context *gin.Context) {

	var user models.User

	account, flag := context.Get("account")
	if !flag {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "account is not found"})
		return
	}

	row := database.DB.Table("users").Where("account = ?", account).Select("name, birthday").Row()
	row.Scan(&user.Name, &user.Birthday)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Get information successfully", "content": user})
}

type LoginStruct struct {
	Account  string `gorm:"account" json:"account" binding:"required"`
	Password string `gorm:"password" json:"password" binding:"required"`
}

func Login(context *gin.Context) {

	var login LoginStruct
	var pwFromDB string
	var userName string
	var id int

	err := context.ShouldBindJSON(&login)
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

	row := database.DB.Table("users").Where("account = ?", login.Account).Select("password, name, id").Row()
	row.Scan(&pwFromDB, &userName, &id)
	verify := verifyPassword(login.Password, pwFromDB)
	if verify {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["id"] = id
		claims["account"] = login.Account
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
