package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/websocket"
)

func InitWebsocketRoutes(router *gin.Engine) {
	hub := websocket.NewHub()
	websocketHandler := websocket.NewHandler(hub)

	go hub.Run()

	router.POST("/ws/create-room/", websocketHandler.CreateRoom)
	router.GET("/ws/join-room/:roomId/", websocketHandler.JoinRoom)
	router.GET("/ws/leave-room/:roomId/", websocketHandler.LeaveRoom)
	router.GET("/ws/get-room/", websocketHandler.GetRooms)
	router.GET("/ws/get-clients/:roomId/", websocketHandler.GetClients)
}
