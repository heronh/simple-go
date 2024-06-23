package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Todo(c *gin.Context) {

	c.HTML(http.StatusOK, "todo.html", gin.H{
		"Title":       "Loterias",
		"Heading":     "Tarefas!",
		"Message":     "Página de tarefas",
		"todo_active": "h5",
	})
}
