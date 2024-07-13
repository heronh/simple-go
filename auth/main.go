package main

import (
	"fmt"
	"jwt-authentication/teste/server"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("1234546")

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
	router.POST("/login", login)

	// Protected routes
	protected := router.Group("/protected")
	protected.Use(authMiddleware())
	{
		protected.GET("/tarefas", protectedData)
	}

	// Start the server
	router.Run(":4004")
}

func login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func protectedData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "This is protected data"})
}
