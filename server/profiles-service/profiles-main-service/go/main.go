package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	//"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"

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

//server for client to connect
type server struct {
	pb.UnimplementedProfilesServer
}

//rtest
const(
	name = "user1"
	email = "user1@email.com"
	description = "first user"

	cityName = "galway1"
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

	token ="fafcd86b7538c740e84a5869df1c838853ada4c79c99aaef9373cee4339af01c"
	tokenEmail ="G00341964@gmit.ie"
)

func (s *server) CreateUser(ctx context.Context, in *pb.UserRequestP) (*pb.UserResponseP, error) {

	log.Printf("Received: %v", "create user")
	//check token
	//isToken := CheckToken(in.Email,in.Token)
/*
	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
*/
	bool, err := CreateUser(in.GetEmail(),in.GetName(),in.GetDescription())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &pb.UserResponseP{
		Email:                in.GetEmail(),
		Valid:                bool,
	}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserRequestP) (*pb.UserResponseP, error) {

	log.Printf("Received: %v", "get user")
	//check token
	isToken := CheckToken(in.Email,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := GetUser(in.GetEmail())
	if err != nil {
		panic(err)
	}
	return &res,nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserRequestP) (*pb.UserResponseP, error) {
	log.Printf("Received: %v", "update user")
	//check token
	isToken := CheckToken(in.Email,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res,err := UpdateUser(in.Email,in.Name,in.Description)
	//nothing to do if error, so restart service
	if err != nil {
		panic(err)
	}
	return &pb.UserResponseP{
		Valid:                res,
		Email:                in.Email,
		Name:                 in.Name,
		Description:          in.Description,

	},nil
}

func (s *server) CreateCity(ctx context.Context, in *pb.CityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "create city")
	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := CreateCity(in.Name,in.Country,in.CreatorEmail,in.Location.Lat,in.Location.Lon,in.Description)
	//nothing to do if error, so restart service
	if err != nil {
		panic(err)
	}
	return &res,nil
}

func (s *server) GetCity(ctx context.Context, in *pb.CityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "get city")
	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err :=GetCity(in.Name,in.Country)
	if err != nil {
		return &pb.CityResponseP{
			Valid:                false,


		},nil
	}

	return &pb.CityResponseP{
		Valid:                true,
		Name:                 res.Name,
		Country:             res.Country,
		CreatorEmail:         res.CreatorEmail,
		Description:          res.Description,
		Location:             &pb.GeolocationP{Lat:res.Location.GetLat(),Lon:res.Location.GetLon()},
		Id :                  res.Id,

	},nil
}

func (s *server) CreatePlace(ctx context.Context, in *pb.PlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "create place")
	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := CreatePlace(in.Name, in.City,in.Country,in.Description,in.CreatorEmail,in.Location.GetLat(),in.Location.GetLon())
	if err!= nil{
		panic(err)
	}

	if res <=0 {
		return &pb.PlaceResponseP{
			Valid:                false,},nil
	}
	return &pb.PlaceResponseP{
		Valid:                true,
		Name:                 in.Name,
		City:                 in.City,
		Country:              in.Country,
		CreatorEmail:         in.CreatorEmail,
		Description:          in.Description,
		Location:             &pb.GeolocationP{Lat:in.Location.GetLat(),Lon:in.Location.GetLon()},
		Id:                   res,
	},nil
}

func (s *server) UpdateCity(ctx context.Context, in *pb.CityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "update city")

	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := UpdateCity(in.Name,in.Country,in.CreatorEmail,in.Description,in.Location.Lat,in.Location.Lon)
	if err!= nil{
		panic(err)
	}

	return &pb.CityResponseP{
		Valid:                res,
		Name:                 in.Name,
		Country:              in.Country,
		CreatorEmail:         in.CreatorEmail,
		Description:          in.Description,
		Location:             &pb.GeolocationP{
			Lon:                  in.Location.Lon,
			Lat:                  in.Location.Lat,

		},
	},nil
}

func (s *server) UpdatePlace(ctx context.Context, in *pb.PlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "Update place")
	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := UpdatePlace(in.Name, in.City,in.Country,in.CreatorEmail,in.Description,in.Location.GetLat(),in.Location.GetLon())
	if err!= nil{
		panic(err)
	}

	return &pb.PlaceResponseP{
		Valid:                res,
		Name:                 in.Name,
		City:                 in.City,
		Country:              in.Country,
		CreatorEmail:         in.CreatorEmail,
		Description:          in.Description,
		Location:             &pb.GeolocationP{Lat:in.Location.GetLat(),Lon:in.Location.GetLon()},
	},nil
}



