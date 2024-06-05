package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

const (
	host     = "192.168.30.5"
	port     = 5432
	user     = "sat"
	password = "U3t8QYg"
	dbname   = "portal"
)

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func Pgsql(c *gin.Context) {
	fmt.Println("postgres")
	connection := connect()
	fmt.Printf("%s\n", connection)

	c.HTML(http.StatusOK, "postgres.html", gin.H{
		"title":      "Postgres",
		"connection": connection,
	})

}

func connect() string {

	fmt.Println("Tentativa de conexão")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Printf("query: %s\n", psqlInfo)

	// Abre conexão com db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return "Erro ao tentar conectar com DB"
	}
	defer db.Close()

	// Teste de conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return "DB não respondeu ao ping"
	}

	return "Banco de dados conectado"
}
