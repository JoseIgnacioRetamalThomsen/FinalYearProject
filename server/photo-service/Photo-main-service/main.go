package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"strconv"

	//"google.golang.org/grpc"
	"io"

	"log"
	//"net"
	"os"
	"time"
)
var m = make(map[string]string)
var mc = make(map[string][]string)
var mp = make(map[string][]string)
var mpo = make (map[string][]string)
const (
	port = ":30051"

)
const(

	defaultProfile = "https://storage.googleapis.com/wcity-images-1/profile-1/profile_0.jpg"
)
const(
	google_api_key = "GOOGLE_APPLICATION_CREDENTIALS"
	start_url =  "https://storage.googleapis.com/"
)



type server struct {
	pb.UnimplementedPhotosServiceServer
}

func (s *server) GetProfilePhoto(ctx context.Context, in *pb.ProfilePhotoRequestP) (*pb.ProfilePhotoResponseP, error) {

	var url string

	if m[in.Email] =="" {
		url = defaultProfile

	}else{
		url = m[in.Email]
	}
	//

return &pb.ProfilePhotoResponseP{
	Email: in.Email,
	Valid: true,
	Url:   url,

},nil

}

func (s *server) UploadProfilePhoto(ctx context.Context, in *pb.ProfileUploadRequestP) (*pb.ProfileUploadResponseP, error) {

	randonIden1 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	randonIden2 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	fileName := profileFolders[0]+ "/"+randonIden1 + "_" +randonIden2+".jpg"
	url := start_url+ buckets[0]+"/" + fileName

	Write(in.Image,buckets[0],fileName)
	m[in.Email] = url

	return &pb.ProfileUploadResponseP{
		Email:                "",
		Valid:                true,
		Url:                  url,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil

}

func (s *server) UploadCityPhoto(ctx context.Context, in *pb.CityUploadRequestP) (*pb.CityUploadResponseP, error) {
	randonIden1 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	randonIden2 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	fileName := cityFolder[0]+ "/"+randonIden1 + "_" +randonIden2+".jpg"
	url := start_url+ buckets[0]+"/" + fileName

	Write(in.Image,buckets[0],fileName)
	mc[in.CityId] = append(mc[in.CityId],url)

	return &pb.CityUploadResponseP{
		Valid:                true,
		CityID:               in.CityId,
		Url:                  url,

	},nil
}

func (s *server) GetCityImage(ctx context.Context, in *pb.CityPhotoRequestP) (*pb.CityPhotoResponseP, error) {
	//var urls [] string;
	urls := mc[in.CityId]

	return &pb.CityPhotoResponseP{
		Valid:                true,
		CityID:               in.CityId,
		Url:                  urls,
		Active:               0,

	},nil
}


func (s *server) UploadPlacePhoto(ctx context.Context, in *pb.PlaceUploadRequestP) (*pb.PaceUploadResponseP, error) {
	randonIden1 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	randonIden2 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	fileName := placeFoler[0]+ "/"+randonIden1 + "_" +randonIden2+".jpg"
	url := start_url+ buckets[0]+"/" + fileName

	Write(in.Image,buckets[0],fileName)
	mp[in.PlaceId] = append(mp[in.PlaceId],url)

	return &pb.PaceUploadResponseP{
		Valid:                true,
		PlaceId:               in.PlaceId,
		Url:                  url,

	},nil
}

func (s *server) GetPlacePhoto(ctx context.Context, in *pb.PlacePhotoRequestP) (*pb.PlacePhotoResponseP, error) {
	//var urls [] string;
	urls := mp[in.PlaceId]

	return &pb.PlacePhotoResponseP{
		Valid:                true,
		PlaceId:              in.PlaceId,
		Url:                  urls,
		Active:               0,
	},nil
}

func (s *server) UploadPostImage(ctx context.Context, in *pb.PostUploadRequestP) (*pb.PostPhotoResponseP, error) {
	randonIden1 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	randonIden2 := strconv.Itoa(rand.Intn(999999 ) + 1000000)
	fileName := postFolder[0]+ "/"+randonIden1 + "_" +randonIden2+".jpg"
	url := start_url+ buckets[0]+"/" + fileName

	Write(in.Image,buckets[0],fileName)
	mp[in.PostId] = append(mc[in.PostId],url)

	return &pb.PostPhotoResponseP{
		Valid:                true,
		PostId:               in.PostId,
		UserEmail:            in.UserEmail,
		Type:                 in.Type,
		Name:                 in.Name,
		Url:                  url,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil
}

func (s *server) GetPostImage(ctx context.Context, in *pb.PostPhotoRequestP) (*pb.PostPhotoResponseP, error) {

	var url string

	if len(mpo[in.PostId]) == 0  {
		url = defaultProfile

	}else{
		url = mpo[in.PostId][0]
	}
	//

	return &pb.PostPhotoResponseP{
		Valid:                true,
		PostId:               in.PostId,
		UserEmail:            in.UserEmail,
		Type:                 in.Type,
		Name:                 in.Name,
		Url:                  url,
	},nil

}




func Save(){
	/*
	ctx := context.Background()
	f, err := os.Open("notes.txt")
	if err != nil {
		return  nil, err
	}
	defer f.Close()*/
}


func main(){



/*
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "final-year-gmit"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "my-new-bucket-7777777"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)*/


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


	var buckets = []string{"wcity-images-1"}
	var profileFolders = []string{"profile-1"}
	var cityFolder = []string{"city-1"}
	var placeFoler = []string{"place-1"}
	var postFolder = []string{"post-1"}





func Write(image []byte,bucketName string,imageName string){
	r := bytes.NewReader(image)
	projectID := os.Getenv(google_api_key)

	if projectID == "" {
		log.Println(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(bucketName).Object(imageName).NewWriter(ctx)
	if _, err = io.Copy(wc, r); err != nil {
		panic( err)
	}
	if err := wc.Close(); err != nil {
		panic(err)
	}
}

//https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}