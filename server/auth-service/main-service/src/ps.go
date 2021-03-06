// main package
package main

// Access to profile service, create user in neo4j database when user is created.

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"time"
)

// the profile service.
type profileServer struct {
	context *profileServiceContext
}

// Profile service connection  struct.
type profileServiceContext struct {
	prosClient pb.ProfilesClient
	timeout    time.Duration
}

// The profile conection.
var profSerConn profileServer

// create connection
func newProfilesServiceContext(endpoint string) (*profileServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &profileServiceContext{
		prosClient: pb.NewProfilesClient(userConn),
		timeout:    time.Second*MAX_CON_TIME,
	}
	return ctx, nil
}

// create user in profile service.
func CreateUser(email string, token string, name string, description string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	r, err := profSerConn.context.prosClient.CreateUser(ctx, &pb.CreateUserRequestP{
		Token:       token,
		Email:       email,
		User: &pb.User{
			Email:                email,
			Name:                 name,
			Descripiton:          description,
			UserId:               0,

		},
	})
	if err != nil {
		log.Printf("Received: %v", err)
		return false
	}

	return r.Valid
}
