package auth

import (
	"log"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/config/err"
	"social-network-go/pkg/models"
	"social-network-go/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User

	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))

	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	user.Email = email
	user.Password = password

	db := CON.DB()

	err := db.QueryRow("SELECT id, email, password FROM user1 WHERE email=?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		resp.Error["email"] = "Invalid credentials"
	}

	encErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if encErr != nil {
		resp.Error["password"] = "Invalid password"

	}
	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return

	}
	session := utils.GetSession(c)
	session.Values["id"] = strconv.Itoa(user.ID)
	session.Values["email"] = user.Email
	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{"message": "User created successfully"})

}
