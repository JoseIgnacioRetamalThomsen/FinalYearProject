package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"io"
	"log"
	"strings"
	"time"
)

const(
	CON_DEADLINE = 30;
)
func GetPlacesCity(request pb.CityRequestPDB)(*pb.VisitedPlacesResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetCityPlaces(ctx,&request)
	if err != nil {
		return nil,err
	}
	return r,nil
}

func UpdatePlace(place pb.Place)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdatePlaceRequest(ctx,&place)
	if err != nil {
		return false,err
	}
	return r.Valid,nil
}

func UpdateCity(city pb.City)(bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateCityRequest(ctx,&city)
	if err != nil {
		return false,nil
	}
	return r.Valid,nil

}

func UpdateUser(user pb.User)(*pb.CreateUserResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateUserRequest(ctx, &user)
	if err != nil {
		return nil,err
	}

	return &pb.CreateUserResponsePDB{
		Valid:                r.Valid,
		User:                 &user,
		},nil
}
func GetVisitedCitys(email string)(*pb.VisitedCitysResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetVisitedCitys(ctx,&pb.VisitedCitysRequestPDB{
		Email:           strings.ToLower(email)})
	if err !=nil{
		return nil,err
	}
	return r, nil
}

func VisitCity(request pb.VisitCityRequestPDB)(*pb.VisitCityResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.VisitCity(ctx,&request)
	if err !=nil{
		return nil,err
	}

	return r,nil

}

func GetVisitedPlaces(email string)(*pb.VisitedPlacesResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetVisitedPlaces(ctx,&pb.VisitedPlacesRequestPDB{
		Email:  strings.ToLower(email)})
	if err != nil {
		return nil,err
	}

	return r,nil
}

func VisitPlace(request pb.VisitPlaceRequestPDB)(*pb.VisitPlaceResponsePDB,error){

	ctx, cancel := context.WithTimeout(context.Background(),CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.VisitPlace(ctx,&request)
	if err != nil {
		return nil,err
	}

	return r,nil
}

func GetPlace(request pb.PlaceRequestPDB)(*pb.Place,error){

	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetPlace(ctx,&request)
	if err != nil {
		return nil,err
	}

	return r,nil
}

func CreatePlace(place pb.Place)(*pb.PlaceResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreatePlaceRequest(ctx,&place)
	if err != nil {
		return nil,err
	}
	return r, nil
}

func GetCity(request pb.CityRequestPDB)(*pb.CityResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetCity(ctx,&request)
	if err != nil {
		return nil, err
	}
	return r,nil
}

func CreateCity(city pb.City) (*pb.CityResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateCity(ctx, &city)
	if err != nil {
		return nil,err
	}
	return r,nil
}

func CreateUser(user pb.User) (*pb.CreateUserResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}


	return &pb.CreateUserResponsePDB{
		Valid:                r.Valid,
		User:                 r.User,
	},nil
}

func GetUser(request pb.GetUserRequestPDB)(*pb.UserResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()

	r, err := dbConn.context.dbClient.GetUser(ctx, &request)
	if err != nil {
		return nil,err
	}
	return &pb.UserResponseP{
		Valid:                r.Valid,
		User:                 r.User,
			},nil
}

//structs to pass in channels from dba to main
// need for pass the error
type CityResult struct{
	City *pb.City
	Err error
}
type PlaceResult struct{
	Place *pb.Place
	Err error
}

func GetAllCitysDBA(in *pb.GetAllRequest,c chan CityResult)(*pb.UserResponseP,error) {

	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()

	stream, err := dbConn.context.dbClient.GetAllCitysDBA(ctx, &pb.GetAllRequest{Max: 100})

	if err != nil {
		panic(err)
	}
	for {
		city, err := stream.Recv()
		if err == io.EOF {

			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", ctx, err)
			c <- CityResult{
				City: nil,
				Err:   err,
			}
		}
		log.Println(city)
		c <- CityResult{
			City: city,
			Err:   nil,
		}

	}
	close(c)
	return nil,nil
}


func GetAllPlacesDBA(in *pb.GetAllRequest,c chan PlaceResult)(*pb.UserResponseP,error) {

	ctx, cancel := context.WithTimeout(context.Background(), CON_DEADLINE*time.Second)
	defer cancel()

	stream, err := dbConn.context.dbClient.GetAllPlacesDBA(ctx, &pb.GetAllRequest{Max: 100})

	if err != nil {
		panic(err)
	}
	for {
		place, err := stream.Recv()
		if err == io.EOF {

			break
		}
		if err != nil {
			log.Fatalf("%v.AllPLaces = _, %v", ctx, err)
			c <- PlaceResult{
				Place: nil,
				Err:   err,
			}
		}

		c <- PlaceResult{
			Place: place,
			Err:   nil,
		}

	}
	close(c)
	return nil,nil
}
