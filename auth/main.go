package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heronh/simple-go/auth/server"
)

/*
 *		Cria um servidor que só tem um path implementado, /
 */
func main() {

	// Create a Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Define a route for the root path ("/")
	router.GET("/", server.Home)
	router.GET("/migrate", server.Migrate)

	// Start the server
	router.Run(":4004")
}
