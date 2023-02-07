package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// function for hashing user info
func hasher(user User) User {
	//hashing user info
	hashedUser := User{}
	name := []byte(user.Name)
	hashedName, err := bcrypt.GenerateFromPassword(name, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	//hash password
	pass := []byte(user.Password)

	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	//hash email
	email := []byte(user.Email)
	hashedEmail, err := bcrypt.GenerateFromPassword(email, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedUser.Name = string(hashedName)
	hashedUser.Email = string(hashedEmail)
	hashedUser.Password = string(hashedPass)
	return hashedUser
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
		//path the user struct name, email, and password extract and run thru bcrypt and store hashed info in db
		//
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//hashing user info
		hashedUser := hasher(user)

		res := db.Create(&hashedUser)
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
