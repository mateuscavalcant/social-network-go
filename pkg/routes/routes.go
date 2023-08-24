package routes

import (
	"social-network-go/pkg/controller/app"
	"social-network-go/pkg/controller/auth"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", auth.Signup)
	r.POST("/validate-email", auth.ExistEmail)
	r.POST("/login", auth.UserLogin)
	r.POST("/create-post", app.CreateNewPost)
	r.POST("/follow", app.Follow)
	r.GET("/feed", app.Feed)
	r.GET("/loggout", auth.Logout)
	r.GET("/:username", app.RenderProfileTemplate)
	r.GET("/profile", app.Profile)
	r.GET("/profile/:username", app.AnotherUserProfile)
}
