package db

import (
	"errors"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

const Coneection_type = "tcp"
const MySQL_socket = "127.0.0.1:3306"
const MySQL_user = "test"
const MySQL_pass = "newpassword"
const MySQL_db = "user_login"

type user struct {
	id             int
	email          string
	username       string
	hashedPassword []byte
	salt           []byte
	isEmail        bool
}

func (u *user) GetSalt() []byte {
	return u.salt
}

func (u *user) GetHashedPassword() []byte {
	return u.hashedPassword
}

func (u *user) GetEmail() string {
	return u.email
}

func NewUser(email string, username string, hashedPassword []byte, salt []byte, isEmail bool) *user {

	u := user{email: email, username: username, hashedPassword: hashedPassword,
		salt: salt, isEmail: isEmail}
	u.id = 0
	return &u
}

func NewUserId(id int, email string, username string, hashedPassword []byte, salt []byte, isEmail bool) *user {

	u := user{id: id, email: email, username: username, hashedPassword: hashedPassword,
		salt: salt, isEmail: isEmail}

	return &u
}
func NewUserEmty() *user{
	u := user{}
	return &u
}

//Should be use for create a user for first time only.
//Returns true if user is created.
//Auto id and isEmail false.
func AddUser(u user) bool {
	created := true
	db := mysql.New(Coneection_type, "", MySQL_socket, MySQL_user, MySQL_pass, MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("insert into users (email, username, hashedpassword, salt) values(?,?,?,?)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Run(u.email, u.username, u.hashedPassword, u.salt)
	if err != nil {
		created = false
		//panic(err)}
	}
	defer db.Close()
	return created
}

func ConfirmEmail(email string) (bool,error){
	db := mysql.New("tcp", "", "127.0.0.1:3306", "test", "newpassword", "user_login")
	err := db.Connect()
	if err != nil {
		return false , errors.New(("can't connect"))
	}
	stmt, err := db.Prepare("UPDATE users SET  isEmail = true WHERE email=?")
	//checkError(err)
	if err != nil {
		return false , errors.New(("can't confirm"))

	}

	_, res, err := stmt.Exec(email)
	if err != nil {
		return false , errors.New(("can't confirm"))

	}
	res = res
	return true, nil
}

func UpdateUser(u user) (bool,error){
	db := mysql.New("tcp", "", "127.0.0.1:3306", "test", "newpassword", "user_login")
	err := db.Connect()
	if err != nil {
		return false, errors.New("Can't connect")
	}
	stmt, err := db.Prepare("UPDATE users SET username = ?, hashedpassword = ?, salt = ? , isEmail = ? WHERE email=?")
	//checkError(err)
	if err != nil {
		return false , errors.New(("can't update"))

	}

	_, res, err := stmt.Exec(u.username, u.hashedPassword, u.salt, u.isEmail, u.email)
	if err != nil {
		return false, errors.New(("can't update"))

	}
	res = res
	return true,nil
}


// If user email exist all data of that user will be wiped out.
func DelUser(email string) int64{
	db := mysql.New(Coneection_type, "", MySQL_socket, MySQL_user, MySQL_pass, MySQL_db)
	err := db.Connect()
	if err != nil {
		return -1

	}
	del, err := db.Prepare("DELETE FROM users WHERE email=?")
	_, res, err := del.Exec(email) // OK
	if err != nil {
		return -1
	}

	return int64(res.InsertId())

}

// Return user data.
func GetUser(email string) (user,error) {
	db := mysql.New(Coneection_type, "", MySQL_socket, MySQL_user, MySQL_pass, MySQL_db)
	err := db.Connect()
	if err != nil {
		return *NewUserEmty(),errors.New("Can't connect")
	}
	rows, res, err := db.Query("select * from users where email = '%s'", email)
	if err != nil {
		return *NewUserEmty(),errors.New("user don't exist")
	}
	res = res
	u := NewUserId(rows[0].Int(0), rows[0].Str(1), rows[0].Str(2), rows[0][3].([]byte), rows[0][4].([]byte), rows[0].Bool(5))
	return *u,nil
}
