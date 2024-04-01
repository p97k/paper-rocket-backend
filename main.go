package main

import (
	"log"
	"server/db"
	"server/router"
)

func main() {
	dbConnection, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	} else {
		log.Printf("connected to database!")
	}

	router.InitRouter(dbConnection.GetDB())
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("router couldn't start: %s", err)
	}
}
