package routes

import (
	"social-network-go/pkg/controller/auth"
	"social-network-go/pkg/controller/app"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", auth.Signup)
	r.POST("/validate-email", auth.ExistEmail)
	r.POST("/login", auth.UserLogin)
	r.POST("/create-post", app.CreateNewPost)
	r.GET("/feed", app.Feed)
	r.GET("/loggout", auth.Logout)
	r.GET("/profile", app.ProfilePost)
}
