package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MakeTimestamp function
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// Err Log
func Err(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

// MeOrNot function to checked whether it's me or not
func MeOrNot(c *gin.Context, user string) bool {
	var id interface{}
	id, _ = AllSessions(c)
	if id != user {
		return false
	}
	return true
}

func RenderTemplate(c *gin.Context, tmpl string, p interface{}) {
	c.HTML(http.StatusOK, tmpl+".html", p)
}

func Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Ses(c *gin.Context) interface{} {
	id, username := AllSessions(c)
	return map[string]interface{}{
		"id":       id,
		"username": username,
	}
}

func LoggedIn(c *gin.Context, urlRedirect string) {
	var URL string
	if urlRedirect == "" {
		URL = "/login"
	} else {
		URL = urlRedirect
	}
	id, _ := AllSessions(c)
	if id == nil {
		c.Redirect(http.StatusFound, URL)
	}
}

func NotLoggedIn(c *gin.Context) {
	id, _ := AllSessions(c)
	if id != nil {
		c.Redirect(http.StatusFound, "/")
	}
}

func Invalid(c *gin.Context, what int) {
	if what == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
}
