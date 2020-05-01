package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"time"
)


type profileServer struct {
	context *ProfileServiceContext
}

type ProfileServiceContext struct {
	dbClient pb.ProfilesClient
	timeout time.Duration
}
var ProfSerConn profileServer

// create connection
func newProfilesServiceContext(endpoint string) (*ProfileServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &ProfileServiceContext{
		dbClient: pb.NewProfilesClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

func CreateCity(city pb.CreateCityRequestP)(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.CreateCity(ctx,&city)
	if err != nil{
		return nil,err
	}
	return r,nil
}




func GetCity(city pb.GetCityRequestP )(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetCity(ctx,&city)

	if err != nil{
		return nil,err
	}

	fmt.Println(r.City)
	fmt.Println(r)
	return r,nil
}

func UpdateCity(city pb.CreateCityRequestP)(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.UpdateCity(ctx,&city)
	if err != nil{
		return nil,err
	}
	return r,nil
}

func CreatePlace(request pb.CreatePlaceRequestP)(*pb.PlaceResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.CreatePlace(ctx,&request)
	if err != nil{
		return nil,err
	}

	return r,nil
}
