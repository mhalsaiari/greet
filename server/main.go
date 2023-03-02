package main

import (
	"fmt"
	g "github.com/mhalsaiari/greet/greetEng"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := g.Server{}

	grpcServer := grpc.NewServer()

	g.RegisterGreetingEngServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
