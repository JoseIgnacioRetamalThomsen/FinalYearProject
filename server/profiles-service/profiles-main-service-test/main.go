package main

import (
	"context"
	"fmt"

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
	token ="119b8f15250c9f0f5dad71375e8dd8b06aa42a13500b62e4ec780e50c112bcf4"
	tokenEmail ="G00341964@gmit.ie"
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

	//fmt.Println(CreateUser("email7","namef","description4",token))
	fmt.Println(GetUser(tokenEmail,token))
	//fmt.Println(UpdateUser(tokenEmail,"pepe","student",token))
	//fmt.Println(CreateCity(tokenEmail,token,"San Pedro","Chile","Bacn",12,12))
	//fmt.Println(GetCity(tokenEmail,token,"San Pedro" , "Chile"))
	//fmt.Println(CreatePlace(tokenEmail,token,"plaza","san pedro","chile","nada",3,3))
	//fmt.Println(UpdateCity(tokenEmail,token,"San Pedro","Chile","Bafome",12,12))
	//fmt.Println(UpdatePlace(tokenEmail,token,"plaza","san pedro","chile","Algo",3,3))
	//fmt.Println(GetPlace(tokenEmail,token, "plaza","San Pedro", "chile"))
//	fmt.Println(VisitCity(tokenEmail,token,"San Pedro","Chile"))
//fmt.Println(VisitPlace(tokenEmail,token,"plaza","san pedro","chile"))
//fmt.Println(GetVisitedCity(tokenEmail,token))
//fmt.Println(GetVisitedPlaces(tokenEmail,token))
//fmt.Println(GetCityPlaces(tokenEmail,token,"san Pedro","chile"))
}


func CreateUser(email string,name string, description string,token string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateUser(ctx,&pb.UserRequestP{
		Token:                token,
		Email:                email,
		Name:                 name,
		Description:          description,

	})
	if err != nil{
		panic(err)
	}

	return r.Valid,nil
}

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

func CreateCity(email string,token string,cityName string,cityCountry string,cityDescription string,lat float32,lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateCity(ctx,&pb.CityRequestP{
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

func CreatePlace(email string, token string, name string, city string,country string,description string, lat float32, lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreatePlace(ctx,&pb.PlaceRequestP{
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
	fmt.Println(r)
	return r.Valid
}

func VisitCity(email string, token string, name string,country string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.VisitCity(ctx,&pb.VisitCityRequestP{
		Token:       token,
		Email:       email,
		CityName:    name,
		CityCountry: country,
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}

func VisitPlace(email string, token string, name string,city string,country string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.VisitPlace(ctx,&pb.VisitPlaceRequestP{
		Token:                token,
		Email:                email,
		PlaceName:            name,
		PlaceCity:            city,
		PlaceCountry:         country,

	})
	if err != nil{
		panic(err)
	}
	fmt.Println(r)
	return r.Valid
}


func GetVisitedCity(email string, token string)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
	fmt.Println(r)
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
