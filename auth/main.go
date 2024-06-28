package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heronh/simple-go/auth/server"
)

func main() {

	// Simulate some private data
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	// Create a Gin router
	router := gin.Default()
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",    // user:foo password:bar
		"austin": "1234",   // user:austin password:1234
		"lena":   "hello2", // user:lena password:hello2
		"manu":   "4321",   // user:manu password:4321
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// Get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Define a route for the root path ("/")
	router.Use(auth.Required())
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
