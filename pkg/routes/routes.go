package routes

import (
	"social-network-go/pkg/controller"
	"social-network-go/pkg/controller/auth"
	"social-network-go/pkg/controller/post"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", auth.Signup)
	r.POST("/validate-email", auth.ExistEmail)
	r.POST("/login", auth.UserLogin)
	r.POST("/create-post", post.CreateNewPost)
	r.GET("/feed", controller.Home)
	r.GET("/loggout", auth.Logout)
}
