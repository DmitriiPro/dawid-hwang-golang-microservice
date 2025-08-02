package main

import (
	"davidHwang/ecomm/db"
	"davidHwang/ecomm/ecomm-api/handler"
	"davidHwang/ecomm/ecomm-api/server"
	"davidHwang/ecomm/ecomm-api/storer"
	"log"

	"github.com/ianschenck/envflag"
)

const minSecretKeySize = 32

func main() {

	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")

	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters long", minSecretKeySize)
	}

	db, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("error opening connection to database: %v", err)
	}

	defer db.Close()

	log.Println("successfully connected to database")

	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)

	handler.RegisterRoutes(hdl)
	handler.Start(":8080")

}

//* time 26 : 30
//* Ep5  https://www.youtube.com/watch?v=HtsEaKuYY2o
