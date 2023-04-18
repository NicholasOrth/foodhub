package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

/*func AddFollower(user User, target User) {
	user.Following = append(user.Following, target.ID)
}


func BlockUser(user User, targetID uint) []uint {
	newList := user.Blocked
	if Contains(user.Blocked, targetID) {
		newList = RemoveFromSlice(user.Blocked, targetID)
	} else {
		newList = append(user.Blocked, targetID)
	}
	return newList
}*/

func HashStr(data string) string {
	hashedData, err :=
		bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash string.")
		return ""
	}

	return string(hashedData)
}

//
//func AuthUser(c *gin.Context, db *gorm.DB) (User, Claims, error) {
//	cookie, err := c.Cookie("jwt")
//	if err != nil {
//		log.Println("No cookie found.")
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return User{}, Claims{}, err
//	}
//
//	claims := &Claims{}
//
//	token, err := jwt.ParseWithClaims(cookie, claims,
//		func(token *jwt.Token) (interface{}, error) {
//			return backend.JwtKey, nil
//		})
//
//	if err != nil || !token.Valid {
//		log.Println("Invalid token.")
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return User{}, Claims{}, err
//	}
//
//	var user User
//	res := db.First(&user, claims.ID)
//
//	if res.Error != nil {
//		log.Println(res.Error)
//		c.AbortWithStatus(http.StatusInternalServerError)
//		return User{}, Claims{}, res.Error
//	}
//
//	return user, *claims, nil
//}

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
