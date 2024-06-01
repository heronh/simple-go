package main

import (
	"fmt"
	"net/http"

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

	// Define a route for the root path ("/")
	router.GET("/", server)

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Start the server
	router.Run(":8080")
}

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func server(c *gin.Context) {
	fmt.Println("server")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello",
	})
}
