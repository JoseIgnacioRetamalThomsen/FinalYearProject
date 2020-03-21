package main
//  "MySQL_socket" : "10.154.0.6:3306",
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
		PlaceCityId: in.PlaceCityId,

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
	id ,time, err:= dba.AddPostPhoto(*in)

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

func (s *server) GetCitysPhotosDBA(ctx context.Context, in *pb.GetCitysPhotoRequest) (*pb.GetCitysPhotoResponse, error) {
	log.Printf("Received: %v", "Get all citys photos")

	photos,err := dba.GetCitysPhotos()

	if err!= nil{
		return &pb.GetCitysPhotoResponse{
			Success:              false,
			CityPhotos:           nil,

		},err
	}

	return &pb.GetCitysPhotoResponse{
		Success:              false,
		CityPhotos:           photos,

	},nil
}

func (s *server) GetPostsPhotosIdDBA(ctx context.Context, in *pb.GetPostsPhotosPerParentRequest) (*pb.GetPostsPhotosPerParentResponse, error) {
	log.Printf("Received: %v", "Get all post photos per parent called")

	photos ,err  := dba.GetPostsPhotoForOne(*in)

	if err!= nil{
		return &pb.GetPostsPhotosPerParentResponse{
			Success:              false,
			PlacesPhoto:          nil,

		},err
	}
	return &pb.GetPostsPhotosPerParentResponse{
		Success:              true,
		PlacesPhoto:          photos,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil
}

func (s *server) GetPlacesPerCityPhotoDBA(ctx context.Context, in *pb.GetPlacesPhotosPerCityRequest) (*pb.GetPlacesPhotosPerCityResponse, error) {
	log.Printf("Received: %v", "GetPlacesPerCityPhotoDBA")

	photos, err := dba.GetPlacePhotosPerCity(*in)

	if err!= nil{
		return &pb.GetPlacesPhotosPerCityResponse{
			Success:              false,
			PlacePhotos:          nil,

		},err
	}
	return &pb.GetPlacesPhotosPerCityResponse{
		Success:              true,
		PlacePhotos:          photos,

	},nil
}

func (s *server) GetVisitdCityPhotosDBA(ctx context.Context, in *pb.GetVisitedCitysDBARequest) (*pb.GetVisitedCitysDBAResponse, error) {
	log.Printf("Received: %v", "GEt Visitd City Photos")

	response, err := dba.GetVIsitedCitysPhotos(*in)
	if err != nil{
		return nil,err
	}

	return &pb.GetVisitedCitysDBAResponse{
		CityPhotos:           response,

	},nil
}

func (s *server) GetVisitedPlacesPhotosDBA(ctx context.Context, in *pb.GetVisitedPlacesPhotoDBARequest) (*pb.GetVisitedPlacesPhotosDBAResponse, error) {
	log.Printf("Received: %v", "Get Visitid places photos")

	response, err := dba.GetVisitdPlacePhoto(*in)

	if err != nil{
		return nil,err
	}
	return &pb.GetVisitedPlacesPhotosDBAResponse{
		PlacePhotos:          response,

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


	res,_ :=dba.GetVIsitedCitysPhotos(pb.GetVisitedCitysDBARequest{
		CityId:               []int32{1,6},

	})
	fmt.Println(res)

	res1,_ := dba.GetVisitdPlacePhoto(pb.GetVisitedPlacesPhotoDBARequest{
		PlaceId:              []int32{5,3},

	})
	fmt.Println(res1)
	log.Print("Starting Service")
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
	fmt.Print("Done")

}
