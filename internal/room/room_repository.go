package room

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateRoom(ctx context.Context, room *Room) (*Room, error) {
	var lastInsertedId int
	//TODO: temp all rooms in active mode, but must be handled after creating an admin panel!
	query := "INSERT INTO rooms(name, cap, isActive) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, room.Name, room.Cap, room.IsActive).Scan(&lastInsertedId)
	if err != nil {
		return nil, err
	}
	room.ID = int64(lastInsertedId)
	return room, nil
}

func (r *repository) GetRooms(ctx context.Context) ([]*Room, error) {
	query := "SELECT * FROM rooms"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Fatal("Error fetching rooms:", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var rooms []*Room

	for rows.Next() {
		var room *Room

		if err := rows.Scan(&room.ID, &room.Name, &room.Cap, &room.IsActive); err != nil {
			log.Fatal("Error scanning row:", err)
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}
