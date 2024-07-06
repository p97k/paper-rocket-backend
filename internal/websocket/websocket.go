package websocket

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateRoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
