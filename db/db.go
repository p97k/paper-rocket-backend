package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	conSting := "host=server-postgres-1 port=5433 user=root password=password dbname=paper_rocket sslmode=disable"

	db, err := sql.Open("postgres", conSting)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		return
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
