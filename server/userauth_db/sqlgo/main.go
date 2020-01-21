package main

import (
	//"os"
	//	"bytes"
	"fmt"

	"github.com/joseignacioretamalthomsen/sqlgo/db" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine

	"context"
	"log"
	"net"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const (
	port = ":7777"
)

type server struct {
	pb.UnimplementedUserLogDBServer
}

func (s *server) AddUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Create new user")
	u := db.NewUser(in.Email, "", in.HashedPassword, in.Salt, false)
	fmt.Println(db.AddUser(*u))
	return &pb.UserDBResponse{Id: 10, Email: in.Email, HashedPassword: in.HashedPassword, Salt: in.Salt}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Get user")
	u := db.GetUser(in.Email)

	return &pb.UserDBResponse{Email: u.GetEmail(), HashedPassword: u.GetHashedPassword(), Salt: u.GetSalt()}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Update user")
	u := db.NewUser(in.Email, "", in.HashedPassword, in.Salt, false)
	db.UpdateUser(*u)
	return &pb.UserDBResponse{Email: u.GetEmail(), HashedPassword: u.GetHashedPassword(), Salt: u.GetSalt()}, nil
}

func main() {
	fmt.Println("Service Started")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserLogDBServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