func (s *server) GetPlace(ctx context.Context, in *pb.PlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "Get place")
	//check token
	isToken := CheckToken(in.CreatorEmail,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := GetPlace(in.Name,in.City,in.Country)
	if err!= nil{
		panic(err)
	}

	if res.Id ==0 {
		return &pb.PlaceResponseP{
			Valid:                false,
			Name:                 res.Name,
			City:                 res.City,
			Country:              res.Country,
			CreatorEmail:         res.CreatorEmail,
			Description:          res.Description,
			Location:             &pb.GeolocationP{Lat:res.Location.GetLat(),Lon:res.Location.GetLon()},
			Id:                   res.Id,
		},nil
	}

	return &pb.PlaceResponseP{
		Valid:                true,
		Name:                 res.Name,
		City:                 res.City,
		Country:              res.Country,
		CreatorEmail:         res.CreatorEmail,
		Description:          res.Description,
		Location:             &pb.GeolocationP{Lat:res.Location.GetLat(),Lon:res.Location.GetLon()},
		Id:                   res.Id,
	},nil
}

func (s *server) VisitCity(ctx context.Context, in *pb.VisitCityRequestP) (*pb.VisitCityResponseP, error) {
	log.Printf("Received: %v: %v", "Visir city", in.CityName)
	//check token
	isToken := CheckToken(in.Email,in.Token)
	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
		res,err := VisitCity(in.Email,in.CityName,in.CityCountry)
	if err!= nil{
		panic(err)
	}
		return &pb.VisitCityResponseP{
		Valid:                res.Valid,
		Email:                res.Email,
		CityName:             res.CityName,
		CityCountry:          res.CityCountry,
		},nil
}

