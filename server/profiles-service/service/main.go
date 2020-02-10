package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
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
	//url="localhost:5777";
	//url = "35.234.146.99:5777"
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


//rtest
const(
	name = "user1"
	email = "user1@email.com"
	description = "first user"

	cityName = "galway"
	cityUserEmail = email
	cityLon =45
	cityLat = 54
	cityCountry = "Ireland"
	cityPicture = "pic"
	cityDescription =" a small city"

	placeName = "GMIT"
	placeCity =cityName
	placeCountry = cityCountry
	placeDescription = "Institute of technologies"
	placeLon = 2
	placeLat =1
	placeCreatorEmail = email

	placeName1 = "Square"
	placeCity1 = "galway"
	placeCountry1 = "ireland"
	placeDescription1 = "nice place"
)

func main(){
	dbserverCtx, err := newNeo4jDBContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &neo4jDB{dbserverCtx}
	dbConn = *s2

	//res , errr := CreateUser(email,name,description)
	//if errr != nil {
	//	panic(errr)
	//}
	//fmt.Print(res)
	//n,e,d,err := GetUser(email)
	//fmt.Print(n +"\n")
	//fmt.Print(e +"\n")
	//fmt.Print(d +"\n")
	//a,b,err :=CreateCity(cityName,cityCountry,cityUserEmail,cityLat,cityLon,cityDescription)
	//fmt.Print(a)
	//fmt.Print(b)
	//r,err := GetCity(cityName,cityCountry);
	//fmt.Print(r.Country + "\n" +r.Name +"\n" +r.GetCreatorEmail() +"\n" +r.Description)
	//fmt.Println(r.GetLocation())
	//r1,err := CreatePlace(placeName1,placeCity1,placeCountry1,placeDescription1,placeCreatorEmail,placeLat,placeLon)
	//fmt.Println(r1)
	//r2,err := GetPlace(placeName,placeCity,placeCountry)
	//fmt.Println(r2)
	//r3, err := VisitPlace(email,placeName1,placeCity1, placeCountry1)
	//fmt.Println(r3)
	//r4,err := GetVisitedPlaces(email)
	//fmt.Println(r4.GetPlaces()[0].Name)
	//fmt.Println(r4.GetPlaces()[1].Name)
	//for _, value := range r4.GetPlaces(){
	//	fmt.Println(value.Name)
	//}
	//r5,err := VisitCity(email,cityName,cityCountry);
	//fmt.Println(r5)
//	r6,err := GetVisitedCitys(email)
//	for _, value := range r6.GetCitys(){
	//	fmt.Println(value.Name)
	//}
	//b,err := UpdateUser(email,"new Nax","New descdffadfsa")
	//fmt.Println(b)

	//b1,err := UpdateCity("galway","ireland","user1@email.com","this is the last" ,7.8,77)
	//fmt.Println(b1)
	//b2,err := UpdatePlace("gmit","galway","ireland","user1@email.com","very very ",4,4)
	//fmt.Println(b2)

	r7,err := GetPlacesCity("galway","ireland")
	for _,value := range r7.GetPlaces(){
		fmt.Println(value.Name)
	}

}

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
		panic(err)
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
		panic(err)
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

func CreatePlace(name string, city string, country string, description string, creatorEmail string, lat float32, lon float32)(string,error){
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
	return r.GetName() , nil
}

func GetCity(name string, country string)(pb.CityPDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.GetCity(ctx,&pb.CityRequestPDB{
		Name:strings.ToLower(name),
		Country:strings.ToLower(country)})
	if err != nil {
		panic(err)
	}
	return *r,nil
}

func CreateCity(name string, country string, creatorEmail string, lat float32,lon float32, description string) (string,string,error){
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
	return r.Name,r.Country,nil
}

func CreateUser(email string,name string, description string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.CreateUser(ctx, &pb.CreateUserRequestPDB{
		Email: strings.ToLower(email),
		Name: strings.ToLower(name),
		Description:description})
	if err != nil {
		return false, err
	}
	res, err := strconv.ParseBool(r.Valied)

	return res,nil
}

func GetUser(email string)(string,string,string,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := dbConn.context.dbClient.GetUser(ctx, &pb.GetUserRequestPDB{Email: strings.ToLower(email)})
	if err != nil {
		panic(err)
	}
	return r.Email, r.Name,r.Description,nil
}

