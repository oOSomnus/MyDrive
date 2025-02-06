package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oOSomnus/MyDrive/db"
	"github.com/oOSomnus/MyDrive/internal/user/handler"
	"github.com/oOSomnus/MyDrive/internal/user/repository"
	"github.com/oOSomnus/MyDrive/internal/user/service"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env.base", ".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env.base file: %s", err.Error())
	}

	DBManager := db.GetDBManager(
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
	)
	defer DBManager.Close()
	userRepo := repository.NewUserRepository(DBManager.GetDB())
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery(), cors.Default())
	r.GET(
		"/ping", func(c *gin.Context) {
			c.JSON(
				200, gin.H{
					"message": "pong",
				},
			)
		},
	)
	r.POST("/register", userHandler.CreateUser)
	r.POST("/login", userHandler.Authenticate)
	r.Run(":8080")
}
