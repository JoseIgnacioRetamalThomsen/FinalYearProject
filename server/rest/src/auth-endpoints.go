package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"io/ioutil"
	"log"
	"net/http"

)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Create new User")


	//create user request for call services
	var userRequest  pb.UserRequest;

	//parse form
	r.ParseForm()
	// read data from form
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain request data")
	}
	json.Unmarshal(reqBody, &userRequest)

	err, response := createUser(userRequest)

	if err!=nil{
		http.Error(w, "Wrong request: "+err.Error(), 500)
	}
	//w.Header().Add("Access-Control-Allow-Headers" ,"X-Requested-With,Content-Type,Accept, Accept-Language, Content-Language, Origin")
	//w.Header().Add("Access-Control-Allow-Headers" ,"*")
	//w.Header().Add("Access-Control-Allow-Headers" ,"Content-Type")
	//w.Header().Add("Access-Control-Allow-Methods","GET,POST,OPTIONS")
	//w.Header().Add("Access-Control-Allow-Origin","*")
	// send response
	json.NewEncoder(w).Encode(response)
}



func UptadeUserEndPoint(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Get User")

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	fmt.Print(token);
	fmt.Print(email);

	var updateRequest pb.UserRequest

	//parse form
	r.ParseForm()
	// read data from form
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain request data")
	}
	json.Unmarshal(reqBody, &updateRequest)

	updateRequest.Email = email

	err, res := updateUser(updateRequest);

	if err!=nil{
		http.Error(w, "Wrong request: "+err.Error(), 500)
	}

	// send response
	json.NewEncoder(w).Encode(res)

}

func LoginEndPoint(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Login User")

	var request pb.UserRequest

	//parse form
	r.ParseForm()
	// read data from form
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain request data")
	}
	json.Unmarshal(reqBody, &request)

fmt.Print(request)

	err, res := loginUser(request)
	if err!=nil{
		http.Error(w, "Wrong request: "+err.Error(), 500)
	}

	// send response
	json.NewEncoder(w).Encode(res)

}

func LogoutEndPoint(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received: %v", "Log out")

	token := r.Header["Token"][0]
	email := r.Header["Email"][0]

	var request = pb.LogRequest{
		Token:                token,
		Email:                email,

	}

	err, res := logoutUser(request);
	if err!=nil{
		http.Error(w, "Wrong request: "+err.Error(), 500)
	}

	// send response
	json.NewEncoder(w).Encode(res)
}
