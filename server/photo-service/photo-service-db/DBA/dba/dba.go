//Jose I. Retamal
//GMIT 2020

// Provide access methods to mysql database

package dba

import (
	"errors"
	"strconv"

	//"errors"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
	pb "github.com/joseignacioretamalthomsen/wcity"
)

type Configuration struct {
	Port string
	Coneection_type    string
	MySQL_socket string
	MySQL_user string
	MySQL_pass string
	MySQL_db string
}

var db mysql.Conn

//database entries names
const(
	DATABASE_NAME = "Photos"
	PROFILES_TABLE = "Profile"
	CITY_TABLE ="City"
	PLACE_TABLE = "Place"
	POST_TABLE = "Post"
	PROFILE_EMAIL ="UserEmail"
	ALL_TIMESTAMP ="TimeStmp"
	ALL_URL = "Url"
	ALL_ID = "Id"
	ALL_SELECTED = "Selected"
	CITY_ID ="CityId";
	PLACE_ID ="PlaceId";
	POST_ID= "PostId"
)

var configuration Configuration
// create connection to database
func SetupConnection(connectionType string, socket string, user string, pass string, database string) (bool, error) {
	configuration.Coneection_type = connectionType
	configuration.MySQL_socket = socket
	configuration.MySQL_user = user
	configuration.MySQL_pass = pass
	configuration.MySQL_db = database
	log.Printf("Config success: %v", configuration)
	//db = mysql.New(connectionType, "", socket, user, pass, database)

	return true, nil
}

func AddProfilePhoto(photo pb.ProfilePhoto) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtStr := fmt.Sprintf("insert into %s (%s, %s, %s ) values (?,?,?)", PROFILES_TABLE, PROFILE_EMAIL, ALL_URL, ALL_SELECTED)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Run(photo.UserEmail, photo.Url, photo.Selected)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

func AddCityPhoto(photo pb.CityPhoto) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtStr := fmt.Sprintf("insert into %s (%s, %s, %s ) values (?,?,?)", CITY_TABLE, CITY_ID, ALL_URL, ALL_SELECTED)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Run(photo.CityId, photo.Url, photo.Selected)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

func AddPlacePhoto(photo pb.PlacePhoto) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtStr := fmt.Sprintf("insert into %s (%s, %s, %s ) values (?,?,?)", PLACE_TABLE, PLACE_ID, ALL_URL, ALL_SELECTED)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Run(photo.PlaceId, photo.Url, photo.Selected)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

func AddPostPhoto(photo pb.PostPhoto) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtStr := fmt.Sprintf("insert into %s (%s, %s, %s ) values (?,?,?)", POST_TABLE, POST_ID, ALL_URL, ALL_SELECTED)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Run(photo.PostId, photo.Url, photo.Selected)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

func GetProfilePhotos(email string) ([]*pb.ProfilePhoto, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _, err := db.Query("select * from %s where %s = '%s'",PROFILES_TABLE,PROFILE_EMAIL ,email)
	if err != nil {
		return nil,err
	}
	if len(rows) <=0 {
		return nil,errors.New("No photos found for that email.")
	}

	var list []*pb.ProfilePhoto

	for _, value := range rows{
		fmt.Println(value.Str(3))
		list = append(list,&pb.ProfilePhoto{
			Id:                   int32(value.Int64(0)),
			UserEmail:            value.Str(1),
			Url:                  value.Str(3),
			Timestamp:            value.Str(2),
			Selected:             value.Bool(4),
		})
	}

	return list, nil
}

func GetCityPhotos(cityId int) ([]*pb.CityPhoto, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _, err := db.Query("select * from %s where %s = '%s' ",CITY_TABLE,CITY_ID ,strconv.Itoa(cityId))
	fmt.Printf("select * from %s where %s = '%s'",CITY_TABLE,CITY_ID ,strconv.Itoa(cityId))
	if err != nil {
		return nil,err
	}
	if len(rows) <=0 {
		return nil,errors.New("No photos found for that id.")
	}

	var list []*pb.CityPhoto

	for _, value := range rows{
		fmt.Println(value.Str(3))
		list = append(list,&pb.CityPhoto{
			Id:                   int32(value.Int64(0)),
			CityId:               int32(value.Int64(1)),
			Url:                  value.Str(3),
			Timestamp:            value.Str(2),
			Selected:             value.Bool(4),
		})
	}

	return list, nil
}

func GetPlacePhotos(placeid int) ([]*pb.PlacePhoto, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _, err := db.Query("select * from %s where %s = '%s' ",PLACE_TABLE,PLACE_ID ,strconv.Itoa(placeid))

	if err != nil {
		return nil,err
	}
	if len(rows) <=0 {
		return nil,errors.New("No photos found for that id.")
	}

	var list []*pb.PlacePhoto

	for _, value := range rows{
		fmt.Println(value.Str(3))
		list = append(list,&pb.PlacePhoto{
			Id:                   int32(value.Int64(0)),
			PlaceId:               int32(value.Int64(1)),
			Url:                  value.Str(3),
			Timestamp:            value.Str(2),
			Selected:             value.Bool(4),
		})
	}
fmt.Println(list)
	return list, nil
}

func GetPostPhotos(postid string) ([]*pb.PostPhoto, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket , configuration.MySQL_user, configuration.MySQL_pass , configuration.MySQL_db )
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _, err := db.Query("select * from %s where %s = '%s'",POST_TABLE,POST_ID ,postid)
	if err != nil {
		return nil,err
	}
	if len(rows) <=0 {
		return nil,errors.New("No photos found for that email.")
	}

	var list []*pb.PostPhoto

	for _, value := range rows{
		fmt.Println(value.Str(3))
		list = append(list,&pb.PostPhoto{
			Id:                   int32(value.Int64(0)),
			PostId:            value.Str(1),
			Url:                  value.Str(3),
			Timestamp:            value.Str(2),
			Selected:             value.Bool(4),
		})
	}

	fmt.Println(list)
	return list, nil
}
