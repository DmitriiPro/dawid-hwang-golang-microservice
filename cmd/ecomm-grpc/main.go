package main

import (
	"davidHwang/ecomm/db"
	"davidHwang/ecomm/ecomm-grpc/pb"
	"davidHwang/ecomm/ecomm-grpc/server"
	"davidHwang/ecomm/ecomm-grpc/storer"
	"log"
	"net"

	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
)

func main() {
	var (
		svcAddr = envflag.String("SVC_ADDR", "0.0.0.0:9091", "address where the ecomm-grpc service is listening on")
	)

	//*создадим

	//* 1 экземпляр базы данных
	db, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("error opening connection to database: %v", err)
	}

	defer db.Close()

	log.Println("successfully connected to database")

	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())

	//* 2 экземпляр сервера
	srv := server.NewServer(st)

	//* 3 зарегистрируем сервер в GRPC сервере
	grpcServer := grpc.NewServer()
	pb.RegisterEcommServer(grpcServer, srv)

	listener, err := net.Listen("tcp", *svcAddr)
	if err != nil {
		log.Fatalf("failed to listener: %v", err)
	}

	log.Printf("server listening on %s", *svcAddr)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("<< failed to serve >>: %v", err)
	}

}
