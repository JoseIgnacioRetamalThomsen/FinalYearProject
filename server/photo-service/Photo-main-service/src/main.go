package main

import (
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"net"

	"log"


)



const (
	port = ":30051"
)

const (
	DBA_URL = "35.197.221.57:7172"

	//DBA_URL  = "0.0.0.0:7172"
	AUTH_URL = "3.85.62.10:50051"
)


func main() {

	//conect to database server
	dbserverCtx, err := newPhotosDBAServiceContext(DBA_URL)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &photosDBAServer{dbserverCtx}
	dbaConn = *s2

	//conect to ps
	// start server pass connection
	psserverCtx, err := newAuthClientContext(AUTH_URL)
	if err != nil {
		log.Fatal(err)
	}
	s1 := &authClient{psserverCtx}
	prsCon = *s1

	//start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPhotosServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}


