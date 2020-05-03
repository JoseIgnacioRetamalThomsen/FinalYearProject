package main

import (
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	//"net/http"

)

const(
	MongoDBURI = "mongodb://127.0.0.1:27017"
	//MongoDBURI = "mongodb://172.17.0.1:27017"
	//MongoDBURI = "mongodb://10.154.0.6:27017"
	DatabaseName ="PostDatabase"
	CollectionName  = "Posts"
	Port = ":2787"
)

type server struct {
	pb.UnimplementedPostsServiceDBAServer
}


func main(){
	fmt.Println("t")
/*
	 _, err := CreateCityPost(&CityPost{
		IndexId:      0,
		CreatorEmail: "d",
		Name:         "d",
		Country:      "d",
		Title:        "s",
		Body:         "d",
		TimeStamp:    "dsf",
		Likes:        nil,
		MongoId:      "cvcvxvxc",
	})

	 if err != nil{
	 	panic(err)
	 }*/

	fmt.Println(GetCityPost(1))

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServiceDBAServer(s, &server{})
	log.Print("Service Started in port: ", Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

