package main

import (
	"log"
	"net/http"
	"social-network-go/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.LoadHTMLGlob("C:/social-network-go/templates/*")
	r.Static("/public", "./public")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/create-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", gin.H{})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	routes.InitRoutes(r.Group("/"))

	errServer := r.Run(":8080")
	if errServer != nil {
		log.Fatal(errServer)
	}
}
