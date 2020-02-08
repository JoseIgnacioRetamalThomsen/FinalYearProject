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
)

func main(){
	dbserverCtx, err := newNeo4jDBContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &neo4jDB{dbserverCtx}
	dbConn = *s2

	res , errr := CreateUser(email,name,description)
	if errr != nil {
		panic(errr)
	}
	fmt.Print(res)
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
	//r,err := CreatePlace(placeName,placeCity,placeCountry,placeDescription,placeCreatorEmail,placeLat,placeLon)
	//fmt.Println(r)
	//r,err := GetPlace(placeName,placeCity,placeCountry)
	//fmt.Println(r)
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
	r, err := dbConn.context.dbClient.GetCity(ctx,&pb.CityRequestPDB{Name:name,Country:country})
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
	r, err := dbConn.context.dbClient.CreateUser(ctx, &pb.CreateUserRequestPDB{Email: strings.ToLower(email),
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

