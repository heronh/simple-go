package main

import (
	"net/http"

	"com.github/heronh/simple-go/postgres/server"
	"github.com/gin-gonic/gin"
)

/*
 *		Cria um servidor que só tem um path implementado, /
 */
func main() {

	// Create a Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Define a route for the root path ("/")
	//router.GET("/", server.Home)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Title":   "My Website",
			"Heading": "Welcome!",
			"Message": "This is the main content.",
		})
	})
	router.GET("/pgsql", server.Pgsql)
	router.GET("/migrate", server.Migrate)

	// Start the server
	router.Run(":4004")
}
