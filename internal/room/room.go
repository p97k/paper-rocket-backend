package room

import "context"

type Room struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Cap      int    `json:"cap" db:"cap"`
	IsActive bool   `json:"isActive" db:"isActive"`
}

type Repository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
	GetRooms(ctx context.Context) ([]*Room, error)
}
