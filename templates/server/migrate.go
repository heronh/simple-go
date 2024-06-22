package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate(c *gin.Context) {

	mensagens := []string{"Iniciando migrações..."}

	runOnce := true
	for runOnce {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		mensagens = append(mensagens, psqlInfo)

		// Abre conexão com db
		db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=loterias password=mysecretpassword sslmode=disable")
		if err != nil {
			mensagens = append(mensagens, "failed to connect database")
			mensagens = append(mensagens, err.Error())
			break
		}
		mensagens = append(mensagens, "Banco de dados conectado!")
		defer db.Close()

		// Initialize the driver for PostgreSQL
		driver, _ := postgres.WithInstance(db.DB(), &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			"file:///Users/heronhurpia/Sites/simple-go/templates/migrations",
			"postgres", driver)
		if err != nil {
			mensagens = append(mensagens, "Falha ao executar migração")
			mensagens = append(mensagens, err.Error())
			break
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
				mensagens = append(mensagens, "No migration changes were applied")
				break
			}
			mensagens = append(mensagens, "Falha ao aplicar migração")
			mensagens = append(mensagens, err.Error())
			break
		}

		runOnce = false
	}

	c.HTML(http.StatusOK, "migrations.html", gin.H{
		"Title":          "Migrações",
		"Heading":        "Teste do comunicação!",
		"Message":        "",
		"migrate_active": "h5",
		"Mensagens":      mensagens,
	})
}
