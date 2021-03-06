package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"io"
	"log"
	"time"
)



func CreateCity(city pb.CreateCityRequestP) (*pb.CityResponseP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.CreateCity(ctx, &city)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetCity(city pb.GetCityRequestP) (*pb.CityResponseP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetCity(ctx, &city)

	if err != nil {
		return nil, err
	}

	fmt.Println(r.City)
	fmt.Println(r)
	return r, nil
}

func GetAllCities(request pb.GetAllRequest) (*[]pb.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()

	stream, err := ProfSerConn.context.dbClient.GetAllCitys(ctx, &request)
	if err != nil {
		return nil, err
	}

	var cities []pb.City
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", ProfSerConn.context.dbClient, err)
		}

		cities = append(cities, *c)
	}

	return &cities, nil
}

func UpdateCity(city pb.CreateCityRequestP) (*pb.CityResponseP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.UpdateCity(ctx, &city)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreatePlace(request pb.CreatePlaceRequestP) (*pb.PlaceResponseP, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.CreatePlace(ctx, &request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetAllCityPlaces(request pb.CreateCityRequestP)(*pb.VisitedPlacesResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,err := ProfSerConn.context.dbClient.GetCityPlaces(ctx,&request)
	if err != nil {
		return nil, err
	}
	return res,nil

}

func Search(request pb.SearchAllRequest)(*[]pb.SearchAllResult,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := ProfSerConn.context.dbClient.SearchAllDBA(ctx,&request)

	if err != nil {
		return nil, err
	}
	var sr []pb.SearchAllResult
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			//log.Fatalf("%v.ListFeatures(_) = _, %v", ProfSerConn.context.dbClient, err)
			return nil, err
		}

		sr = append(sr, *c)
	}

	return &sr,nil
}

func GetUserProfile(request pb.GetUserRequestP)(*pb.UserResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,err := ProfSerConn.context.dbClient.GetUser(ctx,&request)
	if err != nil {
		return nil, err
	}
	return res,nil
}

func GetPlace(request pb.GetPlaceRequestP)(*pb.PlaceResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetPlace(ctx,&request)
	if err != nil{
		return nil,err
	}

	return r,nil
}

func VisitCity(request pb.VisitCityRequestP) (*pb.VisitCityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.VisitCity(ctx,&request)
	if err != nil{
		return nil,err
	}
	return r,nil
}


//  rpc GetVisitedCitys (VisitedRequestP) returns (VisitedCitysResponseP);
func GetUserVisitedCities(request pb.VisitedRequestP)(*pb.VisitedCitysResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetVisitedCitys(ctx,&request)
	if err != nil{
		return nil,err
	}
	return r,err
}

//rpc GetVisitedPlaces (VisitedRequestP) returns (VisitedPlacesResponseP);
func GEtVisitedPlaces(request pb.VisitedRequestP)(*pb.VisitedPlacesResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetVisitedPlaces(ctx,&request)
	if err != nil{
		return nil,err
	}
	return r,err
}
