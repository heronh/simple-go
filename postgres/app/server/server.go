package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func Server(c *gin.Context) {
	fmt.Println("server")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello",
	})
}
