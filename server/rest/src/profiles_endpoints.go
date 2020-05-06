package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type CityResponseJson struct {
	City  pb.City
	Photo pb.CityPhoto
}

type CityResponse struct {
	Success bool             `json:"success"`
	City    pb.City `json:"city"`
	Images []*pb.CityPhoto
	Posts []*pb.CityPost

}

func CreateCityRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received: %v", "Create city")

	err2 := r.ParseForm()

	if err2 != nil {
		fmt.Print("err:", err2)
	}
	//image := []byte(r.Form["image"][0])

	for _, value := range r.Header {
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
	cityRequest.City = &city
	fmt.Println(city)

	response, err := CreateCity(cityRequest)

	if err != nil {
		fmt.Println(err)
	}

	if val, ok := r.Form["image"]; ok {
		image := []byte(val[0])

		cityId := response.City.CityId

		photo, err := SendCityimage(image, r.Header["Email"][0], r.Header["Token"][0], int(cityId))
		if err != nil {
			fmt.Println(err)
		}
		jsonResponse := CityResponseJson{
			City:  *response.City,
			Photo: *photo,
		}

		json.NewEncoder(w).Encode(jsonResponse)
	} else {

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

func GetCityRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received: %v", "Get city")

token:=       r.Header["Token"][0]
	name :=        r.Header["Email"][0]
		cityName:=    mux.Vars(r)["name"]
		cityCountry:= mux.Vars(r)["country"]

	response, err := GetCity(pb.GetCityRequestP{
		Token:       token,
		Name:        name,
		CityName:    cityName,
		CityCountry: cityCountry,
	})

	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}

	//get images
	images,err := GetCityimage(pb.CityPhotoRequestP{
		Token:                token,
		Email:                name,
		CityId:               response.City.CityId,

	})

	var city CityResponse;
	if images == nil {
		city = CityResponse{
			Success: true,
			City:    *response.City,
			Images:  nil,
		}
	}else {
		city = CityResponse{
			Success: true,
			City:    *response.City,
			Images:  images.Photos,
		}
	}

	//get posts
	posts,err := GetCityPosts(pb.PostsRequest{
		IndexId:              response.City.CityId,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	if posts != nil{
		city.Posts = posts.Posts;
	}
	fmt.Println(response.City)
	json.NewEncoder(w).Encode(city)
}

func UpdateCityRequest(w http.ResponseWriter, r *http.Request) {
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
	request.Name = mux.Vars(r)["name"]
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

func CreatePlaceRequest(w http.ResponseWriter, r *http.Request) {

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	if len(email) == 0 || len(token) == 0 {
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
	request.Name = mux.Vars(r)["name"]

	// create grpc coity using body data
	json.Unmarshal(reqBody, &place)
	response, err := CreatePlace(request)

	placeId := response.Place.PlaceId
	placeId = placeId

	json.NewEncoder(w).Encode(response)

}
type AllCityResponse struct{
	Cities *[]pb.City
	Images []*pb.CitysPhoto
}
func GetAllCityEndPoint(w http.ResponseWriter, r *http.Request) {

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	search := r.FormValue("search")

	if len(search) > 0 {

		sr, err := Search(pb.SearchAllRequest{
			Max:                  1000,
			Search:               search,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(sr)
	} else
	{
		cities, err := GetAllCities(pb.GetAllRequest{
			Max: 1000,
		})

		cr := AllCityResponse{
			Cities:cities}

		if err != nil {
			fmt.Fprintf(w, "Wrong request, body must contain city data")
		}

		images,err := GetAllCityImages(pb.GetCitysPhotoRequestP{
			Email:                email,
			Token:                token,

		})
		if err==nil{
			cr.Images= images.CityPhotos
		}

		json.NewEncoder(w).Encode(cr)
	}
}

func GetCityPlacesEndPoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get City Places  : %s", r)

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	if len(email) == 0 || len(token) == 0 {
		http.Error(w, "Wrong request auth", 400)
		return
	}

	cityName := mux.Vars(r)["city"]
	cityCountry := mux.Vars(r)["country"]

	var request = pb.CreateCityRequestP{
		Token: token,
		Name:  email,
		City: &pb.City{
			Name:    cityName,
			Country: cityCountry,
		},
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	response, err := GetAllCityPlaces(request)

	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}

	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Get User")

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	fmt.Print(token)
	fmt.Print(email)
	user, err := GetUserProfile(pb.GetUserRequestP{
		Token: token,
		Email: email,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetPlaceRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received: %v", "Get place")

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]
	country := mux.Vars(r)["country"]
	city := mux.Vars(r)["city"]
	name := mux.Vars(r)["name"]
	response, err := GetPlace(pb.GetPlaceRequestP{
		Token:                token,
		Email:                email,
		PlaceName:            name,
		PlaceCity:           city,
		PlaceCountry:         country,

	})
	if err != nil {
		log.Printf("Error: %v", err)
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}

	json.NewEncoder(w).Encode(response)
}

func VisitCityEndPoint(w http.ResponseWriter, r *http.Request) {

		log.Printf("Received: %v", "Visit city")
	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	i, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	res,err := VisitCity(pb.VisitCityRequestP{
		Token:                token,
		Email:                email,
		Id:                   int32(i),

	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)


}
