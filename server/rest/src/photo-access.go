package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"time"
)

func SendCityimage(image []byte,email string,token string,cityId int)(*pb.CityPhoto,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadCityPhoto(ctx,&pb.CityUploadRequestP{
		Token : token,
		Email : email,
		CityId : int32(cityId),
		Image : image,
	})

	if err!= nil{
		return nil,err
	}

	return r.Photo,nil
}
//  rpc GetCityImage (CityPhotoRequestP) returns (CityPhotoResponseP);

func GetCityimage(request pb.CityPhotoRequestP)(*pb.CityPhotoResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetCityImage(ctx,&request)

	if err!= nil{
		return nil,err
	}

	return r,err;
}
//    rpc GetCitysPhotosP(GetCitysPhotoRequestP) returns (GetCitysPhotoResponseP);

func GetAllCityImages(request pb.GetCitysPhotoRequestP)(*pb.GetCitysPhotoResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetCitysPhotosP(ctx,&request)
	if err!= nil{
		return nil,err
	}

	return r,err;
}
