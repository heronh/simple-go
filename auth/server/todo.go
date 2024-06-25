package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id          int
	Description string
	Status      bool
	Created_at  time.Time
	Created_fmt string
	Updated_at  time.Time
	Updated_fmt string
}

const YYYYMMDD = "2006-01-02"

func SaveTodo(c *gin.Context) {

	// Abre conexão com o banco de dados
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("failed to connect database")
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	description := c.PostForm("description")
	fmt.Println(description)
	current := time.Now()
	ft := current.Format("2006-01-02 15:04:05")
	query := fmt.Sprintf("INSERT INTO todos (description, status, created_at, updated_at) VALUES ('%s', false, '%s', '%s');", description, ft, ft)
	fmt.Println(query)

	// Busca lista de usuários
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	//fmt.Println(rows)
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
	query := "select * from todos order by status asc, updated_at desc"
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
		err := rows.Scan(&t.Id, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
		if err != nil {
			mensagens = append(mensagens, "Erro ao fazer o parse das tarefas")
			mensagens = append(mensagens, err.Error())
			return
		}
		tarefas = append(tarefas, t)
	}

	for i := range tarefas {
		tarefas[i].Updated_fmt = tarefas[i].Updated_at.Format("02/01/2006 15:04")
		tarefas[i].Created_fmt = tarefas[i].Created_at.Format("02/01/2006 15:04")
	}
	mensagens = append(mensagens, "Busca de tarefas finalizada com sucesso!")
}
