package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id          int
	Description string
	Status      bool
}

func Todo(c *gin.Context) {

	mensagens := []string{"Lista de pendencias"}
	tarefas := []Tarefa{}
	defer func() {
		c.HTML(http.StatusOK, "todo.html", gin.H{
			"Title":       "Loterias",
			"Heading":     "Tarefas!",
			"Message":     "Página de tarefas",
			"todo_active": "h5",
			"Tarefas":     tarefas,
			"Mensagens":   mensagens,
		})
	}()

	// Abre conexão com o banco de dados
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	mensagens = append(mensagens, psqlInfo)

	// Abre conexão com db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		mensagens = append(mensagens, "failed to connect database")
		mensagens = append(mensagens, err.Error())
		return
	}
	defer db.Close()

	// Busca lista de tarefas
	query := "select * from todos order by status, updated_at"
	mensagens = append(mensagens, query)
	rows, err := db.Query(query)
	if err != nil {
		mensagens = append(mensagens, "Erro ao buscar lista de tarefas")
		mensagens = append(mensagens, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t Tarefa
		err := rows.Scan(&t.Id, &t.Description, &t.Status)
		if err != nil {
			mensagens = append(mensagens, "Erro ao fazer o parse das tarefas")
			mensagens = append(mensagens, err.Error())
			return
		}
		tarefas = append(tarefas, t)
	}
	mensagens = append(mensagens, "Busca de tarefas finalizada com sucesso!")
}
