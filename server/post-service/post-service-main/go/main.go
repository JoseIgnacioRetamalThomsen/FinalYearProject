package main

import (
	"log"
	"time"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
)

const(
	//url = "0.0.0.0:60051"
	url="35.197.216.42:60051";
	//url = "35.234.146.99:5777"
	token ="a31e31a2fcdf2a9a230120ea620f3b24f7379d923fb122323d3cb9bc56fe6508"
	tokenEmail ="a@a.com"
)
// dba connection
type postDBA struct {
	context *postDBAContext
}

type postDBAContext struct {
	dbClient pb.PostsServiceClient
	timeout time.Duration
}
var dbaConn postDBA

// create connection
func newPostServiceContext(endpoint string) (*postDBAContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &postDBAContext{
		dbClient: pb.NewPostsServiceClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}


//server


func main(){
	//conect to server
	dbserverCtx, err := newPostServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &postDBA{dbserverCtx}
	dbaConn = *s2
}
