package main

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
)

type JsonModel struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `gorm:"index" json:"deletedAt"`
}

type User struct {
	JsonModel

	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	Posts []Post `json:"posts"`
}

type Post struct {
	JsonModel

	UserID  uint   `json:"userId"`
	Caption string `json:"caption"`
	ImgPath string `json:"imgPath"`

	Likes []Like `json:"likes"`
}

type Like struct {
	JsonModel

	UserID uint
	PostID uint
}

type Follow struct {
	JsonModel

	UserID     uint // user being followed
	FollowerID uint // user following
}

type Block struct {
	JsonModel

	UserID    uint // user being blocked
	BlockerID uint // user blocking
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
