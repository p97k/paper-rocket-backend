package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"server/internal/user"
)

func InitUserRoutes(db *sql.DB, router *gin.Engine) {
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	router.POST("/signup/", userHandler.CreateUser)
	router.POST("/login/", userHandler.Login)
	router.GET("/logout/", userHandler.Logout)
}
