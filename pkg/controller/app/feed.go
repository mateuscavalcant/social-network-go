package app

import (
	"encoding/base64"
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
		PostID       int    `json:"post-id"`
		PostUserID   int    `json:"post-user-id"`
		UserID       int    `json:"user-id"`
		Content      string `json:"content"`
		CreatedBy    string `json:"createdby"`
		ProfileImage string `json:"profile-image"`
	}
	var post Post
	post.UserID = id

	posts := []Post{}

	query := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username, user1.profile_image FROM user_post JOIN user1 ON user1.id = user_post.id"

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Failed to query statement", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to execute query",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var profileImageBlob []byte
		err := rows.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy, &profileImageBlob)
		if err != nil {
			log.Println("Failed to scan statement", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan rows",
			})
			return
		}
		log.Println("CreatedBy:", post.CreatedBy)

		// Codificar a imagem em base64 e incluir no campo ProfileImage
		post.ProfileImage = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(profileImageBlob)

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
