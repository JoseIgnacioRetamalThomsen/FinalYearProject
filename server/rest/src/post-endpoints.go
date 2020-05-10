package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"log"
	"net/http"
	"strconv"
)


func GetCityPostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get city post")

	i1, err := strconv.Atoi(mux.Vars(r)["indexid"])
	if err !=nil{

	}

	response , err := GetCityPosts(pb.PostsRequest{
		IndexId:              int32(i1),

	})

	json.NewEncoder(w).Encode(response)
}


func GetPlacePostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get place post")

	i1, err := strconv.Atoi(mux.Vars(r)["indexid"])
	if err !=nil{

	}

	response , err := GetPlacePosts(pb.PostsRequest{
		IndexId:              int32(i1),

	})

	json.NewEncoder(w).Encode(response)
}


