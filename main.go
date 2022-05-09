package main

import (
	"server-service/routes"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// DONE:
// user struct done
// all struct done

// TODO:
// update user service
// update membership service
// update membership payment service
// update event service
// update contact service
// update category service

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	routes.UserRoute(r)
	routes.MembershipRoute(r)
	routes.MembershipPaymentRoute(r)
	routes.EventRoute(r)
	routes.ContactRoute(r)
	routes.CategoryRoute(r)
	r.Run()
}
