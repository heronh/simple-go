package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(c *gin.Context) {

	fmt.Println("Migrate")

	//db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/your_database?sslmode=disable")
	db, err := sql.Open("postgres", "postgres://sat:U3t8QYg@192.168.30.5:5432/portal?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	coneccao := "Banco de dados connectado com sucesso"
	fmt.Println(coneccao)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	drv := "Driver carregado"
	fmt.Println(drv)

	m, err := migrate.NewWithDatabaseInstance(
		"migrations.sql", // Migration file location
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	file := "Arquivo com as migrações"
	fmt.Println(file)

	if err := m.Up(); err != nil { // Apply all pending migrations
		log.Fatal(err)
	}
	result := "Migrations successful!"
	fmt.Println(result)

	c.HTML(http.StatusOK, "migrations.html", gin.H{
		"coneccao": coneccao,
		"driver":   drv,
		"file":     file,
		"result":   result,
	})

	/*
	 */
}
