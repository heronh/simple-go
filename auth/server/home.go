package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":       "Templates",
		"Heading":     "Benvindo!",
		"Message":     "Página principal",
		"home_active": "h5",
	})
}
