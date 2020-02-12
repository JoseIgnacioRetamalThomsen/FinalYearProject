package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":30051"

)
const(

	testURL = "https://storage.googleapis.com/wcity-images-1/profile-1/profile_0.jpg"
)
type server struct {
	pb.UnimplementedPhotosServiceServer
}

func (s *server) GetProfilePhoto(ctx context.Context, in *pb.ProfilePhotoRequestP) (*pb.ProfilePhotoResponseP, error) {

return &pb.ProfilePhotoResponseP{
	Email:                in.Email,
	Valid:                true,
	Url:                  testURL,

},nil

}



func main(){

	//start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPhotosServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
