package main

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Following []uint `json:"following"`
	//Followers []uint `json:"followers"`

	Posts []Post `json:"posts" gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Caption string `json:"caption"`
	ImgPath string `json:"imgPath"`

	UserID uint `json:"userId"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

//func AddFollower(user User, target User) {
//	// TODO add friend to user, and add follower to friend
//	user.Following = append(user.Following, target.ID)
//	target.Followers = append(target.Followers, user.ID)
//}

// TODO: add remove follower function
func Contains(slice []uint, val uint) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []uint, val uint) []uint {
	// iterate over the slice and copy all elements except u to a new slice
	var result []uint
	for _, i := range slice {
		if i != val {
			result = append(result, i)
		}
	}
	return result
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

func AuthUser(c *gin.Context, db *gorm.DB) (User, Claims, error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		log.Println("No cookie found.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return User{}, Claims{}, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(cookie, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		log.Println("Invalid token.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return User{}, Claims{}, err
	}

	var user User
	res := db.First(&user, claims.ID)

	if res.Error != nil {
		log.Println(res.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return User{}, Claims{}, res.Error
	}

	return user, *claims, nil
}

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// db connection
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		log.Println(err)
	}

	// gin web server
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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

	router.GET("/user", func(c *gin.Context) {
		user, _, err := AuthUser(c, db)
		if err != nil {
			return
		}

		var posts []Post

		err = db.Model(&user).Association("Posts").Find(&posts)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"email": user.Email,
			"name":  user.Name,
			"posts": posts,
		})
	})

	router.POST("/auth/login", func(c *gin.Context) {
		var creds Credentials

		// Get the JSON body and decode into credentials
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Get the user details from the database
		var query User
		res := db.Where("email = ?", creds.Email).First(&query)

		if res.Error != nil {
			log.Println(res.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Compare the stored hashed password, with the hashed version of the password that was received
		err := bcrypt.CompareHashAndPassword([]byte(query.Password), []byte(creds.Password))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If authentication is successful, generate a token
		expiration := time.Now().Add(time.Hour).Unix()

		claims := &Claims{
			ID:    query.ID,
			Name:  query.Name,
			Email: query.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiration,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("Content-Type", "application/json")
		c.SetCookie(
			"jwt",
			tokenString,
			3600,
			"/",
			"localhost",
			false,
			true)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.POST("/auth/signup", func(c *gin.Context) {
		/*
			path the user struct name, email, and password extract
			and run through bcrypt and store hashed info in db
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

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, nil)
	})

	router.POST("/image/upload", func(c *gin.Context) {
		user, _, err := AuthUser(c, db)

		caption := c.PostForm("caption")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		filename := filepath.Base(file.Filename)
		path := "images/user/" + strconv.Itoa(int(user.ID)) + "/"

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = c.SaveUploadedFile(file, path+filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = db.Model(&user).Association("Posts").Append(&Post{
			Caption: caption,
			ImgPath: path + filename,
		})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	err = router.Run(":7100")
	if err != nil {
		log.Fatal(err)
	}
}
