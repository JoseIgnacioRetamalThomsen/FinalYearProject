package main

import (
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Create User")
}
