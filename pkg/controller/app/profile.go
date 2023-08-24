package app

import (
	"database/sql"
	"log"
	"net/http"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/config/err"
	"social-network-go/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int    `json:"user-id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Bio           string `json:"bio"`
	Posts         int    `json:"countposts"`
	FollowBy      bool   `json:"followby"`
	FollowByCount int    `json:"followbycount"`
	FollowToCount int    `json:"followtocount"`
}

type Post struct {
	PostID     int    `json:"post-id"`
	PostUserID int    `json:"post-user-id"`
	UserID     int    `json:"user-id"`
	Content    string `json:"content"`
	CreatedBy  string `json:"createdby"`
	Name       string `json:"createdbyname"`
}

func CreateProfile(c *gin.Context) {
	var user User
	name := strings.TrimSpace(c.PostForm("name"))
	bio := strings.TrimSpace(c.PostForm("bio"))

	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	if name == "" {
		resp.Error["name"] = "Some values are missing!"
	}
	if len(name) < 1 || len(name) > 64 {
		resp.Error["name"] = "Name should be between 1 and 64"
	}
	if len(bio) > 120 {
		resp.Error["bio"] = "bio should be 120"
	}

	db := CON.DB()

	user.Name = name
	user.Bio = bio

	query := "INSERT INTO user1 (name, bio) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Name, user.Bio)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "Profile created successfully"})

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
		queryUser := `
		SELECT
			user1.username, user1.name, user1.bio,
			IFNULL(follower_counts.follower_count, 0) AS follower_count,
			IFNULL(followed_counts.following_count, 0) AS following_count
		FROM user1
		LEFT JOIN (
			SELECT followTo, COUNT(followBy) AS follower_count
			FROM user_follow
			GROUP BY followTo
		) AS follower_counts ON follower_counts.followTo = user1.id
		LEFT JOIN (
			SELECT followBy, COUNT(followTo) AS following_count
			FROM user_follow
			GROUP BY followBy
		) AS followed_counts ON followed_counts.followBy = user1.id
		WHERE user1.id = ?
	`

		errUser := db.QueryRow(queryUser, id).Scan(&user.Username, &user.Name, &user.Bio, &user.FollowByCount, &user.FollowToCount)
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

		query := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username, user1.name FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

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
			err := rows.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy, &post.Name)
			if err != nil {
				log.Println("Failed to scan statement", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to scan rows",
				})
				return
			}
			log.Println("@", post.CreatedBy)
			log.Println("Name:", post.Name)

			posts = append(posts, post)
		}

		if err := rows.Err(); err != nil {
			log.Println("Failed 3", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while iterating rows",
			})
			return
		}
		countPosts := len(posts)
		user.Posts = countPosts

		log.Println("Count Posts:", user.Posts)

		// Consulta para verificar se o usuário atual está seguindo o usuário-alvo
		queryFollow := "SELECT COUNT(*) FROM user_follow WHERE followBy = ? AND followTo = ?"
		var followCount int
		errFollow := db.QueryRow(queryFollow, id, targetUserID).Scan(&followCount)
		if errFollow != nil {
			log.Println("Failed to check follow status:", errFollow)
			// Trate o erro, se necessário
		}

		user.FollowBy = followCount > 0

		c.JSON(http.StatusOK, gin.H{
			"profile": user,
			"posts":   posts,
		})

	} else {
		var user User
		queryUser := `
		SELECT
			user1.username, user1.name, user1.bio,
			IFNULL(follower_counts.follower_count, 0) AS follower_count,
			IFNULL(followed_counts.following_count, 0) AS following_count
		FROM user1
		LEFT JOIN (
			SELECT followTo, COUNT(followBy) AS follower_count
			FROM user_follow
			GROUP BY followTo
		) AS follower_counts ON follower_counts.followTo = user1.id
		LEFT JOIN (
			SELECT followBy, COUNT(followTo) AS following_count
			FROM user_follow
			GROUP BY followBy
		) AS followed_counts ON followed_counts.followBy = user1.id
		WHERE user1.id = ?
	`
		errUser := db.QueryRow(queryUser, targetUserID).Scan(&user.Username, &user.Name, &user.Bio, &user.FollowByCount, &user.FollowToCount)
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

		targetUserPostsQuery := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username, user1.name FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

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
			err := rowsTargetUserPosts.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy, &post.Name)
			if err != nil {
				log.Println("Failed to scan target user post", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to scan rows for target user posts",
				})
				return
			}
			log.Println("@", post.CreatedBy)
			log.Println("Name:", post.Name)

			targetUserPosts = append(targetUserPosts, post)
		}

		if err := rowsTargetUserPosts.Err(); err != nil {
			log.Println("Failed while iterating target user posts", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while iterating target user posts",
			})
			return
		}
		countPosts := len(targetUserPosts)
		user.Posts = countPosts

		log.Println("Count Posts:", user.Posts)

		// Consulta para verificar se o usuário atual está seguindo o usuário-alvo
		queryFollow := "SELECT COUNT(*) FROM user_follow WHERE followBy = ? AND followTo = ?"
		var followCount int
		errFollow := db.QueryRow(queryFollow, id, targetUserID).Scan(&followCount)
		if errFollow != nil {
			log.Println("Failed to check follow status:", errFollow)
			// Trate o erro, se necessário
		}

		// Se followCount for maior que 0, o usuário atual está seguindo o usuário-alvo
		user.FollowBy = followCount > 0

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
	queryUser := `
	SELECT
		user1.username, user1.name, user1.bio,
		IFNULL(follower_counts.follower_count, 0) AS follower_count,
		IFNULL(followed_counts.following_count, 0) AS following_count
	FROM user1
	LEFT JOIN (
		SELECT followTo, COUNT(followBy) AS follower_count
		FROM user_follow
		GROUP BY followTo
	) AS follower_counts ON follower_counts.followTo = user1.id
	LEFT JOIN (
		SELECT followBy, COUNT(followTo) AS following_count
		FROM user_follow
		GROUP BY followBy
	) AS followed_counts ON followed_counts.followBy = user1.id
	WHERE user1.id = ?
`

	err := db.QueryRow(queryUser, id).Scan(&user.Username, &user.Name, &user.Bio, &user.FollowByCount, &user.FollowToCount)
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

	query := "SELECT user_post.post_id, user_post.id AS user_post_id, user_post.content, user1.id AS user1_id, user1.username, user1.name FROM user_post JOIN user1 ON user1.id = user_post.id WHERE user1.id = ?"

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

		err := rows.Scan(&post.PostID, &post.PostUserID, &post.Content, &post.UserID, &post.CreatedBy, &post.Name)
		if err != nil {
			log.Println("Failed to scan statement", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan rows",
			})
			return
		}
		log.Println("@", post.CreatedBy)
		log.Println("Name:", post.Name)

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Println("Failed 3", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error occurred while iterating rows",
		})
		return
	}
	countPosts := len(posts)
	user.Posts = countPosts

	log.Println("Count Posts:", user.Posts)

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
