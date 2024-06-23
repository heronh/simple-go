package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id          int
	Description string
	Status      bool
}

// Data Structure
type Item struct {
	Name  string
	Price float64
}

func Todo(c *gin.Context) {

	// Sample data
	fruits := []Item{
		{"Apple", 0.99},
		{"Banana", 0.55},
		{"Orange", 1.25},
	}

	tarefas := []Tarefa{
		{1, "Tarefa 1", true},
		{2, "Tarefa 2", true},
		{3, "Tarefa 3", false},
		{4, "Tarefa 4", false},
		{5, "Tarefa 5", true},
	}

	c.HTML(http.StatusOK, "todo.html", gin.H{
		"Title":       "Loterias",
		"Heading":     "Tarefas!",
		"Message":     "Página de tarefas",
		"todo_active": "h5",
		"Tarefas":     tarefas,
		"Fruits":      fruits,
	})
}
