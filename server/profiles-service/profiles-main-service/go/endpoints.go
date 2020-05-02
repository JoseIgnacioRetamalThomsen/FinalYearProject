package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)


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

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequestP) (*pb.UserResponseP, error) {

	log.Printf("Received: %v", "create user")

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

