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
