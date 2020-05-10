package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)



func CreateCityPost(city *CityPost)(interface{},error){
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
		fmt.Println("here")
		return nil,err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID);
	ok =ok;
	return oid.Hex(),nil
}

func GetCityPost(IndexId int32)[]*pb.CityPostPSDB{
	var posts []*pb.CityPostPSDB
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

	cur, err := collection.Find(ctx, bson.M{"indexid":IndexId})
	if err != nil { log.Fatal(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {

		post := &pb.CityPostPSDB{}
		err := cur.Decode(&post)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		var post1 bson.M
		err = cur.Decode(&post1)


		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}

func CreatePlacePost(place *PlacePost)(interface{},error){
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

	data, err := bson.Marshal(place)
	if err != nil {
		panic(err)
	}


	res, err := collection.InsertOne(ctx,data)
	if err!=nil{
		return "",err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID);
	ok =ok;
	return  oid.Hex(),nil
}

func GetPlacePost(indexId int32)[]*pb.PlacePostPSDB{
	var posts []*pb.PlacePostPSDB
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

	cur, err := collection.Find(ctx, bson.M{"indexid":indexId})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {

		post := &pb.PlacePostPSDB{}
		err := cur.Decode(&post)
		if err != nil { log.Fatal(err) }

		var post1 bson.M
		err = cur.Decode(&post1)
		fmt.Println(post1["_id"])

		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return posts
}

func UpdatePost(objectId string,title string, description string)bool {

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDBURI))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	defer ctx.Done()
	if err != nil {
		panic(err)
	}

	collection := client.Database(DatabaseName).Collection(CollectionName)

	id, _ := primitive.ObjectIDFromHex(objectId)
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"title", title}, {"description", description}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount == 1 {
		return true;
	}
	return false;
}

