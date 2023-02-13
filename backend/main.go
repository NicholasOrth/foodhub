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

func HashStr(data string) string {
	hashedData, err :=
		bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash string.")
		return ""
	}

	return string(hashedData)
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
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:id", func(c *gin.Context) {

	})

	router.GET("/auth/login", func(c *gin.Context) {
		var data struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var query User
		res := db.Where("email = ?", data.Email).First(&query)
		if res.Error != nil {
			log.Println(res.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if query.Password != data.Password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"message": "authenticated",
		})
	})

	router.POST("/auth/signup", func(c *gin.Context) {
		/*
			path the user struct name, email, and password extract
			and run thru bcrypt and store hashed info in db
		*/
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		res := db.Create(&User{
			Name:     user.Name,
			Email:    user.Email,
			Password: HashStr(user.Password),
		})

		if res.Error != nil {
			log.Println(res.Error)
		}

		log.Println("User created. Rows affected ", res.RowsAffected)
		log.Println(user)

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, nil)
	})

	err = router.Run(":7100")
	if err != nil {
		log.Fatal(err)
	}
}
