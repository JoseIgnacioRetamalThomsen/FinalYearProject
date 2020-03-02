package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joseignacioretamalthomsen/photos-dba/dba"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct {
 pb.UnimplementedPhotosDBAServiceServer
}
type Configuration struct {
	Port string
	Coneection_type    string
	MySQL_socket string
	MySQL_user string
	MySQL_pass string
	MySQL_db string
}
var configuration Configuration


func readConfig(fileName string){
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (s *server) AddProfilePhotoDBA(ctx context.Context, in *pb.AddProfilePhotoDBARequest) (*pb.AddProfilePhotoDBAResponse, error) {
	log.Printf("Received: %v", "Add profile photo")
	id ,time, err:= dba.AddProfilePhoto(pb.ProfilePhoto{
		UserEmail:            in.UserEmail,
		Url:                  in.Url,
		Selected:             in.Selected,
	})

	if err != nil{
		log.Printf("Error: %v", err)
		return  &pb.AddProfilePhotoDBAResponse{
			Success:              false,
		},err
	}
	return &pb.AddProfilePhotoDBAResponse{
		Success:              true,
		Id:                   int32(id),
		TimeStamp:            time,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil
}

func (s *server) GetProfilePhotoDBA(ctx context.Context, in *pb.GetProfilePhotosDBARequest) (*pb.GetProfilePhotosDBAResponse, error) {
	log.Printf("Received: %v", "Get profile photo")

	photos,err := dba.GetProfilePhotos(in.UserEmail)

	if err!= nil{
		return &pb.GetProfilePhotosDBAResponse{
			Sucess:               false,


		},err
	}

	return &pb.GetProfilePhotosDBAResponse{
		Sucess:               true,
		Photos:               photos,

	},nil
}

func (s *server) AddCityPhotoDBA(ctx context.Context, in *pb.AddCityPhotoDBARequest) (*pb.AddCityPhotoDBAResponse, error) {
	log.Printf("Received: %v", "Add city photo")
	id ,time, err:= dba.AddCityPhoto(pb.CityPhoto{

		Url:                  in.Url,
		CityId:   in.CityId,
		Selected:             false,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	if err != nil{
		log.Printf("Error: %v", err)
		return  &pb.AddCityPhotoDBAResponse{
			Success:              false,

		},err
	}
	return &pb.AddCityPhotoDBAResponse{
		Success:              true,
		Id:                   int32(id),
		TimeStamp:            time,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil
}

func (s *server) GetCityPhotoDBA(ctx context.Context, in *pb.GetCityPhotosDBARequest) (*pb.GetCityPhotosDBAResponse, error) {
	log.Printf("Received: %v", "Get city photo")

	photos,err := dba.GetCityPhotos(int(in.CityId))

	if err!= nil{
		return &pb.GetCityPhotosDBAResponse{
			Sucess:               false,


		},err
	}

	return &pb.GetCityPhotosDBAResponse{
		Sucess:               true,
		Photos:               photos,

	},nil
}

func (s *server) AddPlacePhotoDBA(ctx context.Context, in *pb.AddPlacePhotoDBARequest) (*pb.AddPlacePhotoDBAResponse, error) {
	log.Printf("Received: %v", "Add place photo")
	id ,time, err:= dba.AddPlacePhoto(pb.PlacePhoto{

		Url:                  in.Url,
		PlaceId:   in.PlaceId,
		Selected:             false,

	})

	if err != nil{
		log.Printf("Error: %v", err)
		return  &pb.AddPlacePhotoDBAResponse{
			Success:              false,

		},err
	}
	return &pb.AddPlacePhotoDBAResponse{
		Success:              true,
		Sd:                   int32(id),
		TimeStamp:            time,

	},nil
}

func (s *server) GetPlacePhotoDBA(ctx context.Context, in *pb.GetPlacePhotosDBARequest) (*pb.GetPlacePhotosDBAResponse, error) {
	log.Printf("Received: %v", "Get city photo")

	photos,err := dba.GetPlacePhotos(int(in.PlaceId))

	if err!= nil{
		return &pb.GetPlacePhotosDBAResponse{
			Sucess:               false,


		},err
	}

	return &pb.GetPlacePhotosDBAResponse{
		Sucess:               true,
		Photos:               photos,

	},nil
}


func (s *server) AddPostPhotoDBA(ctx context.Context, in *pb.AddPostPhotoDBARequest) (*pb.AddPostPhotoDBAResponse, error) {
	log.Printf("Received: %v", "Add post photo")
	id ,time, err:= dba.AddPostPhoto(pb.PostPhoto{
		PostId:              in.PostId,
		Url:                  in.Url,
		Selected:             false,

	})

	if err != nil{
		log.Printf("Error: %v", err)
		return  &pb.AddPostPhotoDBAResponse{
			Sucess:              false,

		},err
	}
	return &pb.AddPostPhotoDBAResponse{
		Sucess:               true,
		Id:                   int32(id),
		TimeStamp:            time,

	},nil
}

func (s *server) GetPostPhotoDBA(ctx context.Context, in *pb.GetPostPhotosDBARequest) (*pb.GetPostPhotosDBAResponse, error) {
	log.Printf("Received: %v", "Get place photo")

	photos,err := dba.GetPostPhotos(in.PlaceId)

	if err!= nil{
		return &pb.GetPostPhotosDBAResponse{
			Sucess:               false,


		},err
	}

	return &pb.GetPostPhotosDBAResponse{
		Sucess:               true,
		Photos:               photos,

	},nil
}

func main(){

	//read config file name from console input
	args := os.Args[1]
	readConfig(args)

	dba.SetupConnection(configuration.Coneection_type,
		configuration.MySQL_socket,
		configuration.MySQL_user,
		configuration.MySQL_pass,
		configuration.MySQL_db)
	log.Print("Starting Service")

//test
/*
	dba.AddProfilePhoto(pb.ProfilePhoto{

		UserEmail: "email1",
		Url:       "url",

		Selected: false,
	})
/*
	dba.AddCityPhoto(pb.CityPhoto{

		CityId:               6,
		Url:                  "uel",
		Selected:             false,

	})

	idp,_ := dba.AddPlacePhoto(pb.PlacePhoto{

		PlaceId:              5,
		Url:                  "fgsdfg",

		Selected:             false,

	})
	fmt.Println(idp)
	idpo,_ := dba.AddPostPhoto(pb.PostPhoto{

		PostId:               "hj6",
		Url:                  "gfhdgf",

		Selected:             false,

	})
	fmt.Println(idpo)/*
	_,err := dba.GetProfilePhotos("email1")
	if err!= nil{
		panic(err)
	}*//*
   _,err:= dba.GetPostPhotos("hj6")
	if err!= nil{
		panic(err)
	}

 */
	//end test
	lis, err := net.Listen("tcp", configuration.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPhotosDBAServiceServer(s, &server{})
	log.Print("Service Started in port: ", configuration.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
