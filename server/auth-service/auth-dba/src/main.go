// Provide access to mysql database using grpc interface
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joseignacioretamalthomsen/sqlgo/db"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// the database configuration.
type Configuration struct {
	Port            string
	Coneection_type string
	MySQL_socket    string
	MySQL_user      string
	MySQL_pass      string
	MySQL_db        string
}

// Database configuration.
var configuration Configuration

// The authentication server.
type server struct {
	pb.UnimplementedUserAuthDBServer
}

// Reads the configuration file.
func readConfig(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// End point:Add new user.
// Create a new user in the database. When the user creates the account, it will create
// the profile automatically in the profiles database.
func (s *server) AddUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v , from: %v", "Get user", in.String())
	u := db.NewUser(in.Email, in.PasswordHash, in.PasswordSalt, false)
	id, err := db.AddUser(*u)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &pb.UserDBResponse{Id: id, Email: in.Email, PasswordHash: in.PasswordHash, PasswordSalt: in.PasswordSalt}, nil
}

// End point: GEt user data.
// Returns the user data used to authenticate the user.
func (s *server) GetUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v , from: %v", "Get user", in.String())
	u, err := db.GetUser(in.Email)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.New("Can't get user")
	}
	return &pb.UserDBResponse{Id: u.GetId(), Email: u.GetEmail(), PasswordHash: u.GetHashedPassword(), PasswordSalt: u.GetSalt()}, nil
}

// End point: update user data.
//Update the user data, is used for changing the password.
func (s *server) UpdateUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Update user")
	u := db.NewUser(in.Email, in.PasswordHash, in.PasswordSalt, false)
	isUpdated, err := db.UpdateUser(*u)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.New("cant update")
	}
	isUpdated = isUpdated
	return &pb.UserDBResponse{Email: u.GetEmail(), PasswordHash: u.GetHashedPassword(), PasswordSalt: u.GetSalt()}, nil
}

// End point: create a new seassion.
// Create a new session in the database. Used when the user login using the password.
func (s *server) CreateSeassion(ctx context.Context, in *pb.UserSessionRequest) (*pb.UserSessionResponse, error) {

	key, email, d1, d2, err := db.CreateSession(in.Token, in.Email)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &pb.UserSessionResponse{Email: email, Token: key, LoginTime: d1, LastSeenTime: d2}, nil
}

// End point: Get a session.
// Return a user session if exist.  Used to check if the session exists so the user can
//log in without the password.
func (s *server) GetSeassion(ctx context.Context, in *pb.UserSessionRequest) (*pb.UserSessionResponse, error) {
	is, se, err := db.GetSession(in.Token)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	if is == false {
		log.Printf("Error: %v", "Token do not exists")
		return nil, errors.New("Token do not exist")
	}
	return &pb.UserSessionResponse{Token: se.SessionKey, Email: se.Email, LoginTime: se.LoginTime, LastSeenTime: se.LastSeemTime}, nil
}

// End point: Delete a session.
// Delete user session if exits. Used when the user logs out from the device.
func (s *server) DeleteSession(ctx context.Context, in *pb.UserSessionRequest) (*pb.UserDeleteSessionResponse, error) {
	res, err := db.DeleteSession(in.Token)
	if err != nil {
		log.Printf("Error: %v", err)
		return &pb.UserDeleteSessionResponse{Success: false}, err
	}
	if res <= 0 {
		return &pb.UserDeleteSessionResponse{Success: false}, err
	}
	return &pb.UserDeleteSessionResponse{Success: true}, nil
}

// Read config file, setup database connection and start server.
func main() {
	//read config file name from console input
	args := os.Args[1]
	readConfig(args)

	db.SetupConnection(configuration.Coneection_type,
		configuration.MySQL_socket,
		configuration.MySQL_user,
		configuration.MySQL_pass,
		configuration.MySQL_db)
	log.Print("Starting Service")

	lis, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserAuthDBServer(s, &server{})
	log.Print("Service Started in port: ", configuration.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
