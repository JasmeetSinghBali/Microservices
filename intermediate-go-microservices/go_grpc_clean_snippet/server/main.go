package main

import (
	"log"
	"net"

	pb "github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/go_grpc_clean_snippet/proto"
	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

type helloServer struct {
	pb.RegisterGreetServiceServer
}

func main() {
	// create a listner with net package with tcp protocol on port 3000
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	// instantiate the grpc server
	grpcServer := grpc.NewServer()
	// register the greet service defined in greet.proto with grpc server instance in reff to the .pb.proto generated
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	// attach the listner to the grpcServer
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
