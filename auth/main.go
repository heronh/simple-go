package main

import (
	"net/http"

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
	router.GET("/todo", server.Todo)

	// Salva nova tarefa e retorna a página de tarefas
	router.POST("/save-todo", func(context *gin.Context) {
		server.SaveTodo(context)
		context.Redirect(http.StatusFound, "/todo")
	})

	// Edita tarefa e retorna a página de tarefas
	router.POST("/edit-todo", func(context *gin.Context) {
		server.EditTodo(context)
		context.Redirect(http.StatusFound, "/todo")
	})

	// Start the server
	router.Run(":4004")
}
