package post

import (
	"log"
	"net/http"
	"os"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/models"
	"social-network-go/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

// CreateNewPost route
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

// DeletePost route
func DeletePost(c *gin.Context) {
	post := c.PostForm("post")
	db := CON.DB()

	_, dErr := db.Exec("DELETE FROM posts WHERE postID=?", post)
	utils.Err(dErr)

	utils.Json(c, map[string]interface{}{
		"mssg": "Post Deleted!!",
	})
}

// UpdatePost route
func UpdatePost(c *gin.Context) {
	postID := c.PostForm("postID")
	title := c.PostForm("title")
	content := c.PostForm("content")

	db := CON.DB()
	db.Exec("UPDATE posts SET title=?, content=? WHERE postID=?", title, content, postID)

	utils.Json(c, map[string]interface{}{
		"mssg": "Post Updated!!",
	})
}

// UpdateProfile route
func UpdateProfile(c *gin.Context) {
	resp := make(map[string]interface{})

	id, _ := utils.AllSessions(c)
	username := strings.TrimSpace(c.PostForm("username"))
	email := strings.TrimSpace(c.PostForm("email"))
	bio := strings.TrimSpace(c.PostForm("bio"))

	mailErr := checkmail.ValidateFormat(email)
	db := CON.DB()

	if username == "" || email == "" {
		resp["mssg"] = "Some values are missing!!"
	} else if mailErr != nil {
		resp["mssg"] = "Invalid email format!!"
	} else {
		_, iErr := db.Exec("UPDATE users SET username=?, email=?, bio=? WHERE id=?", username, email, bio, id)
		utils.Err(iErr)

		session := utils.GetSession(c)
		session.Values["username"] = username
		session.Save(c.Request, c.Writer)

		resp["mssg"] = "Profile updated!!"
		resp["success"] = true
	}

	utils.Json(c, resp)
}

// ChangeAvatar route
func ChangeAvatar(c *gin.Context) {
	resp := make(map[string]interface{})
	id, _ := utils.AllSessions(c)

	dir, _ := os.Getwd()
	dest := dir + "/public/users/" + id.(string) + "/avatar.png"

	dErr := os.Remove(dest)
	utils.Err(dErr)

	file, _ := c.FormFile("avatar")
	upErr := c.SaveUploadedFile(file, dest)

	if upErr != nil {
		resp["mssg"] = "An error occured!!"
	} else {
		resp["mssg"] = "Avatar changed!!"
		resp["success"] = true
	}

	utils.Json(c, resp)
}

// Follow route
func Follow(c *gin.Context) {
	id, _ := utils.AllSessions(c)
	user := c.PostForm("user")
	username := Get(user, "username")

	db := CON.DB()
	stmt, _ := db.Prepare("INSERT INTO follow(followBy, followTo, followTime) VALUES(?, ?, ?)")
	_, exErr := stmt.Exec(id, user, time.Now())
	utils.Err(exErr)

	utils.Json(c, gin.H{
		"mssg": "Followed " + username + "!!",
	})
}

// Unfollow route
func Unfollow(c *gin.Context) {
	id, _ := utils.AllSessions(c)
	user := c.PostForm("user")
	username := Get(user, "username")

	db := CON.DB()
	stmt, _ := db.Prepare("DELETE FROM follow WHERE followBy=? AND followTo=?")
	_, dErr := stmt.Exec(id, user)
	utils.Err(dErr)

	utils.Json(c, gin.H{
		"mssg": "Unfollowed " + username + "!!",
	})
}

// Like post route
func Like(c *gin.Context) {
	post := c.PostForm("post")
	db := CON.DB()
	id, _ := utils.AllSessions(c)

	stmt, _ := db.Prepare("INSERT INTO likes(postID, likeBy, likeTime) VALUES (?, ?, ?)")
	_, err := stmt.Exec(post, id, time.Now())
	utils.Err(err)

	utils.Json(c, gin.H{
		"mssg": "Post Liked!!",
	})
}

// Unlike post route
func Unlike(c *gin.Context) {
	post := c.PostForm("post")
	id, _ := utils.AllSessions(c)
	db := CON.DB()

	stmt, _ := db.Prepare("DELETE FROM likes WHERE postID=? AND likeBy=?")
	_, err := stmt.Exec(post, id)
	utils.Err(err)

	utils.Json(c, gin.H{
		"mssg": "Post Unliked!!",
	})
}
