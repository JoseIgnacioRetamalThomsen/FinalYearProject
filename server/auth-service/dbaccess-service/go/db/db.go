// Contains functions to acess the database
package db

// Use the Ziutec driver:
//  https://github.com/ziutek/mymysql.

import (
	"errors"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
)

// Configuration for connect to the database.
type Configuration struct {
	Port            string
	Coneection_type string
	MySQL_socket    string
	MySQL_user      string
	MySQL_pass      string
	MySQL_db        string
}

// The database confection.
var db mysql.Conn

// The database configuraion.
var configuration Configuration

// Name of tables and rows in mysql.
const (
	tableUsers      = "Users"
	rowID           = "Id"
	rowEmail        = "Email"
	rowPasswordHash = "PasswordHash"
	rowPasswordSalt = "PasswordSalt"
	rowIsEmail      = "IsEmail"
	tableSession    = "UserSessions"
	rowSessionKey   = "SessionKey"
)

// Use data
type user struct {
	id           int64
	email        string
	passwordHash []byte
	salt         []byte
	isEmail      bool
}

// session data
type Session struct {
	SessionKey   string
	Email        string
	LoginTime    string
	LastSeemTime string
}


// Get salt fomr session data.
func (u *user) GetSalt() []byte {
	return u.salt
}

// Get hashed password from session data.
func (u *user) GetHashedPassword() []byte {
	return u.passwordHash
}

// Get email from user.
func (u *user) GetEmail() string {
	return u.email
}

// Get id from user.
func (u *user) GetId() int64 {
	return u.id
}

// Create a new user.
func NewUser(email string, hashedPassword []byte, salt []byte, isEmail bool) *user {

	u := user{email: email, passwordHash: hashedPassword,
		salt: salt, isEmail: isEmail}
	u.id = 0
	return &u
}

// Create a new user with id.
func NewUserId(id int64, email string, hashedPassword []byte, salt []byte, isEmail bool) *user {

	u := user{id: id, email: email, passwordHash: hashedPassword,
		salt: salt, isEmail: isEmail}

	return &u
}
func NewUserEmty() *user {
	u := user{}
	return &u
}


// Create a new seassion.
func NewSeassion(sessionKey string, email string, loginTime string, lastSeemtime string) *Session {
	s := Session{SessionKey: sessionKey, Email: email, LoginTime: loginTime, LastSeemTime: lastSeemtime}
	return &s
}

// create connection to database
func SetupConnection(connectionType string, socket string, user string, pass string, database string) (bool, error) {
	configuration.Coneection_type = connectionType
	configuration.MySQL_socket = socket
	configuration.MySQL_user = user
	configuration.MySQL_pass = pass
	configuration.MySQL_db = database
	log.Printf("Connected to database: %v", configuration)
	//db = mysql.New(connectionType, "", socket, user, pass, database)

	return true, nil
}

//Should be use for create a user for first time only.
//Returns true if user is created.
//Auto id and isEmail false.
func AddUser(u user) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmtStr := fmt.Sprintf("insert into %s (%s, %s, %s ) values (?,?,?)", tableUsers, rowEmail, rowPasswordHash, rowPasswordSalt)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Run(u.email, u.passwordHash, u.salt)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

// Return user data.
func GetUser(email string) (user, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _, err := db.Query("select * from Users where Email = '%s'", email)
	if err != nil {
		return *NewUserEmty(), err
	}
	if len(rows) <= 0 {
		return *NewUserEmty(), err
	}

	u := NewUserId(rows[0].Int64(0), rows[0].Str(1), rows[0][2].([]byte), rows[0][3].([]byte), rows[0].Bool(4))
	return *u, nil
}

// Used when the user confirm the email.
func ConfirmEmail(email string) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("UPDATE Users SET  IsEmail = true WHERE email=?")
	//checkError(err)
	if err != nil {

		return -1, errors.New(("can't confirm"))

	}
	_, res, err := stmt.Exec(email)
	if err != nil {

		return -1, errors.New(("can't confirm"))

	}
	res = res
	return int64(res.InsertId()), nil
}

// Update the user account.
func UpdateUser(u user) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmtStr := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?  WHERE %s = ?", tableUsers, rowPasswordHash, rowPasswordSalt, rowIsEmail, rowEmail)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	_, res, err := stmt.Exec(u.passwordHash, u.salt, u.isEmail, u.email)
	if err != nil {
		return -1, err
	}
	return int64(res.InsertId()), nil
}

// If user email exist all data of that user will be wiped out.
func DelUser(email string) int64 {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	del, err := db.Prepare("DELETE FROM users WHERE email=?")
	_, res, err := del.Exec(email) // OK
	if err != nil {
		return -1
	}

	return int64(res.InsertId())

}

// Create a new session using actual time.
func CreateSession(key string, email string) (string, string, string, string, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmtStr := fmt.Sprintf("insert into %s (%s, %s ,LoginTime ,LastSeenTime) values (?,?, CURTIME(),CURTIME())", tableSession, rowSessionKey, rowEmail)
	stmt, err := db.Prepare(stmtStr)
	//checkError(err)
	if err != nil {

		return "", "", "", "", err

	}
	_, res, err := stmt.Exec(key, email)
	if err != nil {
		return "", "", "", "", err
	}
	is, sess, err := GetSession(key)
	res.InsertId()

	if is {
		return sess.SessionKey, sess.Email, sess.LoginTime, sess.LastSeemTime, nil
	}
	return "", "", "", "", err //

}

// Returns a session using the key.
func GetSession(key string) (bool, Session, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmtStr := fmt.Sprintf("select * from %s where %s = '%s'", tableSession, rowSessionKey, key)
	rows, res, err := db.Query(stmtStr)
	if err != nil {
		return false, Session{}, err
	}
	res = res

	if len(rows) <= 0 {
		return false, Session{}, nil
	}
	s := NewSeassion(rows[0].Str(0), rows[0].Str(1), rows[0].Str(2), rows[0].Str(3))
	return true, *s, nil
}

// Delete a session.
func DeleteSession(key string) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableSession, rowSessionKey)
	del, err := db.Prepare(stmtStr)
	if err != nil {
		return -1, err
	}
	_, res, err := del.Exec(key)
	if err != nil {
		return -1, err
	}

	return int64(res.InsertId()), nil
}

// Delete all sessions for a user.
// Used when logout from everywhere.
func DeleteAllSession(email string) (int64, error) {
	db = mysql.New(configuration.Coneection_type, "", configuration.MySQL_socket, configuration.MySQL_user, configuration.MySQL_pass, configuration.MySQL_db)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableSession, rowEmail)
	del, err := db.Prepare(stmtStr)
	if err != nil {

	}
	_, res, err := del.Exec(email)
	if err != nil {
		return -1, err
	}

	return int64(res.InsertId()), nil

}
