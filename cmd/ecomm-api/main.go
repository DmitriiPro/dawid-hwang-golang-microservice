package main

import (
	"davidHwang/ecomm/db"
	"davidHwang/ecomm/ecomm-api/handler"
	"davidHwang/ecomm/ecomm-api/server"
	"davidHwang/ecomm/ecomm-api/storer"
	"log"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("error opening connection to database: %v", err)
	}

	defer db.Close()

	log.Println("successfully connected to database")


	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)

	handler.RegisterRoutes(hdl)
	handler.Start(":8080")

}

//* time 13 : 35
//* Ep4  https://www.youtube.com/watch?v=v0E6JkBry7I
