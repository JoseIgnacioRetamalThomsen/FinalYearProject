package dba

import (
	//"errors"
	//"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
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

//database names

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

