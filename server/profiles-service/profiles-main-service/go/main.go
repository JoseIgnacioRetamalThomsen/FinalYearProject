package main

import (
	"net"
	"os"

	//"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"

	//"errors"
	//"fmt"
	//pb "github.com/joseignacioretamalthomsen/wcity"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/resolver"
	//"time"
)

//server for client to connect
type server struct {
	pb.UnimplementedProfilesServer
}

func main() {

	// read configuration
	args := os.Args[1]
	readConfig(args)

	//conect to neo4j db
	dbserverCtx, err := newNeo4jDBContext(configuration.Database[0])
	if err != nil {
		log.Fatal(err)
	}
	s2 := &neo4jDB{dbserverCtx}
	dbConn = *s2

	//conect to ps
	// start auth servcie connection
	psserverCtx, err := newAuthClientContext(configuration.Auth[0])
	if err != nil {
		log.Fatal(err)
	}
	s1 := &authClient{psserverCtx}
	prsCon = *s1


	//start server
	lis, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfilesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
