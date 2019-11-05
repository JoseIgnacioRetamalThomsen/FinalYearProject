//go:generate protoc -I ../Login --go_out=plugins=grpc:../Login ../Login/UserLogin.proto
//go:generate protoc -I . --go_out=plugins=grpc:. UserLogin.proto
//

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserAuthenticationServer
}

// SayHello implements helloworld.GreeterServerUserResponse
func (s *server) CheckUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: %v %v", in.GetEmail(), in.GetHashPassword())

	return &pb.UserResponse{IsUser: false, Cookie: "cookie"}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	return &pb.UserResponse{IsUser: false, Cookie: "cookie"}, nil
}

func main() {
	fmt.Print("hello")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserAuthenticationServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
