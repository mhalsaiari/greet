package main

import (
	g "github.com/mhalsaiari/greet/greetEng"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	//c := g.NewChatServiceClient(conn)
	c := g.NewGreetingEngClient(conn)

	response, err := c.SayHello(context.Background(), &g.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
