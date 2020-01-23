package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"time"
)

// password service client
type clientPassword struct {
	context *passClientContext
}

type passClientContext struct {
	psClient pb.PasswordServiceClient
	timeout time.Duration
}

//password service connection
var psCon clientPassword
// database connection
var dbConn clientDB
var dbConnLB clientDBLoadBalancing

func newClientContext(endpoint string) (*passClientContext, error) {
	userConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &passClientContext{
		psClient: pb.NewPasswordServiceClient(userConn),
		timeout: time.Second,
	}
	return ctx, nil
}


//hash service
func  validate(pass string, hash []byte,salt []byte) bool{
	// Set up a connection to the server.


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := psCon.context.psClient.Validate(ctx, &pb.ValidateRequest{Password:pass, HasshedPassword: hash , Salt:salt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.Value
}

//hash service
func hash(pass string) ([]byte,[]byte){


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := psCon.context.psClient.Hash(ctx, &pb.HashRequest{Password:pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}


	return r.GetHashedPassword(), r.GetSalt()
}
