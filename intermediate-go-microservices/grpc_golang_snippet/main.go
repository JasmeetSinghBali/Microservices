package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Entry point of grpc server")

	// grpc server instance as gs
	gs := grpc.NewServer()

	protos.RegisterFiatServer(gs)
}
