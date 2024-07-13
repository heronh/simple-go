package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heronh/simple-go/jwt/initializers"
	"github.com/heronh/simple-go/jwt/initializers/controllers"
	"github.com/heronh/simple-go/jwt/middleware"
)

func init() {
	// Load the environment variables
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run(":4004")

}
