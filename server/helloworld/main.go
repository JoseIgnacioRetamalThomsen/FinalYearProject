//basic server for test clint-server comunication
// base in go grpc oficial examples

// user helloword.proto which must be compiled using:
//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto
// then the generated file must be places in $GO-PATH\src\github.com\JoseIgnacioRetamalThomsen\wcity



package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/JoseIgnacioRetamalThomsen/wcity"
)

const (
	port = ":7776"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}