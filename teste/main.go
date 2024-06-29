package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// App configuration (replace with your values)
const (
	SecretKey   = "your_very_secret_key" // Use a strong secret in production!
	TokenExpiry = 1 * time.Minute        // Token validity duration
	Port        = 8008
)

// User struct (replace withyour database model)
type User struct {
	ID       int
	Username string
	Password string
}

// Sample user data (replace with database access)
var users = []User{
	{1, "alice", "password"},
	{2, "bob", "password"},
}

// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Kind string
}

var (
	signKey *rsa.PrivateKey
)

// CustomClaims includes jwt.StandardClaims and additional custom fields
type CustomClaims struct {
	jwt.RegisteredClaims
	CustomerInfo
}

func createToken(user string) (string, error) {
	// Define the standard claims
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiry)),
		},
		CustomerInfo: CustomerInfo{Name: user, Kind: "customer"},
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with our private key
	signKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func Login(ctx *gin.Context) {

	// Get the username and password from the request
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// Check if the user exists and the password is correct
	var user User
	for _, u := range users {
		if u.Username == username && u.Password == password {
			user = u
			fmt.Println("User found")
			break
		}
	}

	// If the user was not found or the password is incorrect
	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Create a token
	token, err := createToken(user.Username)
	if err != nil {
		fmt.Println("Error creating token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating token"})
		return
	}
	fmt.Println(token)

	// Return the token
	ctx.HTML(http.StatusOK, "produtos.html", gin.H{
		"token": token,
	})
}

func main() {

	// Inicia gin
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// página de cadastro
	router.GET("/produtos", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "produtos.html", gin.H{})
	})
	// página de cadastro
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.POST("/login", Login)

	router.Run(":8008")
}
