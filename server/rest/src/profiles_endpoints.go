package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"io/ioutil"
	"log"
	"net/http"
)

type CityResponseJson struct{
	City pb.City
	Photo pb.CityPhoto
}

type CityResponse struct {
	Success   bool      `json:"success"`
	City pb.CityResponseP `json:"city"`
}

func CreateCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Create city")

	err2 := r.ParseForm()

	if(err2!=nil){
		fmt.Print("err:",err2)
	}
	//image := []byte(r.Form["image"][0])




	for _,value := range  r.Header{
		fmt.Println(value)
	}

	// create the city grpc city object
	var city pb.City
	//read body data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}

	// create grpc coity using body data
	json.Unmarshal(reqBody, &city)

	//create Request
	var cityRequest pb.CreateCityRequestP

	cityRequest.Token = r.Header["Token"][0]
	cityRequest.Name = r.Header["Email"][0] // name is email
	cityRequest.City = &city;
	fmt.Println(city)

	response,err := CreateCity(cityRequest)

	if err !=nil {
		fmt.Println(err)
	}

	if val, ok := r.Form["image"]; ok {
		image := []byte(val[0])

		cityId := response.City.CityId

		photo , err :=SendCityimage(image, r.Header["Email"][0],r.Header["Token"][0], int(cityId))
		if err !=nil {
			fmt.Println(err)
		}
		jsonResponse := CityResponseJson{
			City:   *response.City,
			Photo: *photo,
		}

		json.NewEncoder(w).Encode(jsonResponse)
	}else {

		//convert reponse into json and send back
		jsonResponse := CityResponseJson{
			City:  *response.City,
			Photo: pb.CityPhoto{}, //*photo,
		}

		fmt.Println("next")
		fmt.Println(jsonResponse)

		json.NewEncoder(w).Encode(jsonResponse)
	}

}

func GetCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get city")

	response, err := GetCity(pb.GetCityRequestP{
		Token:                r.Header["Token"][0],
		Name:                 r.Header["Email"][0],
		CityName:             mux.Vars(r)["name"],
		CityCountry:          mux.Vars(r)["country"],

	})
	/*
		response, err := GetCity(pb.CityRequestP{
			Token:                r.Header["Token"][0],
			Name:                 mux.Vars(r)["name"],
			Country:              mux.Vars(r)["country"],
			CreatorEmail:         r.Header["Email"][0],

		})*/
	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}
	fmt.Println(response.City)
	json.NewEncoder(w).Encode(response.City)
}

func UpdateCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Update request")

	var request pb.CreateCityRequestP
	var city pb.City
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain request data")
	}
	json.Unmarshal(reqBody, &city)

	fmt.Println(request)


	request.Token = r.Header["Token"][0]
	//request.CreatorEmail =  r.Header["Email"][0]
	request.Name =  mux.Vars(r)["name"]
	//request.Country= mux.Vars(r)["country"]

	request.City = &city

	fmt.Println(request)
	response, err := UpdateCity(request)

	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}

	json.NewEncoder(w).Encode(response)

}


func CreatePlaceRequest(w http.ResponseWriter, r *http.Request){


	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	if len(email) ==0 || len(token) == 0  {
		http.Error(w, "Wrong request", 400)
		return
	}

	var request pb.CreatePlaceRequestP
	// create the city grpc city object
	var place pb.Place
	//read body data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}
	request.Place = &place

	request.Token = r.Header["Token"][0]
	//request.CreatorEmail =  r.Header["Email"][0]
	request.Name =  mux.Vars(r)["name"]

	// create grpc coity using body data
	json.Unmarshal(reqBody, &place)
	response, err := CreatePlace(request)

	placeId := response.Place.PlaceId
	placeId = placeId

	json.NewEncoder(w).Encode(response)


}
