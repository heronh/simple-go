package main

import (
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
	router.LoadHTMLGlob("templates/*")

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Define a route for the root path ("/")
	router.GET("/", server.Home)
	router.GET("/pgsql", server.Pgsql)

	// Start the server
	router.Run(":4004")
}
