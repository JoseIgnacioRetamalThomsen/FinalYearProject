package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"time"
)

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

