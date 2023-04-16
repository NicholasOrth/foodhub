package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbInit(true)

	// gin web server
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	router.Static("/images", "./images")

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/me", userMe)
	router.GET("/user/:id/posts", userPosts)
	router.GET("/user/:id", userInfo)
	router.POST("/user/:id/follow", followUser)

	router.POST("/auth/login", login)
	router.POST("/auth/signup", signup)

	router.POST("/auth/logout", logout)

	router.POST("/post/create", createPost)

	router.POST("/post/like/:id", likePost)
	router.GET("/post/info/:id", postInfo)

	router.GET("/feed", feed)

	err = router.Run(":7100")
	if err != nil {
		log.Fatal(err)
	}
}
