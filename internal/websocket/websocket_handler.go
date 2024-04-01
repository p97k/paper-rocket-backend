package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server/util"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
			Message: "request failed!",
		})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, util.Response{
		Status:  http.StatusOK,
		Data:    req,
		Message: "ok",
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
			Message: "request failed!",
		})
		return
	}

	roomID := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientId,
		RoomId:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "new user has joined the room",
		RoomId:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReedMessage(h.hub)
}

func (h *Handler) LeaveRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
			Message: "request failed!",
		})
		return
	}

	roomID := c.Param("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientId,
		RoomId:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "user left the room",
		RoomId:   roomID,
		Username: username,
	}

	h.hub.Unregister <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReedMessage(h.hub)
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]CreateRoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, CreateRoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, util.Response{
		Status:  http.StatusOK,
		Data:    rooms,
		Message: "ok",
	})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, util.Response{
		Status:  http.StatusOK,
		Data:    clients,
		Message: "ok",
	})
}
