package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joseignacioretamalthomsen/photos-dba/dba"

	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct {

}
type Configuration struct {
	Port string
	Coneection_type    string
	MySQL_socket string
	MySQL_user string
	MySQL_pass string
	MySQL_db string
}
var configuration Configuration


func readConfig(fileName string){
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	/*
	   configuration.Port = port
	   configuration.MySQL_db = MySQL_db
	   configuration.Coneection_type = Coneection_type
	   configuration.MySQL_socket = MySQL_socket
	   configuration.MySQL_user = MySQL_user
	   configuration.MySQL_pass = MySQL_pass*/
}


func main(){

	//read config file name from console input
	args := os.Args[1]
	readConfig(args)

	dba.SetupConnection(configuration.Coneection_type,
		configuration.MySQL_socket,
		configuration.MySQL_user,
		configuration.MySQL_pass,
		configuration.MySQL_db)
	log.Print("Starting Service")


}
