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
func hashUser(user *User) {
	// hash name
	hashedName, err :=
		bcrypt.GenerateFromPassword([]byte(user.Name), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash user struct.")
		return
	}

	//hash email
	hashedEmail, err :=
		bcrypt.GenerateFromPassword([]byte(user.Email), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash user struct.")
		return
	}

	hashedPass, err :=
		bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash user struct.")
		return
	}

	user.Name = string(hashedName)
	user.Email = string(hashedEmail)
	user.Password = string(hashedPass)
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
		/*
			path the user struct name, email, and password extract
			and run thru bcrypt and store hashed info in db
		*/
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		hashUser(&user)

		res := db.Create(&user)
		if res.Error != nil {
			log.Println(res.Error)
		}
		log.Println("User created. Rows affected ", res.RowsAffected)
		log.Println(user)

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, nil)
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
