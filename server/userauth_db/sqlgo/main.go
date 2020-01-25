package main

import (
	"encoding/json"
	"errors"
	"os"

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
    Coneection_type = "tcp"
     MySQL_socket = "127.0.0.1:3306"
     MySQL_user = "goland"
    MySQL_pass = "fg453ty#2334Mx"
     MySQL_db = "UserAuth"
)

type Configuration struct {
	Port string
	Coneection_type    string
	MySQL_socket string
	MySQL_user string
	MySQL_pass string
	MySQL_db string
}

var configuration Configuration
func readConfig(fileName string){
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

type server struct {
	pb.UnimplementedUserAuthDBServer
}

func (s *server) AddUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Create new user")
	u := db.NewUser(in.Email, in.PasswordHash, in.PasswordSalt, false)
	added := db.AddUser(*u)
	if added == false {
		return nil, errors.New("Can't add")
	}
	return &pb.UserDBResponse{Id: 10, Email: in.Email, PasswordHash: in.PasswordHash, PasswordSalt: in.PasswordSalt}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Get user")
	u, err := db.GetUser(in.Email)

	if err != nil {
		return nil, errors.New("Can't get user")
	}
	return &pb.UserDBResponse{Email: u.GetEmail(), PasswordHash: u.GetHashedPassword(), PasswordSalt: u.GetSalt()}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Update user")
	u := db.NewUser(in.Email, in.PasswordHash, in.PasswordSalt, false)
	isUpdated, err := db.UpdateUser(*u)
	if(err!= nil){
		return nil, errors.New("cant update")
	}
	isUpdated = isUpdated
	return &pb.UserDBResponse{Email: u.GetEmail(), PasswordHash: u.GetHashedPassword(), PasswordSalt: u.GetSalt()}, nil
}

func main() {

	readConfig("config.json")

	db.SetupConnection(configuration.Coneection_type,
		configuration.MySQL_socket,
		configuration.MySQL_user,
		configuration.MySQL_pass,
		configuration.MySQL_db)

	log.Print("Starting Service")

	//user11 := db.NewUser("emailui", []byte("passs1"),[]byte("salt"),false)
	//db.AddUser(*user11)
	//u1,err := db.GetUser("emailui")
	//if err!= nil{
	//	panic(err)
	//}
//	fmt.Print(u1.GetEmail())
	//db.UpdateUser(*user11)
	_,se := db.GetSession("sdfads4fadfaeq443q34qf304")
	fmt.Print(se.SessionKey)
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
