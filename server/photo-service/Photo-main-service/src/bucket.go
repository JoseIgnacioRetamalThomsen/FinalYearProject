package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	defaultProfile = "https://storage.googleapis.com/wcity-images-1/profile-1/profile_0.jpg"
)
const (
	google_api_key = "GOOGLE_APPLICATION_CREDENTIALS"
	start_url      = "https://storage.googleapis.com/"
)

var buckets = []string{"wcity-images-1"}
var profileFolders = []string{"profile-1"}
var cityFolder = []string{"city-1"}
var placeFoler = []string{"place-1"}
var postFolder = []string{"post-1"}

func Write(image []byte, bucketName string, imageName string) {
	r := bytes.NewReader(image)
	projectID := os.Getenv(google_api_key)

	if projectID == "" {
		log.Println(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(bucketName).Object(imageName).NewWriter(ctx)
	if _, err = io.Copy(wc, r); err != nil {
		panic(err)
	}
	if err := wc.Close(); err != nil {
		panic(err)
	}
}

//https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

