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

	"google.golang.org/grpc"
	pb "github.com/joseignacioretamalthomsen/wcity" 
)
 	


const (
	port = ":7777"
)


type server struct {
	pb.UnimplementedUserLogDBServer
}
/*
int32 id = 1;
string email = 2;
bytes hashedPassword = 3;
bytes salt = 4; 
AddUser(UserDBRequest) 
*/
// SayHello implements helloworld.GreeterServer
func (s *server) AddUser(ctx context.Context, in *pb.UserDBRequest) (*pb.UserDBResponse, error) {
	log.Printf("Received: %v", "Create new user")
	u := db.NewUser(in.Email,"",in.HashedPassword,in.Salt ,false)
	db.AddUser(*u)
	return &pb.UserDBResponse{Id:1,Email:"email"}, nil
}

func main() {

	//db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "", "test")


	fmt.Println("working")
 //u := db.NewUser("ema210il@gmail.com","name1",[]byte("Here is a string...."),[]byte("Here is a string...."),false)
 //fmt.Print( db.AddUser(*u))
  //GetPassSalt(u)
//  fmt.Println(GetUser("ema8il@gmail.com").email)
 db.DelUser("ema8il@gmail.com")
// UpdateUser(*u)
//ConfirmEmail("ema9il@gmail.com")
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




