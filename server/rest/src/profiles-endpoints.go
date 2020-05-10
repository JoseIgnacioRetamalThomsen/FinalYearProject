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
	Success      bool    `json:"success"`
	City         pb.City `json:"city"`
	Images       []*pb.CityPhoto
	Posts        []*pb.CityPost
	PostImages   []*pb.PostPhoto
	Places       []*pb.Place
	PlacesImages []*pb.PlacesCityPhotos
}
type PostFull struct {
	indexId      int32
	CreatorEmail string
	Title        string
	Body         string
	TimeStamp    string
	MongoId      string
	ImageURL     string
}
type PlaceNoPost struct {
	PlaceId     int32
	Name        string
	City        string
	Location    pb.Geolocation
	Description string
	Images      []*pb.PlacePhoto
}
type CityFull struct {
	cityId       int32
	Name         string
	Country      string
	CreatorEmail string
	Location     pb.Geolocation
	Description  string
	Images       []*pb.CityPhoto
	Posts        [] PostFull
	Places       [] PlaceNoPost
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
		fmt.Print(("image"))
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

	token := r.Header["Token"][0]
	name := r.Header["Email"][0]
	cityName := mux.Vars(r)["name"]
	cityCountry := mux.Vars(r)["country"]

	var cityFull CityFull

	response, err := GetCity(pb.GetCityRequestP{
		Token:       token,
		Name:        name,
		CityName:    cityName,
		CityCountry: cityCountry,
	})

	// set city base data
	cityFull.cityId = response.City.CityId
	cityFull.CreatorEmail = response.City.CreatorEmail
	cityFull.Name = response.City.Name
	cityFull.Country = response.City.Country
	cityFull.Description = response.City.Description
	cityFull.Location = *response.City.Location

	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}

	//get images
	images, err := GetCityimage(pb.CityPhotoRequestP{
		Token:  token,
		Email:  name,
		CityId: response.City.CityId,
	})
	if err == nil {
		cityFull.Images = images.Photos
	}

	//get posts
	posts, err := GetCityPosts(pb.PostsRequest{
		IndexId: response.City.CityId,
	})

	// get post images
	postimages, err := GetCityPostImages(pb.GetPostsPhotosPerParentRequestP{
		Email:                name,
		Token:                token,
		Type:                 pb.PostType_CityTypePhoto,
		ParentId:             response.City.CityId,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	m := make(map[string]string)
	if err == nil {

		for _, img := range postimages.PlacesPhoto {
			m[img.PostId] = img.Url
		}
	}

	var postFull []PostFull
	if posts != nil {

		for _, p := range posts.Posts {

			postFull = append(postFull, PostFull{
				indexId:      p.IndexId,
				CreatorEmail: p.CreatorEmail,
				Title:        p.Title,
				Body:         p.Body,
				TimeStamp:    p.TimeStamp,
				MongoId:      p.MongoId,
				ImageURL:     m[p.MongoId],
			})
		}
	}

	cityFull.Posts = postFull
	//get places
	places, err := GetAllCityPlaces(pb.CreateCityRequestP{
		Token: token,
		Name:  name,
		City:  response.City,
	})

	placesImages, err := GetCityPlacesImages(pb.GetPlacesPhotosPerCityRequestP{
		Email:   name,
		Token:   token,
		PlaceId: response.City.CityId,
	})

	m1 := make(map[int32]([]*pb.PlacePhoto))
	if err == nil {
		for _, img := range placesImages.PlacePhotos {
			m1[img.PlaceId] = img.PlacePhotos
		}
	}

	var placesNoPost []PlaceNoPost
	if err == nil {
		for _, p := range places.Places {
			placesNoPost = append(placesNoPost, PlaceNoPost{
				PlaceId:     p.PlaceId,
				Name:        p.Name,
				City:        cityName,
				Location:    *p.Location,
				Description: p.Description,
				Images:      m1[p.PlaceId],
			})
		}
	}
	cityFull.Places = placesNoPost
	json.NewEncoder(w).Encode(cityFull)
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

type AllCityResponse struct {
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
			Cities: cities}

		if err != nil {
			fmt.Fprintf(w, "Wrong request, body must contain city data")
		}

		images, err := GetAllCityImages(pb.GetCitysPhotoRequestP{
			Email: email,
			Token: token,
		})
		if err == nil {
			cr.Images = images.CityPhotos
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

type CityOverview struct{
	CityId int32
	Name string
	Country string
	CreatorEmail string
	Description string
	Location pb.Geolocation
	Images  []*pb.CityPhoto
}

type PlaceOverview struct{
	PlaceId int32
	Name string
	City string
	Country string
	Location pb.Geolocation
	Description string
	Images []*pb.PlacePhoto

}

type UserProfile struct{

	Email string
	Name string
	Descripiton string
	Avatar []*pb.ProfilePhoto
	VisitedCities [] CityOverview
	VisitedPlaces []PlaceOverview
}
func GetUser(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Get User")

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	var userProfile UserProfile

	user, err := GetUserProfile(pb.GetUserRequestP{
		Token: token,
		Email: email,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userProfile.Name = user.User.Name
	userProfile.Email = user.User.Email
	userProfile.Descripiton = user.User.Descripiton

	//profile image
	img,err := GetProfilePhoto(pb.ProfilePhotoRequestP{
		Email:                email,
		Token:                token,

	})

	if err== nil{
		userProfile.Avatar = img.Photos
	}

	// get visited citys

	//images

	vc ,err := GetUserVisitedCities(pb.VisitedRequestP{
		Token:                token,
		Email:                email,

	})

	var citiesid []int32
	for _,c := range vc.Citys{
		citiesid = append(citiesid,c.CityId)
	}
	vcimg, err := GetVisitedCittiesImages(pb.GetVisitedCitysImagesRequest{
		Email:                email,
		Token:                token,
		CityId:               citiesid,

	})

	m :=  make(map[int32]([]*pb.CityPhoto))
	for _,img := range vcimg.CityPhotos{
		m[img.CitysPhotos[0].CityId] = img.CitysPhotos
	}

	for _,c := range vc.Citys{
		var city CityOverview
		city.CityId =c.CityId
		city.Name = c.Name
		city.Country = c.Country
		city.Description = c.Description
		city.CreatorEmail = c.CreatorEmail
		city.Images = m[c.CityId]
		userProfile.VisitedCities = append(userProfile.VisitedCities,city)
	}

	//get visited places
	vp,err := GEtVisitedPlaces(pb.VisitedRequestP{
		Token:                token,
		Email:                email,

	})

	var placesid []int32
	for _,c := range vp.Places{
		placesid = append(placesid,c.PlaceId)
	}
	//visited places images
	vpimg,err := GetVisitedPlacesPhotos(pb.GetVisitedPlacesPhotosRequest{
		Email:                email,
		Token:                token,
		PlaceId:              placesid,

	})

	m1:=  make(map[int32]([]*pb.PlacePhoto))
	for _,img := range vpimg.PlacePhotos{
		m1[img.PlacePhotos[0].PlaceId] = img.PlacePhotos
	}
	for _,p := range vp.Places{
		userProfile.VisitedPlaces = append(userProfile.VisitedPlaces,PlaceOverview{
			PlaceId:     p.PlaceId,
			Name:        p.Name,
			City:        p.City,
			Country:     p.Country,
			Location:    *p.Location,
			Description: p.Description,
			Images:      m1[p.PlaceId],
		})
	}

	json.NewEncoder(w).Encode(userProfile)
}

type PlaceResponse struct {
	Place *pb.Place
}

func GetPlaceRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received: %v", "Get place")

	// parse reueat data
	token := r.Header["Token"][0]
	fmt.Print(token)
	email := r.Header["Email"][0]
	country := mux.Vars(r)["country"]
	city := mux.Vars(r)["city"]
	name := mux.Vars(r)["name"]
	var placeResponse PlaceResponse

	response, err := GetPlace(pb.GetPlaceRequestP{
		Token:        token,
		Email:        email,
		PlaceName:    name,
		PlaceCity:    city,
		PlaceCountry: country,
	})
	if err != nil {
		log.Printf("Error: %v", err)
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}

	placeResponse.Place = response.Place

	json.NewEncoder(w).Encode(placeResponse)
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
	res, err := VisitCity(pb.VisitCityRequestP{
		Token: token,
		Email: email,
		Id:    int32(i),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(res)

}
