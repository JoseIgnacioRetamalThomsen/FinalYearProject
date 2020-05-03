package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
var mpo = make(map[string][]string)

const (
	port = ":30051"
)
const (
	defaultProfile = "https://storage.googleapis.com/wcity-images-1/profile-1/profile_0.jpg"
)
const (
	google_api_key = "GOOGLE_APPLICATION_CREDENTIALS"
	start_url      = "https://storage.googleapis.com/"
)
const (
	DBA_URL = "35.197.221.57:7172"

	//DBA_URL  = "0.0.0.0:7172"
	AUTH_URL = "35.197.242.214:50051"
)

const (
	MAX_DB_CON_TIME = 3
)

type server struct {
	pb.UnimplementedPhotosServiceServer
}

var dbaConn photosDBAServer

type photosDBAServiceContext struct {
	dbClient pb.PhotosDBAServiceClient
	timeout  time.Duration
}
type photosDBAServer struct {
	context *photosDBAServiceContext
}

func newPhotosDBAServiceContext(endpoint string) (*photosDBAServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &photosDBAServiceContext{
		dbClient: pb.NewPhotosDBAServiceClient(userConn),
		timeout:  2 * time.Second,
	}
	return ctx, nil
}

func (s *server) GetProfilePhoto(ctx context.Context, in *pb.ProfilePhotoRequestP) (*pb.ProfilePhotoResponseP, error) {
	log.Printf("Received: %v", "get profile picture")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.ProfilePhotoResponseP{
			Valid: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.GetProfilePhotoDBA(ctx, &pb.GetProfilePhotosDBARequest{
		UserEmail:            in.Email,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	if err != nil {
		return &pb.ProfilePhotoResponseP{
			Valid: false,
		}, err
	}

	return &pb.ProfilePhotoResponseP{
		Email:  in.Email,
		Valid:  true,
		Photos: res.Photos,
	}, nil

}

func (s *server) UploadProfilePhoto(ctx context.Context, in *pb.ProfileUploadRequestP) (*pb.ProfileUploadResponseP, error) {

	log.Printf("Received: %v", "Upload profile picture")

	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.ProfileUploadResponseP{
			Sucess: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	randonIden1 := strconv.Itoa(rand.Intn(999999) + 1000000)
	randonIden2 := strconv.Itoa(rand.Intn(999999) + 1000000)
	fileName := profileFolders[0] + "/" + randonIden1 + "_" + randonIden2 + ".jpg"
	url := start_url + buckets[0] + "/" + fileName

	Write(in.Image, buckets[0], fileName)

	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.AddProfilePhotoDBA(ctx, &pb.AddProfilePhotoDBARequest{
		UserEmail: in.Email,
		Url:       url,
		Selected:  true,
	})
	if err != nil {
		log.Printf("Error: %v", err)
		return &pb.ProfileUploadResponseP{
			Sucess: false,
		}, err
	}
	photo := &pb.ProfilePhoto{
		Id:        res.Id,
		UserEmail: in.Email,
		Url:       url,
		Timestamp: res.TimeStamp,
		Selected:  true,
	}
	return &pb.ProfileUploadResponseP{
		Sucess: true,
		Photo:  photo,
	}, nil
}
func getUrlEnd() string {
	return "/" + strconv.Itoa(rand.Intn(999999)+1000000) + "_" + strconv.Itoa(rand.Intn(999999)+1000000) + ".jpg"
}

func (s *server) UploadCityPhoto(ctx context.Context, in *pb.CityUploadRequestP) (*pb.CityUploadResponseP, error) {

	log.Printf("Received: %v", "Upload city photo")

	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.CityUploadResponseP{
			Sucess: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	fileName := cityFolder[0] + getUrlEnd()
	url := start_url + buckets[0] + "/" + fileName

	Write(in.Image, buckets[0], fileName)

	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.AddCityPhotoDBA(ctx, &pb.AddCityPhotoDBARequest{
		CityId:   in.CityId,
		Url:      url,
		Selected: false,
	})

	if err != nil {
		return &pb.CityUploadResponseP{
			Sucess: false,
		}, err
	}

	return &pb.CityUploadResponseP{
		Sucess: true,
		Photo: &pb.CityPhoto{
			Id:        res.Id,
			CityId:    in.CityId,
			Url:       url,
			Timestamp: res.TimeStamp,
			Selected:  false,
		},
	}, nil
}

func (s *server) GetCityImage(ctx context.Context, in *pb.CityPhotoRequestP) (*pb.CityPhotoResponseP, error) {
	log.Printf("Received: %v", "get city picture")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.CityPhotoResponseP{
			Valid: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.GetCityPhotoDBA(ctx, &pb.GetCityPhotosDBARequest{
		CityId: in.CityId,
	})

	if err != nil {
		return &pb.CityPhotoResponseP{
			Valid: false,
		}, err
	}

	return &pb.CityPhotoResponseP{
		Valid:  true,
		CityID: in.CityId,
		Photos: res.Photos,
		Active: 0,
	}, nil

}

/*
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
*/

func (s *server) UploadPlacePhoto(ctx context.Context, in *pb.PlaceUploadRequestP) (*pb.PlaceUploadResponseP, error) {

	log.Printf("Received: %v", "Upload place photo")

	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.PlaceUploadResponseP{
			Success: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}
	fileName := placeFoler[0] + getUrlEnd()
	url := start_url + buckets[0] + "/" + fileName

	Write(in.Image, buckets[0], fileName)

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.AddPlacePhotoDBA(ctx, &pb.AddPlacePhotoDBARequest{
		PlaceId:  in.PlaceId,
		Url:      url,
		Selected: false,
		PlaceCityId: in.PlaceCityId,
	})

	if err != nil {
		return &pb.PlaceUploadResponseP{
			Success: false,
		}, err
	}

	return &pb.PlaceUploadResponseP{
		Success: true,
		Photo: &pb.PlacePhoto{
			Id:        res.Sd,
			PlaceId:   in.PlaceId,
			Url:       url,
			Timestamp: res.TimeStamp,
			Selected:  false,
		},
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func (s *server) GetPlacePhoto(ctx context.Context, in *pb.PlacePhotoRequestP) (*pb.PlacePhotoResponseP, error) {
	log.Printf("Received: %v", "get place picture")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.PlacePhotoResponseP{
			Valid: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.GetPlacePhotoDBA(ctx, &pb.GetPlacePhotosDBARequest{
		PlaceId: in.PlaceId,
	})

	if err != nil {
		return &pb.PlacePhotoResponseP{
			Valid: false,
		}, err
	}
	return &pb.PlacePhotoResponseP{
		Valid:   true,
		PlaceId: in.PlaceId,
		Photos:  res.Photos,
		Active:  false,
	}, nil
}

func (s *server) UploadPostImage(ctx context.Context, in *pb.PostUploadRequestP) (*pb.PostUploadResponseP, error) {

	log.Printf("Received: %v", "Upload post picture")

	//check token
	valid := CheckToken(in.UserEmail, in.Token)

	if !valid {
		return &pb.PostUploadResponseP{
			Sucess: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	//save to bucker
	fileName := postFolder[0] + getUrlEnd()
	url := start_url + buckets[0] + "/" + fileName
	Write(in.Image, buckets[0], fileName)

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.AddPostPhotoDBA(ctx, &pb.AddPostPhotoDBARequest{
		PostId:   in.PostId,
		Url:      url,
		Selected: false,
		Type : in.Type,
		ParentId: in.ParentId,
	})
	if err != nil {
		return &pb.PostUploadResponseP{
			Sucess: false,
		}, err
	}

	return &pb.PostUploadResponseP{
		Sucess: true,
		Photo: &pb.PostPhoto{
			Id:        res.Id,
			PostId:    in.PostId,
			Url:       url,
			Timestamp: res.TimeStamp,
			Selected:  false,
		},
	}, nil
}

func (s *server) GetPostImage(ctx context.Context, in *pb.PostPhotoRequestP) (*pb.PostPhotoResponseP, error) {

	log.Printf("Received: %v", "get Post picture")

	//check token
	valid := CheckToken(in.UserEmail, in.Token)

	if !valid {
		return &pb.PostPhotoResponseP{
			Valid: false,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	//get urls from database
	//database conection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_DB_CON_TIME)
	defer cancel()

	res, err := dbaConn.context.dbClient.GetPostPhotoDBA(ctx, &pb.GetPostPhotosDBARequest{
		PlaceId: in.PostId,
	})

	if err != nil {
		return &pb.PostPhotoResponseP{
			Valid: false,
		}, err
	}

	return &pb.PostPhotoResponseP{
		Valid:     true,
		PostId:    in.PostId,
		UserEmail: in.UserEmail,
		Photos:    res.Photos,
	}, nil

}

func (s *server) GetCitysPhotosP(ctx context.Context, in *pb.GetCitysPhotoRequestP) (*pb.GetCitysPhotoResponseP, error) {

	log.Printf("Received: %v", "get al city images")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.GetCitysPhotoResponseP{
			Success:    false,
			CityPhotos: nil,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res, err := dbaConn.context.dbClient.GetCitysPhotosDBA(ctx, &pb.GetCitysPhotoRequest{
		Valid: true,
	})
	if err != nil {
		return &pb.GetCitysPhotoResponseP{
			Success:    false,
			CityPhotos: nil,
		}, err
	}

	return &pb.GetCitysPhotoResponseP{
		Success:    true,
		CityPhotos: res.CityPhotos,
	}, nil
}

func (s *server) GetPostsPhotosIdP(ctx context.Context, in *pb.GetPostsPhotosPerParentRequestP) (*pb.GetPostsPhotosPerParentResponseP, error) {

	log.Printf("Received: %v", "get post image for parent")
	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.GetPostsPhotosPerParentResponseP{
			Success:     false,
			PlacesPhoto: nil,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res, err := dbaConn.context.dbClient.GetPostsPhotosIdDBA(ctx, &pb.GetPostsPhotosPerParentRequest{
		Type:     in.Type,
		ParentId: in.ParentId,
	})

	if err != nil {
		return &pb.GetPostsPhotosPerParentResponseP{
			Success:     false,
			PlacesPhoto: nil,
		}, err
	}

	return &pb.GetPostsPhotosPerParentResponseP{
		Success:     true,
		PlacesPhoto: res.PlacesPhoto,
	}, nil

}

func (s *server) GetPlacesPerCityPhotoP(ctx context.Context, in *pb.GetPlacesPhotosPerCityRequestP) (*pb.GetPlacesPhotosPerCityResponseP, error) {

	log.Printf("Received: %v", "get all Post cityu ")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return &pb.GetPlacesPhotosPerCityResponseP{
			Success:     false,
			PlacePhotos: nil,
		}, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res, err := dbaConn.context.dbClient.GetPlacesPerCityPhotoDBA(ctx, &pb.GetPlacesPhotosPerCityRequest{
		CityPlaceId:          in.PlaceId,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	if err != nil {
		return &pb.GetPlacesPhotosPerCityResponseP{
			Success:     false,
			PlacePhotos: nil,
		}, err
	}

	return &pb.GetPlacesPhotosPerCityResponseP{
		Success:     true,
		PlacePhotos: res.PlacePhotos,
	}, nil
}

func (s *server) GetVisitedCitysPhotos(ctx context.Context, in *pb.GetVisitedCitysImagesRequest) (*pb.GetCitysPhotoResponseP, error) {

	log.Printf("Received: %v", "Get visited citys photos")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}
	res, err := dbaConn.context.dbClient.GetVisitdCityPhotosDBA(ctx,&pb.GetVisitedCitysDBARequest{
		CityId:               in.CityId,})
	if err!= nil{
		return nil,err
	}

	return &pb.GetCitysPhotoResponseP{
		Success:              true,
		CityPhotos:           res.CityPhotos,

	},nil

}

func (s *server) GetVisitedPlacesPhotos(ctx context.Context, in *pb.GetVisitedPlacesPhotosRequest) (*pb.GetVisitedPlacesPhotosResponse, error) {

	log.Printf("Received: %v", "Get visited places photos")

	//check token
	valid := CheckToken(in.Email, in.Token)

	if !valid {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	res,err := dbaConn.context.dbClient.GetVisitedPlacesPhotosDBA(ctx,&pb.GetVisitedPlacesPhotoDBARequest{
		PlaceId:              in.PlaceId,

	})

	if err!= nil{
		return nil,err
	}
	return &pb.GetVisitedPlacesPhotosResponse{
		Success:              true,
		PlacePhotos:          res.PlacePhotos,

	},nil

}

func main() {

	//conect to database server
	dbserverCtx, err := newPhotosDBAServiceContext(DBA_URL)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &photosDBAServer{dbserverCtx}
	dbaConn = *s2

	//conect to ps
	// start server pass connection
	psserverCtx, err := newAuthClientContext(AUTH_URL)
	if err != nil {
		log.Fatal(err)
	}
	s1 := &authClient{psserverCtx}
	prsCon = *s1

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

func Write(image []byte, bucketName string, imageName string) {
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
		panic(err)
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

/*
 CHECK TOKEN
*/
// password service client
type authClient struct {
	context *authClientContext
}

type authClientContext struct {
	psClient pb.UserAuthenticationClient
	timeout  time.Duration
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
		timeout:  time.Second,
	}
	return ctx, nil
}

func CheckToken(email string, token string) bool {
	// Set up a connection to the server.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := prsCon.context.psClient.CheckToken(ctx, &pb.LogRequest{
		Token: token,
		Email: email,
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return r.Sucess
}
