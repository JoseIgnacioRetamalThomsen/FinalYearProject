package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"time"
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

	//CreateCityPost()
	//CreatePlacePost()
	GetPlacePosts()
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

func CreatePlacePost(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.CreatePlacePost(ctx,&pb.PlacePost{
		IndexId:              2,
		CreatorEmail:         "xxxxxxx",
		CityName:             "xxxxx",
		CountryName:          "xxxx",
		PlaceName:            "xxxxx",
		Title:                "xxxxx",
		Body:                 "xxxx",

		Likes:                nil,
		MongoId:              "",

	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func GetPlacePosts(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.GetPlacePosts(ctx,&pb.PostsRequest{
		IndexId:              1,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	err=err
	fmt.Println(r)
}
