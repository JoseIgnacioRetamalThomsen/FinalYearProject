package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"time"
)

type profileServer struct {
	context *profileServiceContext
}

type profileServiceContext struct {
	prosClient pb.ProfilesClient
	timeout time.Duration
}
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
		timeout:  time.Second,
	}
	return ctx, nil
}

func CreateUser(email string,token string,name string,description string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.prosClient.CreateUser(ctx,&pb.UserRequestP{
		Token:                token,
		Email:                email,
		Name:                 name,
		Description:          description,

	})
	if err != nil{
		log.Printf("Received: %v", err)
		return false
	}

	return r.Valid
}
