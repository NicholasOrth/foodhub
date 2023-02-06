package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	log.Println("Starting server...")

	// db connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}

	// gin web server
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/user/:id", func(c *gin.Context) {

	})

	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		res := db.Create(&user)
		if res.Error != nil {
			log.Println(res.Error)
		}
		log.Println("User created. Rows affected ", res.RowsAffected)

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, user)
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
