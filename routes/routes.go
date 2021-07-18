package routes

import (
	"MemberSystem/controllers"
	"MemberSystem/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()
	userGroup := router.Group("/api/v1/users").Use(middlewares.ParseToken())
	{
		userGroup.GET("/profile", controllers.GetUser)
		userGroup.PUT("/profile", controllers.UpdateUser)
	}

	router.POST("/api/v1/register", controllers.CreateUser)
	router.POST("/api/v1/login", controllers.Login)

	router.Run()
}
