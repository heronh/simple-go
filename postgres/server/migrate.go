package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Migration_t struct {
	Exists   string
	Create   string
	Populate string
}

var Migrations []Migration_t

func init() {

	Migrations = []Migration_t{
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a", "b", "c"},
	}
}

func Migrate(c *gin.Context) {
	fmt.Println("postgres")
	//connection := connect()

	for _, m := range Migrations {
		fmt.Printf("Name: %s %s, City: %s\n", m.Exists, m.Create, m.Populate)
	}
}

/*
var create_users := "CREATE TABLE users (
   user_id SERIAL PRIMARY KEY,
   name VARCHAR(50) UNIQUE NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);"
*/
