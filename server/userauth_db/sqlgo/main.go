package main

import (
	"errors"

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
     MySQL_user = "golandAcces"
    MySQL_pass = "supedPss"
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
func readConfig(fileName string){/*
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}*/
configuration.Port = port
configuration.MySQL_db = MySQL_db
configuration.Coneection_type = Coneection_type
configuration.MySQL_socket = MySQL_socket
configuration.MySQL_user = MySQL_user
configuration.MySQL_pass = MySQL_pass
}

type server struct {
	pb.UnimplementedUserAuthDBServer
}

func (s *server) AddUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Create new user")
	u := db.NewUser(in.Email, in.PasswordHash, in.PasswordSalt, false)
	id,err := db.AddUser(*u)
	if err!=nil {
		return nil, err
	}
	return &pb.UserDBResponse{Id: id, Email: in.Email, PasswordHash: in.PasswordHash, PasswordSalt: in.PasswordSalt}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Get user")
	u, err := db.GetUser(in.Email)

	if err != nil {
		return nil, errors.New("Can't get user")
	}
	return &pb.UserDBResponse{Id: u.GetId(),Email: u.GetEmail(), PasswordHash: u.GetHashedPassword(), PasswordSalt: u.GetSalt()}, nil
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

func (s *server) CreateSeassion(ctx context.Context,in *pb.UserSessionRequest)  (*pb.UserSessionResponse,error){
	key,email,d1,d2,err :=db.CreateSession(in.Token,in.Email)
	if err!= nil{
		return nil, err
	}
	return &pb.UserSessionResponse{Email:email,Token:key,LoginTime:d1,LastSeenTime:d2},nil
}

func (s *server) GetSeassion(ctx context.Context,in *pb.UserSessionRequest)  (*pb.UserSessionResponse,error){
	is,se,err := db.GetSession(in.Token)
	if err!= nil{
		return nil, err
	}
	if is == false{
		return nil, errors.New("Token do not exist")
	}
	return &pb.UserSessionResponse{Token:se.SessionKey,Email:se.Email,LoginTime:se.LoginTime,LastSeenTime:se.LastSeemTime}, nil
}

func (s *server) DeleteSession(ctx context.Context,in *pb.UserSessionRequest)  (*pb.UserDeleteSessionResponse,error){
	res,err := db.DeleteSession(in.Token)
	if err!=nil{
		return &pb.UserDeleteSessionResponse{Success:false},err
	}
	if res<=0{
		return &pb.UserDeleteSessionResponse{Success:false},err
	}
	return &pb.UserDeleteSessionResponse{Success:true},nil
}

func main() {

	readConfig("config.json")

	db.SetupConnection(configuration.Coneection_type,
		configuration.MySQL_socket,
		configuration.MySQL_user,
		configuration.MySQL_pass,
		configuration.MySQL_db)

	log.Print("Starting Service")

	//key,email,d1,d2,err :=db.CreateSession("adefadbf23ffeg","email31")
	//if err != nil{
	//panic(err)}
	//fmt.Print(key,email,d1,d2)


	//user11 := db.NewUser("emailui7y", []byte("passs1"),[]byte("salt"),false)
	//db.AddUser(*user11)
	//u1,err := db.GetUser("emailui")
	//if err!= nil{
	//	panic(err)
	//}
//	fmt.Print(u1.GetEmail())
	//db.UpdateUser(*user11)
	//_,se,err := db.GetSession("sdfads4fadfaeq443q34qf304")
	//err= err
	//fmt.Print(se.SessionKey)
	//db.DeleteSession("sdfads4fadfaeq443q34qf304")
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
