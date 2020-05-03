package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"log"
	"time"
)


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
