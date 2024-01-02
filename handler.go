package handler

import (
	"log"
	"net/http"
	"social-network-go/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.LoadHTMLGlob("C:/social-network-go/templates/*")
	router.Static("/public", "./public")

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", gin.H{})
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	routes.InitRoutes(router.Group("/"))

	// Agora, você pode usar o método ServeHTTP para a execução no Vercel
	router.ServeHTTP(w, r)
}
