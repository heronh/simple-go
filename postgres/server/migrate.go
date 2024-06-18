package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/golang-migrate/migrate"
	//"github.com/golang-migrate/migrate/database/postgres"
	//_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(c *gin.Context) {

	fmt.Println("Migrate")

	//db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/your_database?sslmode=disable")
	db, err := sql.Open("postgres", "postgres://postgres:strong_password@172.18.0.2:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	coneccao := "Banco de dados connectado com sucesso"
	fmt.Println(coneccao)

	err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging database:", err)
    }
    fmt.Println("Successfully connected to PostgreSQL database!")	

	c.HTML(http.StatusOK, "migrations.html", gin.H{
		"coneccao": coneccao,
	})
}
