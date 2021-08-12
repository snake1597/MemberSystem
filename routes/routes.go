package routes

import (
	"MemberSystem/controllers"
	"MemberSystem/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	router := gin.Default()

	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	userGroup := router.Group("/api/v1/users").Use(middlewares.ParseToken())
	{
		userGroup.GET("/profile", controllers.GetUser)
		userGroup.PUT("/profile", controllers.UpdateUser)
	}

	router.Run()
}