func (s *server) VisitPlace(ctx context.Context, in *pb.VisitPlaceRequestP) (*pb.VisitPlaceResponseP, error) {
	log.Printf("Received: %v: %v", "Visir city", in.PlaceName)
	//check token
	isToken := CheckToken(in.Email, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := VisitPlace(in.Email,in.PlaceName,in.PlaceCity,in.PlaceCountry)
	if err!= nil{
		panic(err)
	}
	return &pb.VisitPlaceResponseP{
		Valid:                res.Valid,
		Email:                res.Email,
		PlaceName:            res.PlaceName,
		PlaceCity:            res.PlaceCity,
		PlaceCountry:         res.PlaceCountry,
		},nil
}


func (s *server) GetVisitedCitys(ctx context.Context, in *pb.VisitedRequestP) (*pb.VisitedCitysResponseP, error) {
	log.Printf("Received: %v: %v", "Visited city's", in.Email)
	//check token'
	isToken := CheckToken(in.Email, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := GetVisitedCitys(in.Email)
	if err!= nil{
		panic(err)
	}
	var a []*pb.CityResponseP
	for _,e := range res.Citys{
		a =append(a, &pb.CityResponseP{
			Valid:                true,
			Name:                 e.Name,
			Country:              e.Country,
			CreatorEmail:         e.CreatorEmail,
			Description:          e.Description,
			Location:             &pb.GeolocationP{Lat:e.Location.Lat,Lon:e.Location.Lon},
			Id:                   e.Id,

		})
	}
	return &pb.VisitedCitysResponseP{
		Valid:                true,
		Email:                res.Email,
		Citys:                a,
		},nil
}

func (s *server) GetVisitedPlaces(ctx context.Context, in *pb.VisitedRequestP) (*pb.VisitedPlacesResponseP, error) {
	log.Printf("Received: %v: %v", "Visited place's", in.Email)
	//check token'
	isToken := CheckToken(in.Email, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}
	res, err := GetVisitedPlaces(in.Email)
	if err != nil {
		panic(err)
	}
	var a []*pb.PlaceResponseP
	for _,e := range res.Places{
		a = append(a,&pb.PlaceResponseP{
			Valid:                true,
			Name:                 e.Name,
			City:                 e.City,
			Country:              e.Country,
			CreatorEmail:         e.CreatorEmail,
			Description:          e.Description,
			Location:             &pb.GeolocationP{Lat:e.Location.Lat,Lon:e.Location.Lon},
			Id:                   e.Id,
		})
	}
	return &pb.VisitedPlacesResponseP{
		Valid:                true,
		Email:                in.Email,
		Places:               a,
	},nil
}

func (s *server) GetCityPlaces(ctx context.Context, in *pb.CityRequestP) (*pb.VisitedPlacesResponseP, error) {
	log.Printf("Received: %v: %v", "Visited place's", in.CreatorEmail)
	//check token'
	isToken := CheckToken(in.CreatorEmail, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := GetPlacesCity(in.Name,in.Country)
	if err != nil {
		panic(err)
	}
	var a []*pb.PlaceResponseP
	for _,e := range res.Places{
		a = append(a,&pb.PlaceResponseP{
			Valid:                true,
			Name:                 e.Name,
			City:                 e.City,
			Country:              e.Country,
			CreatorEmail:         e.CreatorEmail,
			Description:          e.Description,
			Location:             &pb.GeolocationP{Lat:e.Location.Lat,Lon:e.Location.Lon},
			Id:                    e.Id,
		})
	}
	return &pb.VisitedPlacesResponseP{
		Valid:                true,
		Email:                in.CreatorEmail,
		Places:               a,
	},nil
}

func main(){
	//conect to neo4j db
	dbserverCtx, err := newNeo4jDBContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &neo4jDB{dbserverCtx}
	dbConn = *s2

	//conect to ps
	// start server pass connection
	psserverCtx, err := newAuthClientContext("35.197.216.42:50051")
	if err != nil {
		log.Fatal(err)
	}
	s1 := &authClient{psserverCtx}
	prsCon = *s1

	//fmt.Println(CheckToken(tokenEmail,token))

	//start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfilesServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}



	////crete user
	//res , errr := CreateUser(email,name,description)
	//if errr != nil {
	//	panic(errr)
	//}
	//fmt.Print(res)
	//
	//get user
	//n,e,d,err := GetUser(email)
	//fmt.Print(n +"\n")
	//fmt.Print(e +"\n")
	//fmt.Print(d +"\n")

	//create city
	//a,err :=CreateCity(cityName,cityCountry,cityUserEmail,cityLat,cityLon,cityDescription)
	//fmt.Print(a.Id)


	////get city
	//r,err := GetCity(cityName,cityCountry);
	//fmt.Print(r.Country + "\n" +r.Name +"\n" +r.GetCreatorEmail() +"\n" +r.Description)
	//fmt.Println(r.Id)
	//
	////create place
	//r1,err := CreatePlace(placeName,placeCity,placeCountry,placeDescription,placeCreatorEmail,placeLat,placeLon)
	//fmt.Println(r1)
	////get place
	//r2,err := GetPlace(placeName,placeCity,placeCountry)
	//fmt.Println(r2)
	//
	//
	//r3, err := VisitPlace(email,placeName,placeCity, placeCountry1)
	//fmt.Println(r3)
	//r4,err := GetVisitedPlaces(email)
	//fmt.Println(r4.GetPlaces()[0].Name)
	////fmt.Println(r4.GetPlaces()[1].Name)
	//for _, value := range r4.GetPlaces(){
	//	fmt.Println(value.Name)
	//}
	//r5,err := VisitCity(email,cityName,cityCountry);
	//fmt.Println(r5)
	//r6,err := GetVisitedCitys(email)
	//for _, value := range r6.GetCitys(){
	//	fmt.Println(value.Name)
	//}
	//b11,err := UpdateUser(email,"new Nax","New descdffadfsa")
	//fmt.Println(b11)
	//
	//b1,err := UpdateCity("galway","ireland","user1@email.com","this is the last" ,7.8,77)
	//fmt.Println(b1)
	//b2,err := UpdatePlace("gmit","galway","ireland","user1@email.com","very very last",4,4)
	//fmt.Println(b2)
	//
	//r7,err := GetPlacesCity("galway","ireland")
	//for _,value := range r7.GetPlaces(){
	//	fmt.Println(value.Name)
	}


	//check token

// password service client
type authClient struct {
	context *authClientContext
}

type authClientContext struct {
	psClient pb.UserAuthenticationClient
	timeout time.Duration
}

//password service connection
var prsCon authClient


func newAuthClientContext(endpoint string) (*authClientContext, error) {
	authConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &authClientContext{
		psClient: pb.NewUserAuthenticationClient(authConn),
		timeout: time.Second,
	}
	return ctx, nil
}


func  CheckToken(email string, token string) bool{
	// Set up a connection to the server.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := prsCon.context.psClient.CheckToken(ctx,&pb.LogRequest{
		Token:                token,
		Email:                email,

	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return r.Sucess
}
