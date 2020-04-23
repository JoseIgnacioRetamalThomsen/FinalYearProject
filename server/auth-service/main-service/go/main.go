// main package provide end points.
package main

//Provide main service for authentication.
// React native app and rest api use endpoint for provide authentication.
// User dba, profile service and hash service.

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

// the public port.
const (
	port = ":50051"
)

const MAX_CON_TIME =30;
// configuration
const (
	dbConnectionScheme = "dbConnectionScheme"
	exampleServiceName = "ie.gmit.wcity.auth"
)

// configuration
type Configuration struct {
	Port string
	Dbs  []string
	Pss  []string
	Prof [] string
}

// the server.
type server struct {
	pb.UnimplementedUserAuthenticationServer
}

// End Point:
func (s *server) LoginUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	log.Printf("Received: %v", "LoginUser  called")
	//get user data from database
	email, hash, salt, err := getUser(in.GetEmail())
	if err != nil {
		//user do not exist
		return &pb.UserResponse{IsUser: false}, status.Error(codes.NotFound, "user not found")
	}
	email = email
	// user hash service for check password
	isValid := validate(in.HashPassword, hash, salt)
	if isValid {
		token := GenerateSecureToken(32)
		is, err := CreateSession(in.Email, token)
		if err != nil {
			return nil, status.Error(codes.Internal, "db problem")
		}

		return &pb.UserResponse{IsUser: is, Token: token}, nil
	}
	return &pb.UserResponse{IsUser: false, Token: ""}, nil
}

// End Point:
// return false if user is updated
func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	log.Printf("Received: %v", "create user")
	if len(in.HashPassword) <= 6 {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}
	hash, salt := hash(in.HashPassword)

	email, id, err := addUser(in.Email, hash, salt)

	id = id
	if err != nil {
		return nil, status.Error(codes.Internal, "database problem")
	}
	token := GenerateSecureToken(32)
	is, err := CreateSession(in.Email, token)
	if err != nil {
		return nil, status.Error(codes.Internal, "db problem")
	}

	// add to profiles db
	CreateUser(email, token, "Plese input your name", "Your description")
	/*if res == false{
		log.Printf("can create: %v", err)
		return nil,status.Error(codes.Internal,"cant create")
	}
	*/

	return &pb.UserResponse{IsUser: is, Token: token}, nil
}

// End Point:  Update user.
func (s *server) UpdateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: %v", "Update user")
	hash, salt := hash(in.HashPassword)
	// if user exist we update
	email, pass, salt, err := updateUser(in.Email, hash, salt)

	pass = pass
	salt = salt
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Can't create or update.")
	}
	token := GenerateSecureToken(32)
	is, err := CreateSession(email, token)
	if err != nil {
		return nil, status.Error(codes.Internal, "database problem")
	}

	return &pb.UserResponse{IsUser: is, Token: token}, nil
}

// End Point:  Check token
func (s *server) CheckToken(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v", "Check token")
	is, err := CheckToken(in.Email, in.Token)
	if err != nil {
		return &pb.LogResponse{Sucess: false}, nil
	}
	return &pb.LogResponse{Sucess: is}, nil
}

// End Point: Log out a user, basically delete the token from database
func (s *server) Logout(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	log.Printf("Received: %v", "Logout")
	suc, err := DeleteToken(in.Email, in.Token)
	if err != nil {
		return nil, err
	}
	return &pb.LogResponse{Sucess: suc}, nil
}

// Generate a token.
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

//init resolvers.
func init() {
	resolver.Register(&databasesResolverBuilder{})
}

// Server configuration.
var configuration Configuration

// reads config file.
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

//Acces to services and provide server.
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

	// conect to profiles

	profilesCtx, err := newProfilesServiceContext(configuration.Prof[0])
	if err != nil {
		log.Fatal(err)
	}
	s4 := &profileServer{profilesCtx}
	profSerConn = *s4


	CreateUser("use45","token","name","description");
	//fmt.Print("Service started")
	log.Printf("Started: %v", " service")

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
