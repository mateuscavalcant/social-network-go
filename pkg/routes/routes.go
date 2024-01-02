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
	r.POST("/unfollow", app.Unfollow)
	r.POST("/feed", app.Feed)
	r.POST("/loggout", auth.Logout)
	r.GET("/:username", app.RenderProfileTemplate)
	r.POST("/profile", app.Profile)
	r.POST("/profile/:username", app.AnotherUserProfile)

}
