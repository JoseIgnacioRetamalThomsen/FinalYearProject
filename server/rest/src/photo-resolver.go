package main

import (
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)

type photosServer struct {
	context *photosServiceContext
}

type photosServiceContext struct {
	dbClient pb.PhotosServiceClient
	timeout time.Duration
}
var photoConn photosServer

// create connection
func newPhotosServiceContext(endpoint string) (*photosServiceContext, error) {

	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &photosServiceContext{
		dbClient: pb.NewPhotosServiceClient(userConn),
		timeout:  time.Second*2,
	}
	return ctx, nil
}

