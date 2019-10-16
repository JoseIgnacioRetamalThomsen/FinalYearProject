//go:generate protoc -I ../Login --go_out=plugins=grpc:../Login ../Login/UserLogin.proto
//go:generate protoc -I . --go_out=plugins=grpc:. UserLogin.proto
// 



package main

import (
	"log"
	"fmt"
	"context"
	"net"
	"google.golang.org/grpc"
	pb "github.com/JoseIgnacioRetamalThomsen/wcity"
)


const (
	port = ":50051"
)


type server struct {
	pb.UnimplementedUserAuthenticationServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Check(ctx context.Context, in *pb.UserData) (*pb.LoginResponse, error) {
	log.Printf("Received: %v %v", in.GetEmail(), in.GetHashPassword())
	
	return &pb.LoginResponse{IsUser: false,Cookie: "cookie"}, nil
}

func (s *server) Create(ctx context.Context, in *pb.UserData)(*pb.LoginResponse, error){

	return &pb.LoginResponse{IsUser: false,Cookie: "cookie"}, nil
}

func main(){
	fmt.Print("hello");
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
