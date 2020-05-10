package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// configuration
type Configuration struct {
	Port string
	Dbs  []string
	Auth [] string
}
// Server configuration.
var configuration Configuration

// reads config file.
func readConfig(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
