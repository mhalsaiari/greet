package greetEng

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedGreetingEngServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello I am the wonderful engineer who joined Tweeq"}, nil
}
