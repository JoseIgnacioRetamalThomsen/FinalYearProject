package main

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

