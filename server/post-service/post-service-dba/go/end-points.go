package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

/**
*  END POINTS
 */
func (s *server) CreateCityPost(ctx context.Context, in *pb.CityPostPSDB) (*pb.CreatePostResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Create city post", in.String())

	index,err := CreateCityPost(&CityPost{
		IndexId:      in.IndexId,
		CreatorEmail: in.CreatorEmail,
		Name:         in.CityName,
		Country:      in.CityCountry,
		Title:        in.Title,
		Body:         in.Body,
		TimeStamp:    in.TimeStamp,
		Likes:        nil,
		MongoId:      "",
	})
	if err != nil{
		return &pb.CreatePostResponsePSDB{
			Valied:               false,
			MongoId:              fmt.Sprintf("%v",index),

		},err
	}
	index = index
	return &pb.CreatePostResponsePSDB{
		Valied:               true,
		MongoId:              fmt.Sprintf("%v",index),

	},nil
}

func (s *server) CreatePlacePost(ctx context.Context, in *pb.PlacePostPSDB) (*pb.CreatePostResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Create place post", in.String())

	index, err := CreatePlacePost(&PlacePost{
		IndexId:      in.IndexId,
		CreatorEmail: in.CreatorEmail,
		Name:         in.PlaceName,
		City:         in.CityName,
		Country:      in.CountryName,
		Title:        in.Title,
		Body:         in.Body,
		TimeStamp:    in.TimeStamp,
		Likes:        nil,

	})

	if err != nil{
		return &pb.CreatePostResponsePSDB{
			Valied:               false,
			MongoId:              fmt.Sprintf("%v",index),

		},err
	}

	return &pb.CreatePostResponsePSDB{
		Valied:               true,
		MongoId:              fmt.Sprintf("%v",index),

	},nil
}

func (s *server) GetPlacePosts(ctx context.Context, in *pb.PostsRequestPSDB) (*pb.PlacePostsResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Get place posts", in.String())

	posts := GetPlacePost(in.IndexId)

	if posts == nil{
		return &pb.PlacePostsResponsePSDB{
			Valid:                false,
			IndexId:              0,
			Posts:                nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},status.Error(codes.NotFound,"No post found.")
	}

	return &pb.PlacePostsResponsePSDB{
		Valid:                true,
		IndexId:              in.IndexId,
		Posts:                posts,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil

}
func (s *server) GetCityPosts(ctx context.Context, in *pb.PostsRequestPSDB) (*pb.CityPostsResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Get city posts", in.String())

	posts := GetCityPost(in.IndexId)

	if posts == nil{
		return &pb.CityPostsResponsePSDB{
			Valid:                false,
			IndexId:              in.IndexId,
			Posts:                nil,

		},status.Error(codes.NotFound,"No post found.")

	}
	return &pb.CityPostsResponsePSDB{
		Valid:                true,
		IndexId:              in.IndexId,
		Posts:                posts,

	},nil
}

