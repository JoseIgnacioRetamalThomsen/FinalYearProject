package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"strconv"
	"strings"
	"time"
)

func GetPlacesCity(cityName string, cityCountry string)(pb.VisitedPlacesResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetCityPlaces(ctx,&pb.CityRequestPDB{
		Name:                 strings.ToLower(cityName),
		Country:              strings.ToLower(cityCountry)})
	if err != nil {
		panic(err)
	}
	return *r,nil
}

func UpdatePlace(name string, city string, country string,creatorEmail string, description string, lat float32, lon float32)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdatePlaceRequest(ctx,&pb.PlacePDB{
		Name:                 strings.ToLower(name),
		City:                 strings.ToLower(city),
		Country:              strings.ToLower(country),
		CreatorEmail:         strings.ToLower(creatorEmail),
		Location:             &pb.GeolocationPDB{
			Lon:                  lat,
			Lat:                  lon},
		Description:          description})
	if err != nil {
		panic(err)
	}
	return r.Valid,nil
}

func UpdateCity(name string, country string, creatorEmail string, description string, lat float32 , lon float32)(bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateCityRequest(ctx,&pb.CityPDB{
		Name:    strings.ToLower(name),
		Country:              strings.ToLower(country),
		CreatorEmail:         strings.ToLower(creatorEmail),
		Location:             &pb.GeolocationPDB{
			Lon:                  lat,
			Lat:                  lon,

		},
		Description:          strings.ToLower(description),
		Places:               nil,

	})
	if err != nil {
		return false,nil
	}
	return r.Valid,nil

}

func UpdateUser(email string,name string, description string)(bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateUserRequest(ctx, &pb.CreateUserRequestPDB{
		Email: strings.ToLower(email),
		Name: strings.ToLower(name),
		Description:description})
	if err != nil {
		return false,nil
	}
	res, err := strconv.ParseBool(r.Valied)

	return res,nil
}
func GetVisitedCitys(email string)(pb.VisitedCitysResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetVisitedCitys(ctx,&pb.VisitedCitysRequestPDB{
		Email:           strings.ToLower(email)})
	if err !=nil{
		panic(err)
	}
	return *r, nil
}

func VisitCity(email string, cityName string, cityCountry string)(pb.VisitCityResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.VisitCity(ctx,&pb.VisitCityRequestPDB{
		Email:                strings.ToLower(email),
		CityName:             strings.ToLower(cityName),
		CityCountry:          strings.ToLower(cityCountry)})
	if err !=nil{
		panic(err)
	}

	return *r,nil

}

func GetVisitedPlaces(email string)(pb.VisitedPlacesResponsePDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetVisitedPlaces(ctx,&pb.VisitedPlacesRequestPDB{
		Email:  strings.ToLower(email)})
	if err != nil {
		panic(err)
	}

	return *r,nil
}

func VisitPlace(email string, placeName string, placeCity string, placeCountry string)(pb.VisitPlaceResponsePDB,error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.VisitPlace(ctx,&pb.VisitPlaceRequestPDB{
		Email:strings.ToLower(email),
		PlaceName: strings.ToLower(placeName),
		PlaceCity:strings.ToLower(placeCity),
		PlaceCountry:strings.ToLower(placeCountry)})
	if err != nil {
		panic(err)
	}

	return *r,nil
}

func GetPlace(name string, city string,country string)(pb.PlacePDB,error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetPlace(ctx,&pb.PlaceRequestPDB{
		Name:strings.ToLower(name),
		City:strings.ToLower(city),
		Country:strings.ToLower(country)})
	if err != nil {
		panic(err)
	}

	return *r,nil
}

func CreatePlace(name string, city string, country string, description string, creatorEmail string, lat float32, lon float32)(bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreatePlaceRequest(ctx,&pb.PlacePDB{
		Name:strings.ToLower(name),
		City:strings.ToLower(city),
		Country:strings.ToLower(country),
		Description:description,
		CreatorEmail:strings.ToLower(creatorEmail),
		Location: &pb.GeolocationPDB{Lat:lat,Lon:lon}})
	if err != nil {
		panic(err)
	}
	return r.Valid , nil
}

func GetCity(name string, country string)(pb.CityPDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetCity(ctx,&pb.CityRequestPDB{
		Name:strings.ToLower(name),
		Country:strings.ToLower(country)})
	if err != nil {
		return pb.CityPDB{}, err
	}
	return *r,nil
}

func CreateCity(name string, country string, creatorEmail string, lat float32,lon float32, description string) (pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateCity(ctx, &pb.CityPDB{ Name: strings.ToLower(name),
		Country:strings.ToLower(country),
		CreatorEmail:strings.ToLower(creatorEmail),
		Description:description,
		Location:&pb.GeolocationPDB{Lat:lat,Lon:lon}})
	if err != nil {
		panic(err)
	}
	return pb.CityResponseP{
		Valid:                r.Valid,
		Name:                 r.Name,
		Country:              r.Country,
		CreatorEmail:         creatorEmail,
		Description:          description,
		Location:             &pb.GeolocationP{Lat:lat,Lon:lon},

	},nil
}

func CreateUser(email string,name string, description string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateUser(ctx, &pb.CreateUserRequestPDB{
		Email: strings.ToLower(email),
		Name: strings.ToLower(name),
		Description:description})
	if err == nil {
		return false, err
	}
	res, err := strconv.ParseBool(r.Valied)

	return res,nil
}

func GetUser(email string)(pb.UserResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := dbConn.context.dbClient.GetUser(ctx, &pb.GetUserRequestPDB{Email: strings.ToLower(email)})
	if err != nil {
		return pb.UserResponseP{Valid:false},nil
	}
	return pb.UserResponseP{
		Valid:                true,
		Email:                r.Email,
		Name:                 r.Name,
		Description:          r.Description,
		},nil
	//r.Email, r.Name,r.Description,nil
}



