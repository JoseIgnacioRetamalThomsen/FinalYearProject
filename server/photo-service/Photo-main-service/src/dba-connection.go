package main

import (
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)

const (
	MAX_DB_CON_TIME = 3
)

type server struct {
	pb.UnimplementedPhotosServiceServer
}

var dbaConn photosDBAServer

type photosDBAServiceContext struct {
	dbClient pb.PhotosDBAServiceClient
	timeout  time.Duration
}
type photosDBAServer struct {
	context *photosDBAServiceContext
}

func newPhotosDBAServiceContext(endpoint string) (*photosDBAServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &photosDBAServiceContext{
		dbClient: pb.NewPhotosDBAServiceClient(userConn),
		timeout:  2 * time.Second,
	}
	return ctx, nil
}

