package app

import (
	"log"
	"net/http"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	utils.LoggedIn(c, "/welcome")

	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))
	db := CON.DB()

	type Post struct {
		PostID            int    `json:"post-id"`
		PostUserID        int    `json:"post-user-id"`
		UserID            int    `json:"user-id"`
		Content           string `json:"content"`
		CreatedByUsername string `json:"createdby"`
		CreatedByName     string `json:"createdbyname"`
	}
	var post Post
	post.UserID = id

	posts := []Post{}

	query := `
		SELECT user_post.post_id, user_post.id AS post_user_id, user_post.content,
		       user1.id AS user_id, user1.username, user1.name
		FROM user_post
		JOIN user1 ON user1.id = user_post.id
		WHERE user1.id = ? OR user1.id IN (
		    SELECT user_follow.followTo
		    FROM user_follow
		    WHERE user_follow.followBy = ?
		)
	`

	rows, err := db.Query(query, id, id)
	if err != nil {
		log.Println("Failed to query statement", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to execute query",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedByUsername, &post.CreatedByName)
		if err != nil {
			log.Println("Failed to scan statement", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan rows",
			})
			return
		}
		log.Println("CreatedBy:", post.CreatedByUsername)

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
		"posts": posts,
	})
}
