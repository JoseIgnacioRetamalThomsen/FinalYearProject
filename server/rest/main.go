package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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

	fmt.Println(city)

	response,err := CreateCity(city)

	json.NewEncoder(w).Encode(response)

}

func GetCityRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get city")

	name := mux.Vars(r)["name"]
	country := mux.Vars(r)["country"]
	token := r.Header["Token"][0]
	email := r.Header["Email"][0]
	fmt.Println(token)
	fmt.Println(email)


	fmt.Println(name)
	fmt.Println(country)

	response, err := GetCity(pb.CityRequestP{
		Token:                token,
		Name:                 name,
		Country:              country,
		CreatorEmail:         email,

	})
	if err != nil {
	panic(err)}
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





	//GetCity(tokenEmail,token,"san pedro","chile")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/city", CreateCityRequest).Methods("POST")
	router.HandleFunc("/city/{name}/{country}", GetCityRequest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}


///********************* Profiles
const(
	//url = "0.0.0.0:60051"
	url="35.197.216.42:60051";
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


//prfiles methods
func CreateCity1(email string,token string,cityName string,cityCountry string,cityDescription string,lat float32,lon float32)bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateCity(ctx,&pb.CityRequestP{
		Token:                token,
		Name:                 cityName,
		Country:              cityCountry,
		CreatorEmail:         email,
		Description:          cityDescription,
		Location:             &pb.GeolocationP{Lat:lat,Lon:lon},

	})
	if err != nil{
		panic(err)
	}
	return r.Valid
}

func CreateCity(city pb.CityRequestP)(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.CreateCity(ctx,&city)
	if err != nil{
		return nil,err
	}
	return r,nil
}

func GetCity(city pb.CityRequestP )(*pb.CityResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := profSerConn.context.dbClient.GetCity(ctx,&city)

	if err != nil{
		return nil,err
	}
	//fmt.Println(r)// r is city request p
	//temp,_ := json.Marshal(r)
	//fmt.Println(string(temp))
	return r,nil
}
