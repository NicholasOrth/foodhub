package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createUser(user User) {
	log.Println(user)
}

func main() {
	log.Println("Starting server...")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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

		createUser(user)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, user)
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	
	

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "test", ID: 101, Email: "test@email.com", Password: "Password"})

	// Read
	var user User
	db.First(&user, 1)                 // find product with integer primary key
	db.First(&user, "code = ?", "D42") // find product with code D42

	// Update - update product's email to x
	db.Model(&user).Update("Email", "new@email.com")
	// Update - update multiple fields

	// Delete - delete product
	db.Delete(&user, 1)
	
	
}
