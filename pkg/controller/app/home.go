package app

import (
	"log"
	"net/http"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/models"
	"social-network-go/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateNewPost(c *gin.Context) {
	var user models.Post

	content := strings.TrimSpace(c.PostForm("content"))
	idInterface, _ := utils.AllSessions(c)
	if idInterface == nil {
		// Se o usuário não estiver logado, retornar um erro de autenticação
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id, _ := strconv.Atoi(idInterface.(string))
	user.Content = content

	db := CON.DB()

	stmt, err := db.Prepare("INSERT INTO user_post(content, id) VALUES (?, ?)")
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		// Tratar o erro, por exemplo, exibir uma mensagem de erro ou retornar um erro de servidor
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to prepare statement",
		})
		return
	}

	rs, err := stmt.Exec(user.Content, id)
	if err != nil {
		// Tratar o erro, por exemplo, exibir uma mensagem de erro ou retornar um erro de servidor
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to execute statement",
		})
		return
	}

	insertID, _ := rs.LastInsertId()

	resp := map[string]interface{}{
		"postID": insertID,
		"mssg":   "Post Created!!",
	}
	c.JSON(http.StatusOK, resp)
}

func DeletePost(c *gin.Context) {
	postID := c.PostForm("post")
	userIDInterface, _ := utils.AllSessions(c)

	if userIDInterface == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, err := strconv.Atoi(userIDInterface.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	db := CON.DB()

	var postAuthorID int
	err = db.QueryRow("SELECT id FROM user_post WHERE postID=?", postID).Scan(&postAuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch post details",
		})
		return
	}

	if postAuthorID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You don't have permission to delete this post",
		})
		return
	}

	_, err = db.Exec("DELETE FROM user_post WHERE postID=?", postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete post",
		})
		return
	}

	resp := map[string]interface{}{
		"mssg": "Post Deleted!",
	}
	c.JSON(http.StatusOK, resp)
}
