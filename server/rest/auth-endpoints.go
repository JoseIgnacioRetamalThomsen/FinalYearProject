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
    //parse form
	r.ParseForm()

	//create user request for call services
	var userRequest  pb.UserRequest;

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

	// send response
	json.NewEncoder(w).Encode(response)
}
