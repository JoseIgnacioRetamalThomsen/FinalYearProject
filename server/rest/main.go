package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	"context"


	//"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"

)




type CityResponse struct {
	Success   bool      `json:"success"`
	City pb.CityResponseP `json:"city"`
}


func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Create city")


	var city pb.CityRequestP
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}

	json.Unmarshal(reqBody, &city)

	city.Token = r.Header["Token"][0]
	city.CreatorEmail = r.Header["Email"][0]
	fmt.Println(city)

	response,err := CreateCity(city)

	json.NewEncoder(w).Encode(response)

}

func GetCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get city")

	response, err := GetCity(pb.CityRequestP{
		Token:                r.Header["Token"][0],
		Name:                 mux.Vars(r)["name"],
		Country:              mux.Vars(r)["country"],
		CreatorEmail:         r.Header["Email"][0],

	})
	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}
	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}

func UpdateCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Update city")

	var city pb.CityRequestP
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}
	json.Unmarshal(reqBody, &city)

	fmt.Println(city)


	city.Token = r.Header["Token"][0]
	city.CreatorEmail =  r.Header["Email"][0]
	city.Name =  mux.Vars(r)["name"]
    city.Country= mux.Vars(r)["country"]

    fmt.Println(city)
    response, err := UpdateCity(city)

	if err != nil {
		log.Printf("Server problem: %s", err)
		fmt.Fprintf(w, "Server problem: %s", err)
	}

	json.NewEncoder(w).Encode(response)

}

func CreatePlaceRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Create place")

}
func GetPlaceRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get place")

	response, err := GetPlace(pb.PlaceRequestP{
		Token:               r.Header["Token"][0],
		Name:                 mux.Vars(r)["name"],
		City:                 mux.Vars(r)["city"],
		Country:              mux.Vars(r)["country"],
		CreatorEmail:         r.Header["Email"][0],


	})
	if err != nil {
		log.Printf("Error: %v", err)
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}


	json.NewEncoder(w).Encode(response)
}
func UpdatePlaceRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Update place")

}

func CreateCityPostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Craete city post")

	var post pb.CityPost
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}
	json.Unmarshal(reqBody, &post)

	 post.CreatorEmail = r.Header["Email"][0]

	fmt.Println(post)

	response , err := CreateCityPost(post)
	json.NewEncoder(w).Encode(response)
}

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

func main() {

	//conect to profiles server
	dbserverCtx, err := newProfilesServiceContext(url)
	if err != nil {
		log.Fatal(err)
	}
	s2 := &profileServer{dbserverCtx}
	profSerConn = *s2

	//conect to  post server
	dbserverCtx1, err := newPostServiceContext(POST_url)
	if err != nil {
		log.Fatal(err)
	}
	s1 := &postService{dbserverCtx1}
	serviceConn = *s1



	//GetCity(tokenEmail,token,"san pedro","chile")

	//profiles
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/city", CreateCityRequest).Methods("POST")
	router.HandleFunc("/city/{country}/{name}/", GetCityRequest).Methods("GET")
	//put not working
	router.HandleFunc("/city/{country}/{name}/",UpdateCityRequest).Methods("PUT")
	router.HandleFunc("/place", CreatePlaceRequest).Methods("POST")
	router.HandleFunc("/place/{country}/{city}/{name}/", GetPlaceRequest).Methods("GET")
	router.HandleFunc("/place/{country}/{city}/{name}/",UpdatePlaceRequest).Methods("PUT")

	//post
	router.HandleFunc("/city/post", CreateCityPostRequest).Methods("POST")
	router.HandleFunc("/posts/city/{indexid}/", GetCityPostRequest).Methods("GET")


	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	//log.Fatal(http.ListenAndServe(":8080", router))
}


///********************* Profiles
const(
	url = "0.0.0.0:60051"
	//url="35.197.216.42:60051";
	//url = "35.234.146.99:5777"
	token ="a31e31a2fcdf2a9a230120ea620f3b24f7379d923fb122323d3cb9bc56fe6508"
	tokenEmail ="a@a.com"
)

type profileServer struct {
	context *profileServiceContext
}

type profileServiceContext struct {
	dbClient pb.ProfilesClient
	timeout time.Duration
}
var profSerConn profileServer

// create connection
func newProfilesServiceContext(endpoint string) (*profileServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &profileServiceContext{
		dbClient: pb.NewProfilesClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}
//profile methos
const(
	DEADLINE = 5;
	)
func CreateCity(city pb.CityRequestP)(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateCity(ctx,&city)
	if err != nil{
		return nil,err
	}
	return r,nil
}

func GetCity(city pb.CityRequestP )(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetCity(ctx,&city)

	if err != nil{
		return nil,err
	}
fmt.Println(r.Id)
	fmt.Println(r)
	return r,nil
}

func UpdateCity(city pb.CityRequestP)(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.UpdateCity(ctx,&city)
	if err != nil{
		return nil,err
	}
	return r,nil
}

func GetPlace(request pb.PlaceRequestP)(*pb.PlaceResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetPlace(ctx,&request)
	if err != nil{
		return nil,err
	}

	return r,nil
}


/*


*************************POSTTTT

 */


const (
	POST_url ="35.197.216.42:10051"
)

type postService struct {
	context *postServiceContext
}

type postServiceContext struct {
	dbClient pb.PostsServiceClient
	timeout time.Duration
}
var serviceConn postService

// create connection
func newPostServiceContext(endpoint string) (*postServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &postServiceContext{
		dbClient: pb.NewPostsServiceClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

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
