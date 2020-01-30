package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"

	//"errors"
	//"fmt"
	//pb "github.com/joseignacioretamalthomsen/wcity"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/resolver"
	//"time"
)
const(
	url = "0.0.0.0:5777"
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


func main(){
	dbserverCtx, err := newNeo4jDBContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &neo4jDB{dbserverCtx}
	dbConn = *s2

	CreateUser("one","two","three")
}

func CreateUser(email string,name string, description string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateUser(ctx, &pb.CreateUserRequest{Email: email, Name: name,Description:description})
	if err != nil {
		return false, err
	}
	res, err := strconv.ParseBool(r.Valied)

	return res,nil
}
