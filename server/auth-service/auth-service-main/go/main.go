//go:generate protoc -I ../Login --go_out=plugins=grpc:../Login ../Login/UserLogin.proto
//go:generate protoc -I . --go_out=plugins=grpc:. UserLogin.proto
//

package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"

)

const (
	dbConnectionScheme = "dbConnectionScheme"
	exampleServiceName = "ie.gmit.wcity.auth"
)

type Configuration struct {
	Port string
	Dbs    []string
	Pss   []string
}


//first addrs is the master
//var db_addrs = []string{"104.40.206.141:7777","40.118.90.61:7777"}
//var ps_addrs = []string{"40.118.90.61:5701","51.124.149.63:5701"}

//var ps_addrs = []string{"35.197.216.42:5701","localhost:5701"}
//grpc server
type server struct {
	pb.UnimplementedUserAuthenticationServer
}

/*
Init client connections
*/

//https://stackoverflow.com/questions/56067076/grpc-connection-management-in-golang
// this type contains state of the server
/**
*  End Points
 */
func (s *server) LoginUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {


	log.Printf("Received: %v", "LoginUser  called")
	//get user data from database
	email, hash , salt ,err := getUser(in.GetEmail())
	if(err != nil){
		//user do not exist
		return &pb.UserResponse{IsUser: false}, err
	}
	email= email
	// user hash service for check password
	isValid := validate(in.HashPassword,hash,salt)
	if isValid {
		token := GenerateSecureToken(32)
		is, err := CreateSession(in.Email, token)
		if err != nil {
			return nil, err
		}

		return &pb.UserResponse{IsUser : is ,Token:token},nil
	}
	return &pb.UserResponse{IsUser: false, Token: ""}, nil
}

// return false if user is updated
func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {


	log.Printf("Received: %v", "create user")
	if(len(in.HashPassword) <=6){
		return nil, errors.New("Password to short")
	}
	hash, salt := hash(in.HashPassword)

	email,id, err := addUser(in.Email,hash,salt)
	email =email
	id = id
	if err!=nil{
		return nil, err
	}
	token := GenerateSecureToken(32)
	is, err := CreateSession(in.Email, token)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{ IsUser: is, Token: token}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error){
	log.Printf("Received: %v", "Update user")
	hash, salt := hash(in.HashPassword)
	// if user exist we update
	email,pass,salt, err := updateUser(in.Email,hash,salt)

	pass =pass
	salt =salt
	if err!=nil{
		return nil, errors.New("Can't create or update.")
	}
	token := GenerateSecureToken(32)
	is, err := CreateSession(email, token)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{IsUser: is, Token: token}, nil
}

func (s *server) CheckToken (ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error){
	log.Printf("Received: %v", "Check token")
	is,err := CheckToken(in.Email,in.Token)
	if err !=nil{
		return &pb.LogResponse{Sucess:false},nil
	}
	return &pb.LogResponse{Sucess:is},nil
}

func (s *server) Logout(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v", "Logout")
    suc,err := DeleteToken(in.Email,in.Token)
    if err != nil{
    	return nil,err
	}
	return &pb.LogResponse{Sucess:suc},nil
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}



//init resolvers
func init() {
	resolver.Register(&databasesResolverBuilder{})
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

func main() {

	args := os.Args[1]
	fmt.Print(args)
	readConfig(args)
	fmt.Println(configuration.Pss)

	// start server pass connection
	psserverCtx, err := newClientContext(configuration.Pss[0])
	if err != nil {
		log.Fatal(err)
	}
	s1 := &clientPassword{psserverCtx}
	psCon = *s1

	//start db client
	dbserverCtx, err := newDBContext(configuration.Dbs[0])
	if err != nil {
		log.Fatal(err)
	}
	s2 := &clientDB{dbserverCtx}
	dbConn = *s2

	//start load balance connection
	dbserverCtxLB, err := newDBContextLoadBalancing()
	if err != nil {
		log.Fatal(err)
	}
	s3 := &clientDBLoadBalancing{dbserverCtxLB}
	dbConnLB = *s3

	//fmt.Print("helloworld")
	//x,y := hash("helloworld678")
     //fmt.Print(x,y)
	//a,b,c := addUser("7777",x,y)
	//fmt.Print(a,b,c)
	//email,pass,salt,err :=getUser("email756")
	//err=err
	//fmt.Print(email,pass,salt)





//	email,hash,salt,err := getUser("myEmail1")
	//if(err != nil){
		//user do not exits

//	}
	//fmt.Print(email)
	//hash=hash
	//salt=salt
//	fmt.Print(validate("helloworld",hash,salt))
	//email=email
	//x,y := hash("12344567")
	//updateUser("Emailu3i",x,y)

	//fmt.Print("Service started")

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




