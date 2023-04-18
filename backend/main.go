package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbInit(true)

	// gin web server
	router := gin.Default()

	store, err :=
		redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	if err != nil {
		log.Fatal(err)
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	})

	router.Use(sessions.Sessions("session", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	router.Static("/images", "./images")

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:id/posts", userPosts)
	router.GET("/user/:id", userInfo)
	router.POST("/auth/login", login)
	router.POST("/auth/signup", signup)
	router.POST("/auth/logout", logout)
	router.GET("/post/info/:id", postInfo)

	protected := router.Group("/")
	protected.Use(Authentication())
	protected.POST("/user/:id/follow", followUser)
	protected.POST("/user/:id/following", userFollowing)
	protected.GET("/feed", feed)
	protected.POST("/post/create", createPost)
	protected.POST("/post/delete/:id", deletePost)
	protected.POST("/post/like/:id", likePost)
	protected.GET("/user/me", userMe)

	protected.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "protected",
		})
	})

	err = router.Run(":7100")
	if err != nil {
		log.Fatal(err)
	}
}
