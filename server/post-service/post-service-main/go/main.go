package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const(
	MongoDBURI = "mongodb://172.17.0.1:27017"
	DatabaseName ="PostDatabase"
	CollectionName  = "Posts"
)

type CityPost struct{
	CreatorEmail string
	Name string
	Country string
	Title string
	Body string
	TimeStamp string
	Likes[] string

}
type PlacePost struct{
	CreatorEmail string
	Name string
	City string
	Country string
	Title string
	Body string
	TimeStamp string
	Likes[] string
}


func main(){
	temp := &CityPost{
		CreatorEmail: "one",
		Name:         "two",
		Country:      "three",
		Title:        "four",
		Body:         "five",
		TimeStamp:    "six",
		Likes:        []string{"one","two"},
	}
	_ , err :=CreateCityPost(*temp)
	if err!= nil{
		panic(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://172.17.0.1:27017"))
if err!= nil {
	panic(err)
}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	defer ctx.Done()
if err != nil{
	panic(err)
}

	collection := client.Database(DatabaseName).Collection(CollectionName)
	fmt.Println(collection)



	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)




	res, err := collection.InsertOne(ctx, bson.M{})
	if err!=nil{
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)


	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)


	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
	fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}


}


func CreateCityPost(city CityPost)(interface{},error){
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDBURI))
	if err!= nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	defer ctx.Done()
	if err != nil{
		panic(err)
	}

	collection := client.Database(DatabaseName).Collection(CollectionName)

	data, err := bson.Marshal(city)
	if err != nil {
		panic(err)
	}

	res, err := collection.InsertOne(ctx,data)
	if err!=nil{
		return nil,err
	}
	return res.InsertedID,nil
}
