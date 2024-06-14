package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(c *gin.Context) {

	fmt.Println("Migrate")

	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/your_database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"migrations.sql", // Migration file location
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil { // Apply all pending migrations
		log.Fatal(err)
	}
	fmt.Println("Migrations successful!")
}
