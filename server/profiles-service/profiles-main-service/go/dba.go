package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
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



