package controllers

import (
	"net/http"

	"MemberSystem/database"
	"MemberSystem/models"

	"github.com/gin-gonic/gin"
)

func UpdateUser(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "json error"})
		return
	}

	account, flag := context.Get("account")
	if !flag {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "account is not found"})
		return
	}

	dbErr := user.Update(database.DB, account)
	if dbErr != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "message": "update is not succes", "content": dbErr.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update successfully"})
}

func GetUser(context *gin.Context) {

	var user models.User

	account, flag := context.Get("account")
	if !flag {
		context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "account is not found"})
		return
	}

	dbErr := user.FindOne(database.DB, account)
	if dbErr != nil {
		context.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "message": "get user is not success", "content": dbErr.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "get information successfully", "content": user})
}
