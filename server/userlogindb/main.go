package main
// root pass = oass
import (
	//"os"
//	"bytes"
	"fmt"
	 "github.com/ziutek/mymysql/mysql"
	_"github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)
 	
const Coneection_type = "tcp"
const MySQL_socket = "127.0.0.1:3306"
const MySQL_user = "test"
const MySQL_pass = "newpassword"
const MySQL_db = "user_login"

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
	

func NewUserId(id int,email string,username string, hashedPassword []byte, salt []byte) *user {

		
	
	u := user{id : id,email : email,username: username,hashedPassword: hashedPassword,
	salt:salt}
	
	return &u
}

	
func main() {

	//db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "", "test")


	fmt.Println("working")
  //u := NewUser("ema8il@gmail.com","name",[]byte("Here is a string...."),[]byte("Here is a string...."))
  //AddUser(*u)
  //GetPassSalt(u)
 // fmt.Println(GetUser("email9").email)
 DelUser("ema8il@gmail.com")
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

func DelUser(email string) {
	db := mysql.New(Coneection_type, "", MySQL_socket, MySQL_user, MySQL_pass, MySQL_db)	
	err := db.Connect()
	if err != nil {
		panic(err)
		
	}

	del,err := db.Prepare("DELETE FROM users WHERE email=?")
	_,res, err := del.Exec(email)  // OK
	if err != nil {
		panic(err)
		
	}
	res=res
	
}

func GetUser(email string) user{
	db := mysql.New(Coneection_type, "", MySQL_socket, MySQL_user, MySQL_pass, MySQL_db)	
	err := db.Connect()
	if err != nil {
		panic(err)
		
	}
	rows, res,err := db.Query("select * from users where email = '%s'", email)
	if err != nil {
		panic(err)
		
	}
	res=res
	/*
	fmt.Print(rows[0].Int(0))
	fmt.Print(rows[0].Str(1))
	fmt.Print(rows[0].Str(2))
	var t = rows[0][3].([]byte)
	fmt.Print(t)
	fmt.Print(bytes.NewBuffer(t).String())
	var t1 = rows[0][4].([]byte)
	fmt.Print(t1)
	fmt.Print(bytes.NewBuffer(t1).String())
	res=res
//	fmt.Print(res)
	fmt.Println()
	fmt.Println()
	*/
	uu := NewUserId(rows[0].Int(0),rows[0].Str(1),rows[0].Str(2),rows[0][3].([]byte),rows[0][4].([]byte))
//	var  u user 
/*
for rows.Next(){
		err = rows.Scan(&u.id, &u.email,&u.username,&u.hashedPassword,&u.salt)
		if err != nil {
			panic(err)
			
		}
}
*/
		/*
	for{
	   row,err := res.GetRow()
	   if err != nil {
		panic(err)
		
	}
	if row ==nil{
		break
	}
	for _,col:= range row{
		if col ==nil {
			fmt.Print("<NULL>")	fmt.Println()
		fmt.Print(" ")
	}
		fmt.Println()
	}
*/

	return *uu
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

