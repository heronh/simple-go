package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

type User struct {
	Username string
	Password string // This should be hashed
}

// Dummy database of users
var users = map[string]string{}

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares plain password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Register new user
func Register(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	users[newUser.Username] = hashedPassword
	c.Status(http.StatusCreated)
}

// Login user
func Login(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	hashedPassword, exists := users[user.Username]
	if !exists || !CheckPasswordHash(user.Password, hashedPassword) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

/*
 *		Servidor que responde com a página index.html e
 * envia um parâmetro, title.
 */
func Cadastro(c *gin.Context) {

	c.HTML(http.StatusOK, "cadastro.html", gin.H{
		"Heading": "Benvindo!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/cadastro", Cadastro)
	router.POST("/register", Register)
	router.POST("/login", Login)

	router.Run(":8008")
}
