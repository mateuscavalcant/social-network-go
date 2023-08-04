package routes

import (
	"social-network-go/pkg/controller/auth"
	"social-network-go/pkg/controller/home"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", auth.Signup)
	r.POST("/validate-email", auth.ExistEmail)
	r.POST("/login", auth.UserLogin)
	r.POST("/create-post", home.CreateNewPost)
	r.GET("/feed", home.Feed)
	r.GET("/loggout", auth.Logout)
}
