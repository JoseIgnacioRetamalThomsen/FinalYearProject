//go:generate protoc -I ../Login --go_out=plugins=grpc:../Login ../Login/UserLogin.proto
//go:generate protoc -I . --go_out=plugins=grpc:. UserLogin.proto
//

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const (
	port = ":50051"

)

//first addrs is the master
var db_addrs = []string{"104.40.206.141:7777"}
var ps_addrs = []string{"52.236.146.149:5701"}


type server struct {
	pb.UnimplementedUserAuthenticationServer
}

// server
func (s *server) CheckUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	email, hash , salt ,err := getUser(in.GetEmail())
	if(err != nil){
		//user do not exist
		return &pb.UserResponse{IsUser: false}, errors.New("could not get user")

	}
	email= email

	isValid := validate(in.HashPassword,hash,salt)
	return &pb.UserResponse{IsUser: isValid, Cookie: "cookie"}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	hash, salt := hash(in.HashPassword)

	AddUser(in.Email,hash,salt)

	return &pb.UserResponse{IsUser: true, Cookie: "cookie"}, nil
}

func main() {

	//fmt.Print("helloworld")
	x,y := hash("helloworld678")
	//fmt.Print(x)
	//fmt.Print(y)

	fmt.Print(AddUser("myEmail1",x,y))

	//email,hash,salt,err := getUser("myEmail18")
	//if(err != nil){
		//user do not exits

	//}
	//fmt.Print(email)
	//hash=hash
	//salt=salt
	//fmt.Print(validate("helloworld",hash,salt))
	//email=email
	//x,y := hash("12344567")
	//updateUser("myEmail",x,y)

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

//hash service
func validate(pass string, hash []byte,salt []byte) bool{
	// Set up a connection to the server.
	conn, err := grpc.Dial(ps_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPasswordServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Validate(ctx, &pb.ValidateRequest{Password:pass, HasshedPassword: hash , Salt:salt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.Value
}

//hash service
func hash(pass string) ([]byte,[]byte){

	// Set up a connection to the server.
	conn, err := grpc.Dial(ps_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPasswordServiceClient(conn)

	// Contact the server and print out its response.


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hash(ctx, &pb.HashRequest{Password:pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}


	return r.GetHashedPassword(), r.GetSalt()
}


//db
func AddUser(email string, hashedPassword []byte, salt []byte) string{

	// Set up a connection to the server.
	conn, err := grpc.Dial(db_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserLogDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(ctx, &pb.UserDBRequest{Email: email, HashedPassword: hashedPassword,Salt:salt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetEmail()
	//log.Printf("Greeting: %s", r.GetMessage())
}

//db
func getUser(email string) (string,[]byte,[]byte,error){
	// Set up a connection to the server.
	conn, err := grpc.Dial(db_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserLogDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.UserDBRequest{Email: email})
	if err != nil {

		return "",nil,nil, errors.New("could not get user")
	}

	 return r.GetEmail(),r.GetHashedPassword(), r.GetSalt(),nil
}

//db
func updateUser(email string, hash []byte, salt []byte) string {
	// Set up a connection to the server.
	conn, err := grpc.Dial(db_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserLogDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUser(ctx, &pb.UserDBRequest{Email: email,HashedPassword:hash,Salt:salt})
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}
	return r.Email
}
