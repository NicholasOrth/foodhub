package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("uid")

		if uid == nil {
			log.Println("Unauthorized access")
			c.AbortWithStatus(401)
			return
		}
	}
}
