package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const(
	url = "0.0.0.0:2787"
	//url="35.197.216.42:60051";
	//url = "35.197.221.57:2787"
	token ="a31e31a2fcdf2a9a230120ea620f3b24f7379d923fb122323d3cb9bc56fe6508"
	tokenEmail ="a@a.com"
	Port = ":10051"
)
// dba connection
type postDBA struct {
	context *postDBAContext
}

type postDBAContext struct {
	dbClient pb.PostsServiceDBAClient
	timeout time.Duration
}
var dbaConn postDBA

// create connection
func newPostServiceContext(endpoint string) (*postDBAContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &postDBAContext{
		dbClient: pb.NewPostsServiceDBAClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}


//server
type server struct {
	pb.UnimplementedPostsServiceServer
}

/**
*  END POINTS
 */
func (s *server) CreateCityPost(ctx context.Context, in *pb.CityPost) (*pb.CreatePostResponse, error) {
	log.Printf("Received: %v , from: %v", "Get user", in.String())

	post := &pb.CityPostPSDB{
		IndexId:              in.IndexId,
		CreatorEmail:         in.CreatorEmail,
		CityName:             in.CityName,
		CityCountry:          in.CityCountry,
		Title:                in.Title,
		Body:                 in.Body,
		TimeStamp:            time.Now().Format("UnixDate"),
		Likes:                nil,
		MongoId:              "",

	}
	result :=CreateCityPost(*post)

	return &pb.CreatePostResponse{
		Valied:               result.Valied,
		MongoId:              result.MongoId,

	},nil
}

func (s *server) CreatePlacePost(ctx context.Context, in *pb.PlacePost) (*pb.CreatePostResponse, error) {
	log.Printf("Received: %v , from: %v", "Create city post", in.String())

	post := &pb.PlacePostPSDB{
		IndexId:              in.IndexId,
		PlaceName:   in.PlaceName,
		CreatorEmail:         in.CreatorEmail,
		CityName:             in.CityName,
		CountryName:          in.CountryName,
		Title:                in.Title,
		Body:                 in.Body,
		TimeStamp:            time.Now().Format("UnixDate"),
		Likes:                nil,
		MongoId:              "",

	}
	result :=CreatePlacePost(*post)

	return &pb.CreatePostResponse{
		Valied:               result.Valied,
		MongoId:              result.MongoId,

	},nil
}

func (s *server) GetPlacePosts(ctx context.Context, in *pb.PostsRequest) (*pb.PlacePostsResponse, error) {
	log.Printf("Received: %v , from: %v", "all posts", in.String())

	posts,err := GetPlacePosts(pb.PostsRequestPSDB{
		IndexId:              in.IndexId,

	})
	if err!= nil{
		return &pb.PlacePostsResponse{
			Valid:                false,
			IndexId:              0,
			Posts:                nil,

		},err
	}

	var postsRes []*pb.PlacePost

	for _,value := range posts.Posts{
		postsRes = append(postsRes,&pb.PlacePost{
			IndexId:              value.IndexId,
			CreatorEmail:         value.CreatorEmail,
			CityName:             value.CityName,
			CountryName:          value.CountryName,
			PlaceName:            value.PlaceName,
			Title:                value.Title,
			Body:                 value.Body,
			TimeStamp:            value.TimeStamp,
			Likes:                value.Likes,
			MongoId:              value.MongoId,

		})
	}
	return &pb.PlacePostsResponse{
		Valid:                true,
		IndexId:              in.IndexId,
		Posts:                postsRes,

	},nil
}

func (s *server) GetCityPosts(ctx context.Context, in *pb.PostsRequest) (*pb.CityPostsResponse, error) {
	log.Printf("Received: %v , from: %v", "all posts", in.String())

	posts,err  := GetCityPosts(pb.PostsRequestPSDB{
		IndexId:              in.IndexId,

	})
	if err != nil{
		return &pb.CityPostsResponse{
			Valid:                false,
			IndexId:              in.IndexId,
			Posts:                nil,

		},err
	}

	var postsRes []*pb.CityPost

	for _,value := range posts.Posts{
		postsRes = append(postsRes,&pb.CityPost{
			IndexId:              value.IndexId,
			CreatorEmail:         value.CreatorEmail,
			CityName:             value.CityName,
			CityCountry:          value.CityCountry,
			Title:                value.Title,
			Body:                 value.Body,
			TimeStamp:            value.TimeStamp,
			Likes:                value.Likes,
			MongoId:              value.MongoId,

		})
	}
	return &pb.CityPostsResponse{
		Valid:                true,
		IndexId:              in.IndexId,
		Posts:                postsRes,

	},nil
}

func main(){

	fmt.Println(time.Now().Format("Mon Jan _2 15:04:05 MST 2006"))
	//conect to server
	dbserverCtx, err := newPostServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &postDBA{dbserverCtx}
	dbaConn = *s2

	temp := &pb.CityPostPSDB{
		IndexId:              1,
		CreatorEmail:         "email",
		CityName:             "aaaaaaaaaaaaaaaaaaaaa",
		CityCountry:          "b",
		Title:                "c",
		Body:                 "d",
		TimeStamp:            "e",
		Likes:                nil,
		MongoId:              "",

	}
	CreateCityPost(*temp)

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServiceServer(s, &server{})
	log.Print("Service Started in port: ", Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func CreateCityPost(post pb.CityPostPSDB) pb.CreatePostResponsePSDB{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbaConn.context.dbClient.CreateCityPost(ctx,&post)
	if err != nil {
		panic(err)
	}

	return *r
}


func CreatePlacePost(post pb.PlacePostPSDB) pb.CreatePostResponsePSDB{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbaConn.context.dbClient.CreatePlacePost(ctx,&post)
	if err != nil {
		panic(err)
	}

	return *r
}


func GetPlacePosts(request pb.PostsRequestPSDB) (pb.PlacePostsResponsePSDB, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbaConn.context.dbClient.GetPlacePosts(ctx,&request)
	if err!= nil{
		return pb.PlacePostsResponsePSDB{
			Valid:                false,
			IndexId:              0,
			Posts:                nil,

		},err
	}

	return *r, nil
}

func GetCityPosts(request pb.PostsRequestPSDB) (pb.CityPostsResponsePSDB,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbaConn.context.dbClient.GetCityPosts(ctx,&request)
	if err != nil{
		return pb.CityPostsResponsePSDB{
			Valid:                false,
			IndexId:              0,
			Posts:                nil,

		},err
	}
	return *r, nil
}
