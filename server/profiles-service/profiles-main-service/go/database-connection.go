package main

import (
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)

type neo4jDB struct {
	context *neo4jDBContext
}

type neo4jDBContext struct {
	dbClient pb.ProfilesDBClient
	timeout time.Duration
}
var dbConn neo4jDB

// create connection
func newNeo4jDBContext(endpoint string) (*neo4jDBContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &neo4jDBContext{
		dbClient: pb.NewProfilesDBClient(userConn),
		timeout:  time.Second,
	}

	return ctx, nil
}
