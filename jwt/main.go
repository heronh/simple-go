package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heronh/simple-go/jwt/controllers"
	"github.com/heronh/simple-go/jwt/initializers"
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

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Serve static files (CSS) from the 'static' directory
	r.Static("/static", "./static")

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"Title":   "Benvindo",
			"Heading": "Página de acesso!",
			"Message": "",
			"welcome": "h5",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"Title":   "Benvindo",
			"Heading": "Página de acesso!",
			"Message": "",
			"welcome": "h5",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title":   "Login",
			"Heading": "Página de login!",
			"Message": "",
			"login":   "h5",
		})
	})
	r.POST("/login", controllers.Login, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title":   "Login",
			"Heading": "Página de login!",
			"Message": "",
			"login":   "h5",
		})
	})

	r.POST("/signup", controllers.Signup)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run(":4004")

}
