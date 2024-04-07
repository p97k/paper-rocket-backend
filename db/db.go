package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open(
		"postgres",
		"postgresql://ubuntu-paper-rocket-cloud:uBenZsyrsw@188.121.97.228:5432/paper-rocket?sslmode=disable")
	if err != nil {
		return nil, err
	}

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
