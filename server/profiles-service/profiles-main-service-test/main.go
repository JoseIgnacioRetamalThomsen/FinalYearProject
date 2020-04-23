package main

import (
	"context"
	"fmt"
	"io"

	//"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	//"strconv"
	//"strings"

	//	"net"
	"time"

	//"errors"
	//"fmt"
	//pb "github.com/joseignacioretamalthomsen/wcity"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/resolver"
	//"time"
)

const (
	port = ":60051"

)
const(
	//url = "0.0.0.0:60051"
	url="35.197.216.42:60051";
	//url = "35.234.146.99:5777"
	token ="fafcad212731804df628c349717a30855872de34449093dbc7616d2411af1ec7"
	tokenEmail ="g00341964@gmit.ie"
)

type profileServer struct {
	context *profileServiceContext
}

type profileServiceContext struct {
	dbClient pb.ProfilesClient
	timeout time.Duration
}
var profSerConn profileServer

// create connection
func newProfilesServiceContext(endpoint string) (*profileServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &profileServiceContext{
		dbClient: pb.NewProfilesClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}


func main(){
	//conect to server
	dbserverCtx, err := newProfilesServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &profileServer{dbserverCtx}
	profSerConn = *s2
    // GetCity(tokenEmail,token,"Galway","Ireland")
	//GetAllCitys()
	//GetAllPlaces()
	fmt.Println(CreateUser(tokenEmail,"namef","description4",token))
//	fmt.Println(GetUser(tokenEmail,token))
//	fmt.Println(UpdateUser(tokenEmail,"pepe","student",token))
	//fmt.Println(CreateCity(tokenEmail,token,"galway","ireland","Bacn",12,12))
	//fmt.Println(GetCity(tokenEmail,token,"galway", "ireland"))
	//fmt.Println(CreatePlace(tokenEmail,token,"plaza1","san pedro","chile","nada",3,3))
//fmt.Println(UpdateCity(tokenEmail,token,"San Pedro","Chile","Bafome",12,12))
	//fmt.Println(UpdatePlace(tokenEmail,token,"plaza","san pedro","chile","Algo",3,3))

	//fmt.Println(GetPlace(tokenEmail,token, "gmit","galway", "ireland"))
	//
	//fmt.Println(VisitCity(tokenEmail,token,26))
//fmt.Println(VisitPlace(tokenEmail,token,0))
//fmt.Println(GetVisitedCity(tokenEmail,token))
//fmt.Println(GetVisitedPlaces(tokenEmail,token))
//fmt.Println(GetCityPlaces(tokenEmail,token,"san Pedro","chile"))
}


func CreateUser(email string,name string, description string,token string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateUser(ctx,&pb.CreateUserRequestP{
		Token:                token,
		Email:                email,
		User:                 &pb.User{
			Email:                email,
			Name:                 "user2",
			Descripiton:          "a user",

		},

	})
	if err != nil{
		panic(err)
	}

	return r.Valid,nil
}
/*
func GetUser(email string,token string)(pb.UserResponseP,error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetUser(ctx,&pb.UserRequestP{Email: email,Token:token})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Name)
	fmt.Println(r.GetDescription())
	return *r,nil
}


func UpdateUser(email string,name string, description string, token string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.UpdateUser(ctx, &pb.UserRequestP{
		Token:                token,
		Email:                email,
		Name:                 name,
		Description:          description,

	})
	if err != nil{
		panic(err)
	}
	return r.Valid
}



func GetCity(email string,token string,cityName string, cityCounty string )bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetCity(ctx,&pb.CityRequestP{
		Token:                token,
		Name:                 cityName,
		Country:              cityCounty,
		CreatorEmail:         email,
		})

	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}



func UpdateCity(email string,token string,cityName string,cityCountry string,cityDescription string,lat float32,lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.UpdateCity(ctx,&pb.CityRequestP{
		Token:                token,
		Name:                 cityName,
		Country:              cityCountry,
		CreatorEmail:         email,
		Description:          cityDescription,
		Location:             &pb.GeolocationP{Lat:lat,Lon:lon},

	})
	if err != nil{
		panic(err)
	}
	return r.Valid
}

func UpdatePlace(email string, token string, name string, city string,country string,description string, lat float32, lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.UpdatePlace(ctx,&pb.PlaceRequestP{
		Token:                token,
		Name:                 name,
		City:                 city,
		Country:              country,
		CreatorEmail:         email,
		Description:          description,
		Location:             &pb.GeolocationP{Lat:lat, Lon:lon},

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}

func GetPlace(email string, token string, name string, city string,country string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetPlace(ctx,&pb.PlaceRequestP{
		Token:                token,
		Name:                name,
		City:                 city,
		Country:              country,
		CreatorEmail:         email,

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Id)
	fmt.Println(r.Name)
	fmt.Println(r.Description)
	fmt.Println(r.Name)
	fmt.Println(r.Id)
	return r.Valid
}






func GetCityPlaces(email string, token string, cityName string, cityCountry string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetCityPlaces(ctx,&pb.CityRequestP{
		Token:                token,
		Name:                 cityName,
		Country:              cityCountry,
		})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}

*/
func GetCity(email string,token string,cityName string, cityCounty string )bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetCity(ctx,&pb.GetCityRequestP{
		Token:                token,
		Name:                email,
		CityName:             cityName,
		CityCountry:          cityCounty,

	})

	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}


func GetAllCitys(){

	// initialize a pb.Rectangle
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := profSerConn.context.dbClient.GetAllCitys(ctx,&pb.GetAllRequest{
		Max:                  100,

	})
	//stream, err := client.ListFeatures(context.Background(), rect)
	if err != nil {
		panic (err)
	}
	for {
		city, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", profSerConn.context.dbClient, err)
		}
		log.Println(city)
	}
}
func GetAllPlaces(){

	// initialize a pb.Rectangle
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := profSerConn.context.dbClient.GetAllPlaces(ctx,&pb.GetAllRequest{
		Max:                  100,

	})
	//stream, err := client.ListFeatures(context.Background(), rect)
	if err != nil {
		panic (err)
	}
	for {
		city, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", profSerConn.context.dbClient, err)
		}
		log.Println(city)
	}
}

// need to finish
/*
func CreatePlace(email string, token string, name string, city string,country string,description string, lat float32, lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreatePlace(ctx,&pb.CreatePlaceRequestP{
		Token:                "",
		Name:                 "",
		Place:                nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}
*/

func CreateCity(email string,token string,cityName string,cityCountry string,cityDescription string,lat float32,lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateCity(ctx,&pb.CreateCityRequestP{
		Token:                token,
		Name:                 email,
		City:                 &pb.City{
			Name:                 cityName,
			Country:              cityCountry,
			CreatorEmail:         email,
			Location:            &pb.Geolocation{
				Lon:                  lat,
				Lat:                  lon,

			},
			Description:          cityDescription,

		},

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}


func GetVisitedCity(email string, token string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetVisitedCitys(ctx,&pb.VisitedRequestP{
		Token:                token,
		Email:                email,

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}

func VisitCity(email string, token string, id int32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.VisitCity(ctx,&pb.VisitCityRequestP{
		Token:       token,
		Email:       email,
		Id : id,
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}



func GetVisitedPlaces(email string, token string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetVisitedPlaces(ctx,&pb.VisitedRequestP{
		Token:                token,
		Email:                email,

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r.Places)
	return r.Valid
}

func VisitPlace(email string, token string, id int32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := profSerConn.context.dbClient.VisitPlace(ctx,&pb.VisitPlaceRequestP{
		Token:                token,
		Email:                email,
		PlaceId:  id,

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}
