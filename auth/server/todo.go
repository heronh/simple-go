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

const psqlQuery = "host=localhost port=5432 user=postgres password=mysecretpassword dbname=loterias sslmode=disable"
const listTasks = "select * from todos order by status asc, updated_at desc"
const deleteById = "delete from todos where id = "
const setStatusById = "update todos set status = true where id = "
const clearStatusById = "update todos set status = false where id = "

func EditTodo(c *gin.Context) {

	fmt.Println("Editando tarefa")
	// Get all form values as a map
	id := c.PostForm("id")
	formValues := c.Request.PostForm

	// Iterate and print key-value pairs
	for key, values := range formValues {
		fmt.Printf("Key: %s, Values: %v\n", key, values)
		if key == "check" {
			check(id)
		}
		if key == "uncheck" {
			uncheck(id)
		}
		if key == "delete" {
			delete(id)
		}
	}
}

func check(id string) {
	fmt.Printf("Marcar tarefa %s\n", id)
	db, _ := sql.Open("postgres", psqlQuery)
	defer db.Close()

	// Busca lista de tarefas
	query := fmt.Sprintf("%s %s", setStatusById, id)
	rows, _ := db.Query(query)
	rows.Close()
}

func uncheck(id string) {
	fmt.Printf("DesMarcar tarefa %s\n", id)
	db, _ := sql.Open("postgres", psqlQuery)
	defer db.Close()

	// Busca lista de tarefas
	query := fmt.Sprintf("%s %s", clearStatusById, id)
	rows, _ := db.Query(query)
	rows.Close()
}

func delete(id string) {
	fmt.Printf("Apagar tarefa %s\n", id)
	db, _ := sql.Open("postgres", psqlQuery)
	defer db.Close()

	// Busca lista de tarefas
	query := fmt.Sprintf("%s %s", deleteById, id)
	rows, _ := db.Query(query)
	rows.Close()
}

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

func TodoOrigina(c *gin.Context) {

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

func Todo(c *gin.Context) {

	// Abre conexão com o banco de dados
	db, _ := sql.Open("postgres", psqlQuery)
	defer db.Close()

	// Busca lista de tarefas
	rows, _ := db.Query(listTasks)
	defer rows.Close()

	tarefas := []Tarefa{}
	for rows.Next() {
		var t Tarefa
		_ = rows.Scan(&t.Id, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
		tarefas = append(tarefas, t)
	}

	for i := range tarefas {
		tarefas[i].Updated_fmt = tarefas[i].Updated_at.Format("02/01/2006 15:04")
		tarefas[i].Created_fmt = tarefas[i].Created_at.Format("02/01/2006 15:04")
	}

	c.HTML(http.StatusOK, "todo.html", gin.H{
		"Title":       "Loterias",
		"Heading":     "Tarefas!",
		"Message":     "Página de tarefas",
		"todo_active": "h5",
		"Tarefas":     tarefas,
		"Mensagens":   nil,
	})
}
