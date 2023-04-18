package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* User Routes */
func userMe(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid == nil {
		log.Println("Unauthorized access")
		c.AbortWithStatus(401)
		return
	}

	var user User
	res := db.First(&user, uid.(uint))
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var posts []Post

	err := db.Model(&user).Association("Posts").Find(&posts)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var followers []Follow
	err = db.Where("user_id = ?", user.ID).Find(&followers).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var following []Follow
	err = db.Where("follower_id = ?", user.ID).Find(&following).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"name":      user.Name,
		"followers": len(followers),
		"following": len(following),
		"posts":     posts,
	})
}
func userPosts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user User
	res := db.First(&user, id)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
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
		"posts": posts,
	})
}
func userInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user User
	res := db.First(&user, id)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var followers []Follow
	err = db.Where("user_id = ?", user.ID).Find(&followers).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var following []Follow
	err = db.Where("follower_id = ?", user.ID).Find(&following).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"followers": len(followers),
		"following": len(following),
	})
}

func followUser(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid == nil {
		log.Println("Unauthorized access")
		c.AbortWithStatus(401)
		return
	}

	var user User
	res := db.First(&user, uid.(uint))
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var userToFollow User
	res = db.First(&userToFollow, id)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	follow := Follow{
		UserID:     userToFollow.ID,
		FollowerID: user.ID,
	}

	res = db.Create(&follow)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "followed user",
	})
}

/* Auth Routes */
func login(c *gin.Context) {
	session := sessions.Default(c)

	uid := session.Get("uid")

	if uid != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "already logged in",
		})
	}

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

	session.Set("uid", query.ID)
	err = session.Save()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func signup(c *gin.Context) {
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
}
func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

/* Post Routes */
func createPost(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid == nil {
		log.Println("Unauthorized access")
		c.AbortWithStatus(401)
		return
	}

	var user User
	res := db.First(&user, uid.(uint))
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

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
}
func deletePost(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid == nil {
		log.Println("Unauthorized access")
		c.AbortWithStatus(401)
		return
	}

	var user User
	res := db.First(&user, uid.(uint))
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	postID := c.Param("id")

	var post Post
	err := db.Where("id = ? AND user_id = ?", postID, user.ID).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}

	err = os.Remove(post.ImgPath)
	if err != nil {
		log.Println(err)
	}

	err = db.Delete(&post).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func likePost(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")

	if uid == nil {
		log.Println("Unauthorized access")
		c.AbortWithStatus(401)
		return
	}

	var user User
	res := db.First(&user, uid.(uint))
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var post Post
	res = db.First(&post, postID)

	if res.Error != nil {
		log.Println(res.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var like Like
	tx := db.Where("user_id = ? AND post_id = ?", user.ID, post.ID).First(&like)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	if like.ID != 0 { // If the like exists, delete it
		err := db.Model(&post).Association("Likes").Delete(&like)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		like := Like{UserID: user.ID, PostID: post.ID}
		err := db.Model(&post).Association("Likes").Append(&like)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	likes := db.Model(&post).Association("Likes").Count()

	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}
func postInfo(c *gin.Context) {
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

	likes := db.Model(&post).Association("Likes").Count()

	c.JSON(http.StatusOK, gin.H{
		"post":  post,
		"likes": likes,
	})
}

/* Feed Routes */
func feed(c *gin.Context) {
	var posts []Post

	err := db.Order("created_at").Find(&posts).Limit(10).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
