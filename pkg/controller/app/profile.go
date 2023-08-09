package app

import (
	"log"
	"net/http"
	CON "social-network-go/pkg/config"

	"github.com/gin-gonic/gin"
)

type Post struct {
	PostID     int    `json:"post-id"`
	PostUserID int    `json:"post-user-id"`
	UserID     int    `json:"user-id"`
	Content    string `json:"content"`
	CreatedBy  string `json:"createdby"`
}

func ProfilePost(c *gin.Context) {
	username := c.Param("username")
	db := CON.DB()

	// Fetch user ID based on username
	var userID int
	userQuery := "SELECT id FROM user1 WHERE username = ?"
	err := db.QueryRow(userQuery, username).Scan(&userID)
	if err != nil {
		log.Println("Failed to fetch user ID", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user ID",
		})
		return
	}

	posts := []Post{}

	query := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

	rows, err := db.Query(query, userID)
	if err != nil {
		log.Println("Failed to query statement", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to execute query",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
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

	// Render the HTML template
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"username": username,
		"posts":    posts,
	})
}
