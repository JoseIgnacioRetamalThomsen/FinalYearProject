package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"strings"

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
	//url = "35.197.221.57:5777"
	//url = "34.89.92.151:5777"
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

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequestP) (*pb.UserResponseP, error) {

	log.Printf("Received: %v", "create user")
/*	//check token
	//isToken := CheckToken(in.Email,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
*/
	res, err := CreateUser(*in.User)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return &pb.UserResponseP{
		Valid:                res.Valid,
		User:                 res.User,
		XXX_NoUnkeyedLiteral: struct{}{},
	}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequestP) (*pb.UserResponseP, error) {

	log.Printf("Received: %v", "get user")
	//check token
	isToken := CheckToken(in.Email,in.Token)
	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
	res, err := GetUser(pb.GetUserRequestPDB{
		Email:                strings.ToLower(in.Email),
		})
	if err != nil {
		return nil,err
	}
	return &pb.UserResponseP{
		Valid:                res.Valid,
		User:                 res.User,

	},nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.CreateUserRequestP) (*pb.UserResponseP, error) {
	log.Printf("Received: %v", "update user")
	//check token
	isToken := CheckToken(in.Email,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res,err := UpdateUser(*in.User)
	//nothing to do if error, so restart service
	if err != nil {
		return nil,err
	}
	return &pb.UserResponseP{
		Valid:                true,
		User:                 res.User,

	},nil
}

func (s *server) CreateCity(ctx context.Context, in *pb.CreateCityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "create city")
	//check token
	isToken := CheckToken(in.Name,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := CreateCity(*in.City)
	if err != nil {
		return nil,err
	}
	return &pb.CityResponseP{
		Valid:                res.Valid,
		City:                 res.City,
	},nil
}

func (s *server) GetCity(ctx context.Context, in *pb.GetCityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "get city")
	//check token
	isToken := CheckToken(in.Name,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err :=GetCity(pb.CityRequestPDB{
		Name:                 strings.ToLower(in.CityName),
		Country:              strings.ToLower(in.CityCountry),

	})
	if err != nil {
		return nil,err
	}

	return &pb.CityResponseP{
		Valid:                res.Valid,
		City:                 res.City,
	},nil
}

func (s *server) CreatePlace(ctx context.Context, in *pb.CreatePlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "create place")
	//check token
	isToken := CheckToken(in.Name,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := CreatePlace(*in.Place)

	if err != nil{
		return nil,err
	}
	return &pb.PlaceResponseP{
		Valid:                res.Valid,
		Place:                res.Place,
	},nil
}

func (s *server) UpdateCity(ctx context.Context, in *pb.CreateCityRequestP) (*pb.CityResponseP, error) {
	log.Printf("Received: %v", "update city")

	//check token
	isToken := CheckToken(in.Name,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := UpdateCity(*in.City)
	if err!= nil{
		panic(err)
	}

	return &pb.CityResponseP{
		Valid:                res,
		City:                 in.City,
	},nil
}

func (s *server) UpdatePlace(ctx context.Context, in *pb.CreatePlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "Update place")
	//check token
	isToken := CheckToken(in.Name,in.Token)
	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
	res, err := UpdatePlace(*in.Place)
	if err!= nil{
		return nil,err
	}
	return &pb.PlaceResponseP{
		Valid:                res,
		Place:                in.Place,
	},nil
}



func (s *server) GetPlace(ctx context.Context, in *pb.GetPlaceRequestP) (*pb.PlaceResponseP, error) {
	log.Printf("Received: %v", "Get place")
	//check token
	isToken := CheckToken(in.Email,in.Token)

	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}

	res, err := GetPlace(pb.PlaceRequestPDB{
		Name:                 strings.ToLower(in.PlaceName),
		City:                  strings.ToLower(in.PlaceCity),
		Country:               strings.ToLower(in.PlaceCountry),

	})
	if err!= nil{
		return nil, err
	}
	return &pb.PlaceResponseP{
		Valid:                true,
		Place:                res,},nil
}

func (s *server) VisitCity(ctx context.Context, in *pb.VisitCityRequestP) (*pb.VisitCityResponseP, error) {
	log.Printf("Received: %v: %v", "Visir city", in)
	//check token
	isToken := CheckToken(in.Email,in.Token)
	if isToken == false{
		return nil,status.Error(codes.PermissionDenied,"Invalid token")
	}
		res,err := VisitCity(pb.VisitCityRequestPDB{
			UserEmail:            strings.ToLower(in.Email),
			CityId:               in.Id,

		})
	if err!= nil{
		return nil,err
	}
		return &pb.VisitCityResponseP{
			Valid:                res.Valid,
			TimeStamp:            res.TimeStamp,

		},nil
}

