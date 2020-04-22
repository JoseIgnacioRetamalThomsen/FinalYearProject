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

	r.ParseForm()

	var userRequest  pb.UserRequest;

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain request data")
	}
	json.Unmarshal(reqBody, &userRequest)

	fmt.Print(userRequest)

	err, response := createUser(userRequest)

	if err!=nil{
		http.Error(w, "Wrong request: "+err.Error(), 500)
	}

	json.NewEncoder(w).Encode(response)
}
