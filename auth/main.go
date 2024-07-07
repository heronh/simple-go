package main

import (
	"fmt"
	"jwt-authentication/teste/server"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var claims jwt.MapClaims

func main() {

	// Create a Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Serve static files (CSS) from the 'static' directory
	router.Static("/static", "./static")

	// Define a route for the root path ("/")
	router.GET("/", server.Home)
	router.GET("/login", server.Login)
	router.GET("/tarefas", server.Todo)
	router.POST("/login", func(c *gin.Context) {
		if auth(c) {
			server.Welcome(c)
		} else {
			server.Login(c)
		}
	})

	// Start the server
	router.Run(":4004")
}

func auth(c *gin.Context) bool {

	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)
	if password != "123" {
		return false
	}

	// Dados da carga útil (claims)
	claims = jwt.MapClaims{
		"email":    email,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas
	}

	// Chave secreta (mantenha isso em segredo!)
	secretKey := []byte("sua_chave_secreta")

	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Token:", tokenString)
	return true
}
