package main

import (
	//"os"
	//	"bytes"
	//"fmt"

	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine

	"context"
	"fmt"

	//"net"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	//"os"
	"time"
)

const (
//	address     = "35.197.216.42:7777"
	//address     = "40.118.90.61:7777"
	//address     = "localhost:7777"
address = "104.40.206.141:7777"
)

func main() {

	//fmt.Print(CreateUser("e5613tfhhhh", []byte("12345h6"),[]byte("12334h6")))
	//GetUser("e5613tfh")
//	UpdateUser("e560",[]byte("12345678"),[]byte("123345678"))
	//CreateSession("e560","1234556789")
	fmt.Print(GetSession("e560","12345567898"))
	//fmt.Print(DeleteSession("e560","1234556"))

}

func GetUser(email string){
	// Set up a connection to the server.
	log.Println("meesage")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("meesage")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)

	log.Println("meesage")
	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.UserDBRequest{Email: email})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetEmail())
}

func CreateUser(email string, passwordHash []byte,passwordSalt []byte)string{


	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(ctx, &pb.UserDBRequest{Email: email, PasswordHash: passwordHash,PasswordSalt:passwordSalt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Email

}

func UpdateUser(email string, passwordHash []byte,passwordSalt []byte) string{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUser(ctx, &pb.UserDBRequest{Email: email, PasswordHash: passwordHash,PasswordSalt:passwordSalt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Email
}

func CreateSession(email string,token string)string{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateSeassion(ctx, &pb.UserSessionRequest{Email: email, Token:token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Token

}

func GetSession(email string,token string)string{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetSeassion(ctx, &pb.UserSessionRequest{Email: email, Token:token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Email

}
func DeleteSession(email string,token string)bool{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthDBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteSession(ctx, &pb.UserSessionRequest{Email: email, Token:token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Success

}
/*
func CreateUser(email string, password string)string{


	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthenticationClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.UserRequest{Email: email, HashPassword: password})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.GetToken()
	//log.Printf("Greeting: %s", r.GetMessage())
}

func LoginUser(email string, password string) string{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthenticationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.LoginUser(ctx, &pb.UserRequest{Email: email, HashPassword: password})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Token
}

func CheckToken(email string,token string)bool{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthenticationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CheckToken(ctx, &pb.LogRequest{Email: email, Token:token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Sucess

}

func Logout(email string, token string)bool{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthenticationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Logout(ctx, &pb.LogRequest{Email: email, Token:token})
	return r.Sucess
}
func Update(email string, pass string) bool{
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserAuthenticationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUser(ctx, &pb.UserRequest{Email: email, HashPassword:pass})
	return r.IsUser
}
/*

func GetUser(email string){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetEmail())
}

func UpdateUser(email string, hashedPassword []byte, salt []byte){

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserLogDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateUser(ctx, &pb.UserDBRequest{Email: email, HashedPassword: hashedPassword,Salt:salt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	r=r
	//log.Printf("Greeting: %s", r.GetMessage())
}
*/