func (s *server) VisitPlace(ctx context.Context, in *pb.VisitPlaceRequestP) (*pb.VisitPlaceResponseP, error) {
	log.Printf("Received: %v: %v", "Visir city", in)
	//check token
	isToken := CheckToken(in.Email, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := VisitPlace(pb.VisitPlaceRequestPDB{
		UserEmail:            strings.ToLower(in.Email),
		PlaceId:              in.PlaceId,

	})
	if err!= nil{
		return nil,err
	}
	return &pb.VisitPlaceResponseP{
		Valid:                res.Valid,
		TimeStamp:            res.TimeStamp,
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
		return nil,err
	}
	return &pb.VisitedCitysResponseP{
		Valid:                true,
		Citys:                res.Citys,
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
		return nil,err
	}

	return &pb.VisitedPlacesResponseP{
		Valid:                true,
		Places:               res.Places,
	},nil
}

func (s *server) GetCityPlaces(ctx context.Context, in *pb.CreateCityRequestP) (*pb.VisitedPlacesResponseP, error) {
	log.Printf("Received: %v: %v", "Git city  place's", in)
	//check token'
	isToken := CheckToken(in.Name, in.Token)
	if isToken == false {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := GetPlacesCity(pb.CityRequestPDB{
		Name:                 in.City.Name,
		Country:              in.City.Country,
	})
	if err != nil {
		return nil,err
	}

	return &pb.VisitedPlacesResponseP{
		Valid:                true,
		Places:               res.Places,
	},nil
}



func (s *server) GetAllCitys(in *pb.GetAllRequest, stream pb.Profiles_GetAllCitysServer) error {
	// create a channel, that is like a blocking queue
	ch := make(chan CityResult)


	go GetAllCitysDBA(in,ch)
	for i := range ch {
		if i.Err != nil{
			return i.Err
		}
		if err := stream.Send(i.City); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetAllPlaces(in *pb.GetAllRequest, stream pb.Profiles_GetAllPlacesServer) error {
	// create a channel, that is like a blocking queue
	ch := make(chan PlaceResult)

	// Get all places in different thread
	go GetAllPlacesDBA(in,ch)

	// send results back async
	for i := range ch {
		//send the error if
		if i.Err != nil{
			return i.Err
		}
		// get next place from stream
		if err := stream.Send(i.Place); err != nil {
			return err
		}
	}
	// all finish with no errors
	return nil
}

func main() {
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


	//get city
//	r, _ := GetCity("galway", "ireland");
	//fmt.Print(r.Country + "\n" + r.Name + "\n" + r.GetCreatorEmail() + "\n" + r.Description)
//	fmt.Println("\nid :", r.Id)
//	fmt.Println(CheckToken(tokenEmail,token))

/*
	for i := 0; i < 1; i++ {
		//crete user
		res, errr := CreateUser(email, name, description)
		if errr != nil {
			panic(errr)
		}
		fmt.Print(res)

		//create city
		a, _ := CreateCity(cityName, cityCountry, cityUserEmail, cityLat, cityLon, cityDescription)
		fmt.Print(a.Id)

		//get city
		r, _ := GetCity(cityName, cityCountry);
		fmt.Print(r.Country + "\n" + r.Name + "\n" + r.GetCreatorEmail() + "\n" + r.Description)
		fmt.Println(r.Id)

		//create place
		r1, _ := CreatePlace(placeName, placeCity, placeCountry, placeDescription, placeCreatorEmail, placeLat, placeLon)
		fmt.Println(r1)
		//get place
		r2, _ := GetPlace(placeName, placeCity, placeCountry)
		fmt.Println(r2)

		r3, _ := VisitPlace(email, placeName, placeCity, placeCountry1)
		fmt.Println(r3)
		r4, _ := GetVisitedPlaces(email)
		fmt.Println(r4.GetPlaces()[0].Name)
		//fmt.Println(r4.GetPlaces()[1].Name)
		for _, value := range r4.GetPlaces() {
			fmt.Println(value.Name)
		}
		r5, _ := VisitCity(email, cityName, cityCountry);
		fmt.Println(r5)
		r6, _ := GetVisitedCitys(email)
		for _, value := range r6.GetCitys() {
			fmt.Println(value.Name)
		}
	}*/
//	b11, err := UpdateUser(email, "new Nax", "New descdffadfsa")
//	fmt.Println(b11)

	//b1, err := UpdateCity("galway", "ireland", "user1@email.com", "this is the last", 7.8, 77)
	//fmt.Println(b1)
	//b2, err := UpdatePlace("gmit", "galway", "ireland", "user1@email.com", "very very last", 4, 4)
	//fmt.Println(b2)
/*
	r7, err := GetPlacesCity("galway", "ireland")
	for _, value := range r7.GetPlaces() {
		fmt.Println(value.Name)
	}*/

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
