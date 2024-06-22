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
	host     = "localhost" //"0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "loterias"
)

type users struct {
	id    int
	name  string
	email string
}

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func Pgsql(c *gin.Context) {
	// fmt.Println("postgres")
	mensagens := []string{"Iniciando migrações..."}
	connection := connect()
	mensagens = append(mensagens, connection)

	// query := "select id, name, email from users order by id"
	// fmt.Printf("%s\n", query)
	// list := exec_query(query)
	// for _, user := range list {
	// 	fmt.Printf("%02d - %30s %s\n", user.id, user.name, user.email)
	// }

	c.HTML(http.StatusOK, "postgres.html", gin.H{
		"title":       "Postgres",
		"psql_active": "h5",
		"Mensagens":   mensagens,
	})

}

func exec_query(query string) []users {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//fmt.Printf("query: %s\n", psqlInfo)

	// Abre conexão com db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Busca lista de usuários
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var elements []users
	for rows.Next() {
		var element users
		err := rows.Scan(&element.id, &element.name, &element.email) // match these with the struct fields
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, element)
	}
	return elements
}

// Executa conexão com DB
func connect() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
