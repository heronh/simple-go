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

	mensagens := []string{"Iniciando migrações..."}

	runOnce := true
	for runOnce {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		mensagens = append(mensagens, psqlInfo)

		// Abre conexão com db
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(err)
			mensagens = append(mensagens, "Erro ao tentar conectar com DB")
			break
		}
		defer db.Close()
		mensagens = append(mensagens, "Conexão aberta")

		err = db.Ping()
		if err != nil {
			mensagens = append(mensagens, "Error pinging database")
		}
		mensagens = append(mensagens, "Ping executado com sucesso")

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
