package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)



type authService struct {
	context *authServiceContext
}

type authServiceContext struct {
	dbClient pb.UserAuthenticationClient
	timeout time.Duration
}
var authConn authService

// create connection
func newAuthServiceContext(endpoint string) (*authServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &authServiceContext{
		dbClient: pb.NewUserAuthenticationClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}


// Create User, user load balancing connection.
func createUser(user pb.UserRequest) (error, *pb.UserResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := dbConnLB.context.dbClient.CreateUser(ctx, &user)
	if err != nil {

		fmt.Println(err)
		return  errors.New("could not create"),nil
	}

	return   nil,r
}

// Create User, user load balancing connection.
func updateUser(user pb.UserRequest) (error, *pb.UserResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConnLB.context.dbClient.UpdateUser(ctx, &user)
	if err != nil {

		fmt.Println(err)
		return  errors.New("could not uptade."),nil
	}

	return   nil,r
}

//GetUser
func loginUser(user pb.UserRequest) (error, *pb.UserResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConnLB.context.dbClient.LoginUser(ctx, &user)
	if err != nil {

		fmt.Println(err)
		return  errors.New("could not get user"),nil
	}

	return   nil,r
}


//GetUser
func checkUsertoken(request pb.LogRequest) (error, *pb.LogResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConnLB.context.dbClient.CheckToken(ctx, &request)
	if err != nil {

		fmt.Println(err)
		return  errors.New("Could not check the token."),nil
	}

	return   nil,r
}

func logoutUser(request pb.LogRequest) (error, *pb.LogResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConnLB.context.dbClient.Logout(ctx, &request)
	if err != nil {

		fmt.Println(err)
		return  errors.New("Could not log out."),nil
	}

	return   nil,r
}
