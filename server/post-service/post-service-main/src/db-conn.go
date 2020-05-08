package main

import (
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)

// dba connection
type postDBA struct {
	context *postDBAContext
}

type postDBAContext struct {
	dbClient pb.PostsServiceDBAClient
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
		dbClient: pb.NewPostsServiceDBAClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}
