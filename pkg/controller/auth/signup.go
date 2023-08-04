package auth

import (
	"log"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/config/err"
	"social-network-go/pkg/config/validators"
	"social-network-go/pkg/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	username := strings.TrimSpace(c.PostForm("username"))
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	confirmPassword := strings.TrimSpace(c.PostForm("confirm_password"))

	existEmail, err := validators.ExistEmail(email)
	if err != nil {
		log.Println("Error checking email existence:", err)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}

	if username == "" || email == "" || password == "" || confirmPassword == "" {
		resp.Error["missing"] = "Some values are missing!"
	}

	if len(username) < 4 || len(username) > 32 {
		resp.Error["username"] = "Username should be between 4 and 32"
	}

	if validators.ValidateFormatEmail(email) != nil {
		resp.Error["email"] = "Invalid email format!"
	}

	if existEmail {
		resp.Error["email"] = "Email already exists!"
	}

	if len(password) < 8 || len(password) > 16 {
		resp.Error["password"] = "Passwords should be between 8 and 16"
	}

	if password != confirmPassword {
		resp.Error["confirm_password"] = "Passwords don't match"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	user.Username = username
	user.Email = email
	user.Password = password

	db := CON.DB()

	query := "INSERT INTO user1 (username, email, password) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Username, user.Email, validators.Hash(user.Password))
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}
