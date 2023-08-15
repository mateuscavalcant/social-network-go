package app

import (
	"database/sql"
	"log"
	"net/http"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"user-id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	PostID     int    `json:"post-id"`
	PostUserID int    `json:"post-user-id"`
	UserID     int    `json:"user-id"`
	Content    string `json:"content"`
	CreatedBy  string `json:"createdby"`
}

var user User

func YourProfile(c *gin.Context) {
	utils.LoggedIn(c, "/welcome")

	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))
	db := CON.DB()
	var post Post
	post.UserID = id

	// Fetch user information
	var user User
	queryUser := "SELECT username, email FROM user1 WHERE id = ?"
	err := db.QueryRow(queryUser, id).Scan(&user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println("Failed to query user information:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user information",
		})
		return
	}

	posts := []Post{}

	query := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

	rows, err := db.Query(query, id)
	if err != nil {
		log.Println("Failed to query statement", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to execute query",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy)
		if err != nil {
			log.Println("Failed to scan statement", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan rows",
			})
			return
		}
		log.Println("CreatedBy:", post.CreatedBy)

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Println("Failed 3", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error occurred while iterating rows",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": user,
		"posts":   posts,
	})
}

func YourProfileTemplate(c *gin.Context) {
	// Recupere o nome de usuário da URL
	username := c.Param("username")

	// Recupere o ID da sessão (convertido para int)
	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))

	user.Username = username

	db := CON.DB()

	// Verifique se o usuário existe e está relacionado ao ID da sessão
	query := "SELECT username FROM user1 WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println("Failed to query user information:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user information",
		})
		return
	}

	// Renderize o template HTML com os dados
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"username": username,
	})
}
