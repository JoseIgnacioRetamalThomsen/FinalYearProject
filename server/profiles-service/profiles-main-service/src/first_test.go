package main

import (
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println(GetUser(pb.GetUserRequestPDB{
		Email:                "mm",

	}))
}
