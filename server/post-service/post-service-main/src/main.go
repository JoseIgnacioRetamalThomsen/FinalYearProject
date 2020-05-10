package main

import (

	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

//server
type server struct {
	pb.UnimplementedPostsServiceServer
}

func main(){

	// read configuration
	args := os.Args[1]
	readConfig(args)

	fmt.Println(time.Now().Format("Mon Jan _2 15:04:05 MST 2006"))
	//conect to server
	dbserverCtx, err := newPostServiceContext(configuration.Dbs[0])
	if err != nil {
		log.Fatal(err)
	}
	s2 := &postDBA{dbserverCtx}
	dbaConn = *s2



	lis, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServiceServer(s, &server{})
	log.Print("Service Started in port: ", configuration.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


