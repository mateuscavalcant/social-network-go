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

func AnotherUserProfile(c *gin.Context) {
	utils.LoggedIn(c, "/welcome")

	username := c.Param("username")
	var post Post

	// Obtenha o ID do usuário alvo usando o nome de usuário
	db := CON.DB()
	var targetUserID int
	queryUserID := "SELECT id FROM user1 WHERE username = ?"
	err := db.QueryRow(queryUserID, username).Scan(&targetUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println("Failed to query target user information:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch target user information",
		})
		return
	}

	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))

	if id == targetUserID {
		var user User
		queryUser := "SELECT username, email FROM user1 WHERE id = ?"
		errUser := db.QueryRow(queryUser, id).Scan(&user.Username, &user.Email)
		if errUser != nil {
			if errUser == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}
			log.Println("Failed to query user information:", errUser)
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

	} else {
		var user User
		queryUser := "SELECT username, email FROM user1 WHERE id = ?"
		errUser := db.QueryRow(queryUser, targetUserID).Scan(&user.Username, &user.Email)
		if errUser != nil {
			if errUser == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}
			log.Println("Failed to query target user information:", errUser)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch target user information",
			})
			return
		}
		targetUserPosts := []Post{}

		targetUserPostsQuery := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

		rowsTargetUserPosts, err := db.Query(targetUserPostsQuery, targetUserID)
		if err != nil {
			log.Println("Failed to query target user posts", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to execute query for target user posts",
			})
			return
		}
		defer rowsTargetUserPosts.Close()

		for rowsTargetUserPosts.Next() {
			err := rowsTargetUserPosts.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy)
			if err != nil {
				log.Println("Failed to scan target user post", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to scan rows for target user posts",
				})
				return
			}
			log.Println("CreatedBy (Target User Post):", post.CreatedBy)

			targetUserPosts = append(targetUserPosts, post)
		}

		if err := rowsTargetUserPosts.Err(); err != nil {
			log.Println("Failed while iterating target user posts", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while iterating target user posts",
			})
			return
		}

		// Retorne o perfil público do usuário alvo com seus posts públicos
		c.JSON(http.StatusOK, gin.H{
			"profile": user,
			"posts":   targetUserPosts,
		})
	}
}

func Profile(c *gin.Context) {
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

func RenderProfileTemplate(c *gin.Context) {
	// Recupere o nome de usuário da URL
	username := c.Param("username")

	// Recupere o ID da sessão (convertido para int)
	idInterface, _ := utils.AllSessions(c)
	id, _ := strconv.Atoi(idInterface.(string))

	db := CON.DB()

	// Verifique se o usuário existe
	queryExist := "SELECT COUNT(*) FROM user1 WHERE username = ?"
	var count int
	err := db.QueryRow(queryExist, username).Scan(&count)
	if err != nil {
		log.Println("Failed to query user existence:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check user existence",
		})
		return
	}

	if count == 0 {
		// Renderize o template notexistuser.html se o usuário não existir
		c.HTML(http.StatusOK, "notfounduser.html", gin.H{})
		return
	}

	// Verifique se o usuário pertence à sessão
	var userSession User
	queryUserSession := "SELECT id, username FROM user1 WHERE id = ?"
	err = db.QueryRow(queryUserSession, id).Scan(&userSession.ID, &userSession.Username)
	if err != nil {
		log.Println("Failed to query user session information:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user session information",
		})
		return
	}

	if userSession.Username != username {
		// Renderize o template anotheruserprofile.html se o usuário não pertencer à sessão
		c.HTML(http.StatusOK, "another_profile.html", gin.H{
			"username": username,
		})
		return
	}

	// Renderize o template profile.html com os dados
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"username": username,
	})
}
