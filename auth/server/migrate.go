package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost" //"0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "loterias"
)

func Migrate(c *gin.Context) {

	mensagens := []string{"Iniciando migrações..."}
	defer func() {
		c.HTML(http.StatusOK, "migrations.html", gin.H{
			"Title":          "Migrações",
			"Heading":        "Teste do comunicação!",
			"Message":        "",
			"migrate_active": "h5",
			"Mensagens":      mensagens,
		})
	}()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	mensagens = append(mensagens, psqlInfo)

	// Abre conexão com db
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=loterias password=mysecretpassword sslmode=disable")
	if err != nil {
		mensagens = append(mensagens, "failed to connect database")
		mensagens = append(mensagens, err.Error())
		log.Println("")
		return
	}
	mensagens = append(mensagens, "Banco de dados conectado!")
	defer db.Close()

	// Initialize the driver for PostgreSQL
	driver, _ := postgres.WithInstance(db.DB(), &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///Users/heronhurpia/Sites/simple-go/auth/migrations",
		"postgres", driver)
	if err != nil {
		mensagens = append(mensagens, "Falha ao executar migração")
		mensagens = append(mensagens, err.Error())
		return
	}

	// Migrate all the way up
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			mensagens = append(mensagens, "No migration changes were applied")
			return
		}
		mensagens = append(mensagens, "Falha ao aplicar migração")
		mensagens = append(mensagens, err.Error())
		return
	}
	mensagens = append(mensagens, "Migração executada com sucesso!")
}
