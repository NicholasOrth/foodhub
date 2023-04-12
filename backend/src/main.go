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
	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	Posts []Post `json:"posts"`
}

type Post struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`

	UserID  uint   `json:"userId"`
	Caption string `json:"caption"`
	ImgPath string `json:"imgPath"`

	Likes []Like `json:"likes"`
}

type Like struct {
	UserID    uint `gorm:"primaryKey"`
	PostID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
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

func (p *Post) Like(userID uint) {
	u := false

	for i, like := range p.Likes {
		if like.UserID == userID {
			p.Likes = append(p.Likes[:i], p.Likes[i+1:]...)
			u = true
			break
		}
	}

	if !u {
		p.Likes = append(p.Likes, Like{UserID: userID, PostID: p.ID})
	}
}

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

	err = db.AutoMigrate(&User{}, &Post{}, &Like{})
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

	router.GET("/user/me", func(c *gin.Context) {
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
	router.GET("/user/posts", func(c *gin.Context) {
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

		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	})
	router.GET("/user/:id", func(c *gin.Context) {})

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

		// Compare the stored hashed password, with the hashed version of the password that was received\
		if query.Email != creds.Email {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
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

	router.POST("/auth/logout", func(c *gin.Context) {
		c.SetCookie(
			"jwt",
			"",
			-1,
			"/",
			"localhost",
			false,
			true)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.POST("/post/create", func(c *gin.Context) {
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

	router.POST("/post/like/:id", func(c *gin.Context) {
		user, _, err := AuthUser(c, db)
		if err != nil {
			return
		}

		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var post Post
		res := db.First(&post, postID)

		if res.Error != nil {
			log.Println(res.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		post.Like(user.ID)

		err = db.Save(&post).Error
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"likes": len(post.Likes),
		})
	})
	router.GET("/post/info/:id", func(c *gin.Context) {
		_, _, err := AuthUser(c, db)
		if err != nil {
			return
		}

		postID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var post Post
		res := db.First(&post, postID)

		if res.Error != nil {
			log.Println(res.Error)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"post":  post,
			"likes": len(post.Likes),
		})
	})

	router.GET("/feed", func(c *gin.Context) {
		_, _, err := AuthUser(c, db)
		if err != nil {
			return
		}

		var posts []Post

		err = db.Order("created_at").Find(&posts).Limit(10).Error
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	})
	// follow a user
	/*router.POST("/user/follow/:id", func(c *gin.Context) {
		// get the authenticated user
		authUser, _, err := AuthUser(c, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// get the user to follow
		targetID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid user ID",
			})
			return
		}

		// check if the user to follow exists
		var target User
		res := db.First(&target, targetID)
		if res.Error != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		// update the following list of the authenticated user
		authUser.Following = append(authUser.Following, target.ID)
		err = db.Save(&authUser).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update following list",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully followed",
		})
	})*/

	err = router.Run(":7100")
	if err != nil {
		log.Fatal(err)
	}
}
