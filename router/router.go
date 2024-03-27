package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(db *sql.DB) {
	router = gin.Default()
	router.Use(CORSMiddleware())

	InitUserRoutes(db, router)
	InitWebsocketRoutes(router)
}

func Start(address string) error {
	return router.Run(address)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
