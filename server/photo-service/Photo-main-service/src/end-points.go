package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"strconv"
	"time"
)

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
