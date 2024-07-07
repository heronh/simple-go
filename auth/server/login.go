package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title":        "Login",
		"Heading":      "Página de acesso!",
		"Message":      "",
		"login_active": "h5",
	})
}
