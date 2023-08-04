package auth

import (
	"social-network-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	utils.LoggedIn(c, "")
	session := utils.GetSession(c)
	delete(session.Values, "id")
	delete(session.Values, "email")
	session.Save(c.Request, c.Writer)

}
