package main
// root pass = oass
import (
	
	"fmt"
	 "github.com/ziutek/mymysql/mysql"
	_"github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)
 	

type user struct {
	id int
	email string
    username string
	hashedPassword  []byte
	salt []byte
}


func NewUser(email string,username string, hashedPassword []byte, salt []byte) *user {

		
	
		u := user{email : email,username: username,hashedPassword: hashedPassword,
		salt:salt}
		u.id = 0
		return &u
	}
	


func main() {

	//db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "", "test")


	fmt.Println("working")
  u := NewUser("ema8il@gmail.com","name",[]byte("Here is a string...."),[]byte("Here is a string...."))
  //AddUser(*u)
  GetPassSalt(u)
  fmt.Println(u.salt)
}

func AddUser(u user){
//GRANT ALL ON users.* TO 'test'@'localhost';
	db := mysql.New("tcp", "", "127.0.0.1:3306", "test", "newpassword", "user_login")	
	err := db.Connect()
	if err != nil {
		panic(err)
		
	}

	stmt, err := db.Prepare("insert into users (email, username, hashedpassword, salt) values(?,?,?,?)")
	//checkError(err)
	if err != nil {
		panic(err)
		
	}
	//fmt.Println(u.email)	


	_, err = stmt.Run(u.email,u.username,u.hashedPassword,u.salt)
	if err != nil {
		panic(err)
		
	}
	defer db.Close()


}

func GetPassSalt(u *user){
	db := mysql.New("tcp", "", "127.0.0.1:3306", "test", "newpassword", "user_login")	
	err := db.Connect()
	if err != nil {
		panic(err)
		
	}
	//stmt, err := db.Prepare("select hashedPassword, salt from users where email = ?")
	//res, err := db.Start("select hashedPassword, salt from users where email = ?")
	if err != nil {
		panic(err)
		
	}

	//_, err = stmt.Run(u.email)
	//res, err := db.Start("select * from X")
	rows, res, err := db.Query("select hashedPassword, salt from users where email = '%s'", u.email)
	if err != nil {
		panic(err)
		
	}
	rows = rows
	res = res
	fmt.Println(res)
	//u.salt = []byte("12")

}