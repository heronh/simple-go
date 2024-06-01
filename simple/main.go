package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func server(c *gin.Context) {

	fmt.Println("server")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello",
	})
}

func main() {

	absPath, _ := filepath.Abs("templates/**/*")
	log.Println(absPath)

	// // Create a Gin router
	router := gin.Default()

	// // Load HTML templates (assuming "index.html" is in the "templates" directory)
	router.LoadHTMLGlob("templates/*")

	// // Define a route for the root path ("/")
	router.GET("/", server)

	// // Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// // Define a route for "/ping"
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// // Start the server
	router.Run(":8080")

}
