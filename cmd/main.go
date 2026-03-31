package main

import (
	"github.com/Harsh00067/harshad-api/database"
	"github.com/Harsh00067/harshad-api/handlers"
	"github.com/Harsh00067/harshad-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	database.Connect()
	router := gin.Default()
	router.Use(middleware.Logger())
	router.POST("/login", handlers.Login)
	router.POST("refresh", handlers.RefreshToken)
	router.GET("/health", handlers.Health)
	router.POST("/users", handlers.CreateUser)
	//router.GET("/users/:id", handlers.GetUserByID)
	//router.GET("/users", handlers.GetUsers)
	//router.DELETE("/users/:id", handlers.DeleteUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUserByID)
	}

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.Authorize("admin"))
	{
		adminRoutes.DELETE("/users/:id", handlers.DeleteUser)
	}

	router.Run(":8080")

}
