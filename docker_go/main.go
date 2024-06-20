package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	//db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/your_database?sslmode=disable")
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@127.0.0.0:5432/postgres?sslmode=disable")
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
}
