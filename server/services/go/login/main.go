//go:generate protoc -I ../Login --go_out=plugins=grpc:../Login ../Login/UserLogin.proto
//go:generate protoc -I . --go_out=plugins=grpc:. UserLogin.proto
//

package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"time"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const (
	port = ":50051"

)

const (
	exampleScheme      = "example"
	exampleServiceName = "lb.example.grpc.io"
)




//first addrs is the master
var db_addrs = []string{"104.40.206.141:7777"}
var ps_addrs = []string{"52.236.146.149:5701"}


type server struct {
	pb.UnimplementedUserAuthenticationServer
}

//conection
//https://stackoverflow.com/questions/56067076/grpc-connection-management-in-golang
// this type contains state of the server
type psserverContext struct {
	// client to GRPC service
	psClient pb.PasswordServiceClient

	// default timeout
	timeout time.Duration

	// some other useful objects, like config
	// or logger (to replace global logging)
	// (...)
}

// constructor for server context
func newClientContext(endpoint string) (*psserverContext, error) {
	userConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &psserverContext{
		psClient: pb.NewPasswordServiceClient(userConn),
		timeout: time.Second,
	}
	return ctx, nil
}

type serverPass struct {
	context *psserverContext
}
/*
func (s *server) Handler(ctx context.Context, request *Request) (*Response, error) {
	clientCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	response, err := c.GetUserFromTokenID(
		clientCtx,
		&user.GetUserFromTokenRequest{
			TransactionID: transactionID,
			OathToken: *oathToken,
		},
	)
	if err != nil {
		return nil, err
	}
	// ...
}

*/



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

// return false if user is updated
func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {

	hash, salt := hash(in.HashPassword)

	_, err := addUser(in.Email,hash,salt)

	if err!=nil{
		// if user exist we update
		_, err := updateUser(in.Email,hash,salt)
		if err!=nil{
			return nil, errors.New("Can't create or update.")
		}
		return &pb.UserResponse{IsUser: false, Cookie: "cookie"}, nil
	}
	return &pb.UserResponse{IsUser: true, Cookie: "cookie"}, nil
}

var psCon serverPass

func main() {

	// start server pass connection
	psserverCtx, err := newClientContext(ps_addrs[0])
	if err != nil {
		log.Fatal(err)
	}
	s1 := &serverPass{psserverCtx}
	psCon = *s1

	//fmt.Print("helloworld")
	x,y := hash("helloworld678")
	fmt.Print(x)
	fmt.Print(y)

	//fmt.Print(addUser("myEmail1",x,y))

	email,hash,salt,err := getUser("myEmail1")
	if(err != nil){
		//user do not exits

	}
	fmt.Print(email)
	hash=hash
	salt=salt
	fmt.Print(validate("helloworld",hash,salt))
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
/*
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
*/
func validate(pass string, hash []byte,salt []byte) bool{
	// Set up a connection to the server.


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := psCon.context.psClient.Validate(ctx, &pb.ValidateRequest{Password:pass, HasshedPassword: hash , Salt:salt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.Value
}
/*
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
*/
//hash service
func hash(pass string) ([]byte,[]byte){


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := psCon.context.psClient.Hash(ctx, &pb.HashRequest{Password:pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}


	return r.GetHashedPassword(), r.GetSalt()
}


//db
func addUser(email string, hashedPassword []byte, salt []byte) (string,error){

	// Set up a connection to the server.
	conn, err := grpc.Dial(db_addrs[0], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return email, errors.New("Cant connect.")
	}
	defer conn.Close()
	c := pb.NewUserLogDBClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(ctx, &pb.UserDBRequest{Email: email, HashedPassword: hashedPassword,Salt:salt})
	if err != nil {
		return r.GetEmail(), errors.New("Cant add.")
	}
	return r.GetEmail(),nil

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
func updateUser(email string, hash []byte, salt []byte) (string,error) {
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
		return "", errors.New("cant update")
	}
	return r.Email,nil
}
//resolver

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: db_addrs,
		},
	}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*exampleResolver) Close()                                  {}

func init() {
	resolver.Register(&exampleResolverBuilder{})
}

