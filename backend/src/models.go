package main

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
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
	gorm.Model

	UserID uint
	PostID uint
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
