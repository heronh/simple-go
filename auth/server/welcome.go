package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Welcome(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"Title":   "bem-vindo",
		"Heading": "Página de acesso!",
		"Message": "",
		"welcome": "h5",
	})
}
