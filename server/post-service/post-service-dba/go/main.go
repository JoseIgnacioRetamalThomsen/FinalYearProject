package main

import (
	"context"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"time"
)

const(
	//MongoDBURI = "mongodb://172.17.0.1:27017"

	MongoDBURI = "mongodb://127.0.0.1:27017"
	DatabaseName ="PostDatabase"
	CollectionName  = "Posts"
	Port = ":2787"
)

type server struct {
	pb.UnimplementedPostsServiceDBAServer
}

type CityPost struct{
	IndexId int32
	CreatorEmail string
	Name string
	Country string
	Title string
	Body string
	TimeStamp string
	Likes[] string
	MongoId string `bson:"_id,omitempty"`

}
type PlacePost struct{
	IndexId int32
	CreatorEmail string
	Name string
	City string
	Country string
	Title string
	Body string
	TimeStamp string
	Likes[] string
	MongoId string `bson:"_id,omitempty"`
}

/**
*  END POINTS
 */
func (s *server) CreateCityPost(ctx context.Context, in *pb.CityPostPSDB) (*pb.CreatePostResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Create city post", in.String())

	index,_ :=CreateCityPost(&CityPost{
		IndexId:      in.IndexId,
		CreatorEmail: in.CreatorEmail,
		Name:         in.CityName,
		Country:      in.CityCountry,
		Title:        in.Title,
		Body:         in.Body,
		TimeStamp:    in.TimeStamp,
		Likes:        nil,
		MongoId:      "",
	})
	index = index
	return &pb.CreatePostResponsePSDB{
		Valied:               true,
		IndexId:              in.IndexId,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	},nil
}

func (s *server) CreatePlacePost(ctx context.Context, in *pb.PlacePostPSDB) (*pb.CreatePostResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Create place post", in.String())

	index, err := CreatePlacePost(&PlacePost{
		IndexId:      in.IndexId,
		CreatorEmail: in.CreatorEmail,
		Name:         in.PlaceName,
		City:         in.CityName,
		Country:      in.CountryName,
		Title:        in.Title,
		Body:         in.Body,
		TimeStamp:    in.TimeStamp,
		Likes:        nil,

	})
	index = index
	if err != nil{
		return &pb.CreatePostResponsePSDB{
			Valied:               false,
			IndexId:              in.IndexId,

		},err
	}

	return &pb.CreatePostResponsePSDB{
		Valied:               true,
		IndexId:              in.IndexId,

	},nil
}

func (s *server) GetPlacePosts(ctx context.Context, in *pb.PostsRequestPSDB) (*pb.PlacePostsResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Get place posts", in.String())

	posts := GetPlacePost(in.IndexId)

	if posts == nil{
		return &pb.PlacePostsResponsePSDB{
			Valid:                false,
			IndexId:              0,
			Posts:                nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},status.Error(codes.NotFound,"No post found.")
	}

return &pb.PlacePostsResponsePSDB{
	Valid:                true,
	IndexId:              in.IndexId,
	Posts:                posts,
	XXX_NoUnkeyedLiteral: struct{}{},
	XXX_unrecognized:     nil,
	XXX_sizecache:        0,
},nil

}
func (s *server) GetCityPosts(ctx context.Context, in *pb.PostsRequestPSDB) (*pb.CityPostsResponsePSDB, error) {
	log.Printf("Received: %v , from: %v", "Get city posts", in.String())

	posts := GetCityPost(in.IndexId)

	if posts == nil{
		return &pb.CityPostsResponsePSDB{
			Valid:                false,
			IndexId:              in.IndexId,
			Posts:                nil,

		},status.Error(codes.NotFound,"No post found.")

	}
	return &pb.CityPostsResponsePSDB{
		Valid:                true,
		IndexId:              in.IndexId,
		Posts:                posts,

	},nil
}
func main(){
/*
	temp := &CityPost{
		IndexId: 1,
		CreatorEmail: "one",
		Name:         "two",
		Country:      "three",
		Title:        "four",
		Body:         "five",
		TimeStamp:    "six",
		Likes:        []string{"one","two"},
	}
	instid , err :=CreateCityPost(temp)
	if err!= nil{
		panic(err)
	}

	fmt.Println(instid)

	temp1 := &PlacePost{
		IndexId:      2,
		CreatorEmail: "a",
		Name:         "bv",
		City:         "c",
		Country:      "d",
		Title:        "a",
		Body:         "f",
		TimeStamp:    "d",
		Likes:        []string{"a","b","c"},
	}

	instid , err =	CreatePlacePost(temp1)
	if err!= nil{
		panic(err)
	}

	fmt.Println(instid)


 */
/*
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
	//fmt.Println(collection)



	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)



/*
	res, err := collection.InsertOne(ctx, bson.M{})
	if err!=nil{
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)

*//*
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)


	cur, err := collection.Find(ctx, bson.M{"cityid":1})
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
*/
//fmt.Println(GetCityPost(1))
//fmt.Println(GetPlacePost(2))
//UpdatePost("5e4d9e2d6c8b79466543aa4b","assssssssssssssssssssssssssssss","ssssssssffffffffffffffffffffffffffffffffffff")

	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServiceDBAServer(s, &server{})
	log.Print("Service Started in port: ", Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}


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
		return nil,err
	}
	return res.InsertedID,nil
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
	return res.InsertedID,nil
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

