package main

import (
	"context"
	"fmt"
	"io/ioutil"

	//"fmt"
	//"io/ioutil"

	//"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"

	//"strconv"
	//"strings"

	//	"net"
	"time"

	//"errors"
	//"fmt"
	//pb "github.com/joseignacioretamalthomsen/wcity"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/resolver"
	//"time"
)

const(
	url = "0.0.0.0:30051"
	//url="35.197.216.42:30051";
	//url = "35.234.146.99:5777"
	token ="c1ce3461b81275c72c7dd7bbe6372bfcf099d83fb383ade531935ca4610cb4b6"
	tokenEmail ="a@a.com"
)

type photosServer struct {
	context *photosServiceContext
}

type photosServiceContext struct {
	dbClient pb.PhotosServiceClient
	timeout time.Duration
}
var photoConn photosServer

// create connection
func newPhotosServiceContext(endpoint string) (*photosServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &photosServiceContext{
		dbClient: pb.NewPhotosServiceClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
const(
	email = "a@a.com"
)
func main() {


	//conect to server
	dbserverCtx, err := newPhotosServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &photosServer{dbserverCtx}
	photoConn = *s2

	dat, err := ioutil.ReadFile("img/website.jpg")
	check(err)
	dat=dat
	//SendCityimage(dat,"d","1")
//GetCityIamge("1")
	//fmt.Println(GetProfilePhoto("d","token"))
	//SendPlaceImage(dat,"id")

//GetPlacePhotos("id")
//SendPost(dat,"id")
//GetPostImage("id")
}

func GetProfilePhoto(email string,token string) (string,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetProfilePhoto(ctx,&pb.ProfilePhotoRequestP{
		Email:                email,
		Token:                token,

	})
	if err != nil{
		panic(err)
	}

	return r.Url,nil
}


func SendImage(image []byte,email string)string{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadProfilePhoto(ctx,&pb.ProfileUploadRequestP{
		Image : image,
		Email : email,

	})
	if err!= nil{
		panic(err)
	}
	return r.Url
	//fmt.Print(string(image))
}

func SendCityimage(image []byte,email string, cityId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadCityPhoto(ctx,&pb.CityUploadRequestP{
		CityId : cityId,
		Image : image,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)

}

func GetCityIamge(cityId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetCityImage(ctx,&pb.CityPhotoRequestP{
		CityId : cityId,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func SendPlaceImage(image []byte,placeId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadPlacePhoto(ctx, &pb.PlaceUploadRequestP{
		PlaceId :placeId,
		Image : image,
	})

	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func GetPlacePhotos(placeId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetPlacePhoto(ctx, &pb.PlacePhotoRequestP{
		PlaceId : placeId,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func SendPost(image []byte,postId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadPostImage(ctx, &pb.PostUploadRequestP{
		PostId : postId,
		Image : image,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func GetPostImage(postId string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetPostImage(ctx, &pb.PostPhotoRequestP{
		PostId : postId,

	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}
