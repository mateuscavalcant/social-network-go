package app

import (
	"net/http"
	CON "social-network-go/pkg/config"
	"social-network-go/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	id, _ := utils.AllSessions(c)
	user := c.PostForm("user")
	username := c.PostForm("username")

	db := CON.DB()
	stmt, _ := db.Prepare("INSERT INTO follow(followBy, followTo, followTime) VALUES(?, ?, ?)")
	_, err := stmt.Exec(id, user, time.Now())
	if err != nil {

	}

	resp := map[string]interface{}{
		"mssg": "Followed " + username,
	}
	c.JSON(http.StatusOK, resp)
}

func Unfollow(c *gin.Context) {
	id, _ := utils.AllSessions(c)
	user := c.PostForm("user")
	username := c.PostForm("username")

	db := CON.DB()
	stmt, _ := db.Prepare("DELETE FROM follow WHERE followBy=? AND followTo=?")
	_, err := stmt.Exec(id, user)
	if err != nil {

	}

	resp := map[string]interface{}{
		"mssg": "Unfollowed " + username + "!!",
	}
	c.JSON(http.StatusOK, resp)
}
