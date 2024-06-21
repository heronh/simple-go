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

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Title":   "My Website",
		"Heading": "Welcome!",
		"Message": "This is the main content.",
	})
}
