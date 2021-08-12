package controllers

import (
	"net/http"
	"strings"

	"MemberSystem/database"
	"MemberSystem/models"
	"MemberSystem/services"
	"MemberSystem/tools"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "registration information is not complete"})
		return
	}
	user.Password, _ = tools.HashString(user.Password)

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
		token := strings.Split(auth, "Bearer ")[1]

		_, err := services.ValidateToken(token)
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
	errVerify := tools.VerifyHashString(user.Password, pwFromDB)
	if errVerify != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "message": "invalid password"})
	}

	authToken, err := services.GenerateToken(id, user.Account)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "JWT token failed"})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "login successfully", "token": authToken})
}
