package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"time"
)

//check token

// password service client
type authClient struct {
	context *authClientContext
}

type authClientContext struct {
	psClient pb.UserAuthenticationClient
	timeout time.Duration
}

//password service connection
var prsCon authClient


func newAuthClientContext(endpoint string) (*authClientContext, error) {
	authConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &authClientContext{
		psClient: pb.NewUserAuthenticationClient(authConn),
		timeout: time.Second,
	}
	return ctx, nil
}


func  CheckToken(email string, token string) bool{
	// Set up a connection to the server.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := prsCon.context.psClient.CheckToken(ctx,&pb.LogRequest{
		Token:                token,
		Email:                email,

	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return r.Sucess
}

