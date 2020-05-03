package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const(
	//url = "0.0.0.0:2787"
	//url="35.197.216.42:60051";
	url = "35.197.221.57:2787"
	token ="a31e31a2fcdf2a9a230120ea620f3b24f7379d923fb122323d3cb9bc56fe6508"
	tokenEmail ="a@a.com"
	Port = ":10051"
)



//server
type server struct {
	pb.UnimplementedPostsServiceServer
}

func main(){

	fmt.Println(time.Now().Format("Mon Jan _2 15:04:05 MST 2006"))
	//conect to server
	dbserverCtx, err := newPostServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &postDBA{dbserverCtx}
	dbaConn = *s2

	temp := &pb.CityPostPSDB{
		IndexId:              1,
		CreatorEmail:         "email",
		CityName:             "aaaaaaaaaaaaaaaaaaaaa",
		CityCountry:          "b",
		Title:                "c",
		Body:                 "d",
		TimeStamp:            "e",
		Likes:                nil,
		MongoId:              "",

	}
	CreateCityPost(*temp)

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServiceServer(s, &server{})
	log.Print("Service Started in port: ", Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


