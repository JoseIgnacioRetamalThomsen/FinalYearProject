package main

import (
	"context"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"time"
)

func CreateCityPost(post pb.CityPost)(*pb.CreatePostResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.CreateCityPost(ctx,&post)
	if err != nil {
		return nil,err
	}

	return r,nil
}

func GetCityPosts(request pb.PostsRequest)(*pb.CityPostsResponse ,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.GetCityPosts(ctx,&request)
	if err!=nil{
		return nil,err
	}
	return r,nil
}

func GetPlacePosts(request pb.PostsRequest)(*pb.PlacePostsResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceConn.context.dbClient.GetPlacePosts(ctx,&request);
	if err!=nil{
		return nil,err
	}
	return r,nil
}

