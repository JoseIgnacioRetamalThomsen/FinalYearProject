package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const (
	url ="0.0.0.0:10051"
)

type postService struct {
	context *postServiceContext
}

type postServiceContext struct {
	dbClient pb.PostsServiceClient
	timeout time.Duration
}
var serviceConn postService

// create connection
func newPostServiceContext(endpoint string) (*postServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &postServiceContext{
		dbClient: pb.NewPostsServiceClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

func main(){

	//conect to server
	dbserverCtx, err := newPostServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &postService{dbserverCtx}
	serviceConn = *s2

	CreateCityPost()
}

func CreateCityPost(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.CreateCityPost(ctx,&pb.CityPost{
		IndexId:              1,
		CreatorEmail:         "aaaaa",
		CityName:             "bbbb",
		CityCountry:          "cccc",
		Title:                "ddd",
		Body:                 "eee",
		TimeStamp:            "ffff",

		MongoId:              "",

	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
