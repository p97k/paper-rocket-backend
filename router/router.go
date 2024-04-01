package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/util"
)

var router *gin.Engine

func InitRouter(db *sql.DB) {
	router = gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", initBaseRouter)

	InitUserRoutes(db, router)
	InitWebsocketRoutes(router)
}

func Start(address string) error {
	return router.Run(address)
}

func initBaseRouter(c *gin.Context) {
	c.JSON(http.StatusOK, util.Response{
		Data:    nil,
		Status:  http.StatusOK,
		Message: "welcome to paper rocket app! :))",
	})
}
