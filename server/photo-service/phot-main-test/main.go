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
	url = "35.197.216.42:30051"
	//url="35.197.216.42:30051";
	//url = "35.234.146.99:5777"
	token ="ef236fdcb42d55d00703a8737d343574b9766c7a78fd5a268425f6bbe8753b9e"
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
		timeout:  time.Second*2,
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
	//AddProfileImage(dat,tokenEmail,token)
	//GetProfilePhoto(tokenEmail,token);
	SendCityimage(dat,tokenEmail,1,token)
GetCityIamge(1,tokenEmail,token)
	//fmt.Println(GetProfilePhoto("d","token"))
	//endPlaceImage(dat,2,tokenEmail,token)

//GetPlacePhotos(5,tokenEmail,token)
//SendPost(dat,"id1",tokenEmail,token)
//GetPostImage("id1",tokenEmail,token)
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
fmt.Println(r)
	return "",nil
}

func AddProfileImage(image []byte,email string,token string) string{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadProfilePhoto(ctx,&pb.ProfileUploadRequestP{
		Email : email,
		Token: token,
		Image : image,
	})
	if err!= nil{
		panic(err)
	}
	return r.Photo.Url
}
/*
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
*/
func SendCityimage(image []byte,email string, cityId int,token string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadCityPhoto(ctx,&pb.CityUploadRequestP{
		Token : token,
		Email : email,
		CityId : int32(cityId),
		Image : image,
	})

	if err!= nil{
		panic(err)
	}
	fmt.Print(r)
}

func GetCityIamge(cityId int,email string , token string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetCityImage(ctx,&pb.CityPhotoRequestP{
		CityId : int32(cityId),
		Token : token,
		Email : email,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func SendPlaceImage(image []byte,placeId int,email string, token string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadPlacePhoto(ctx, &pb.PlaceUploadRequestP{
		PlaceId :int32(placeId),
		Image : image,
		Token : token,
		Email : email,
	})

	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func GetPlacePhotos(placeId int,email string,token string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetPlacePhoto(ctx, &pb.PlacePhotoRequestP{
		PlaceId : int32(placeId),
		Token : token,
		Email : email,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func SendPost(image []byte,postId string,email string, token string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.UploadPostImage(ctx, &pb.PostUploadRequestP{
		PostId : postId,
		Image : image,
		UserEmail: email,
		Token : token,
	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

func GetPostImage(postId string,email string, token string){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := photoConn.context.dbClient.GetPostImage(ctx, &pb.PostPhotoRequestP{
		PostId : postId,
		UserEmail : email,
		Token : token,

	})
	if err!= nil{
		panic(err)
	}
	fmt.Println(r)
}

