package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=loterias password=mysecretpassword sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Initialize the driver for PostgreSQL
	driver, _ := postgres.WithInstance(db.DB(), &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///Users/heronhurpia/Sites/simple-go/migrate/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	// if err := m.Down(); err != nil {
	// 	if err == migrate.ErrNoChange {
	// 		fmt.Println("No migration changes were applied")
	// 		return
	// 	}
	// 	panic(err)
	// }

	// Migrate all the way up
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migration changes were applied")
			return
		}
		panic(err)
	}

	fmt.Println("Migration successful")
}
