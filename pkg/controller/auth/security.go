package auth

import (
	"log"
	"social-network-go/pkg/config/err"
	"social-network-go/pkg/config/validators"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExistEmail(c *gin.Context) {
	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	email := strings.TrimSpace(c.PostForm("email"))
	existEmail, err := validators.ExistEmail(email)
	if err != nil {
		log.Println("Error checking email existence:", err)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}

	if email == "" {
		resp.Error["missing"] = "Some values are missing!"
	}

	if validators.ValidateFormatEmail(email) != nil {
		resp.Error["email"] = "Invalid email format!"
	}

	if existEmail {
		resp.Error["email"] = "Email already exists!"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	c.JSON(200, gin.H{"message": "Email Validad"})

}
