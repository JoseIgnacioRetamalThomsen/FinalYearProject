package db

import (
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	type args struct {
		u user
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfirmEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfirmEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConfirmEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSession(t *testing.T) {
	type args struct {
		key   string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		want3   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := CreateSession(tt.args.key, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSession() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateSession() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("CreateSession() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("CreateSession() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestDelUser(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelUser(tt.args.email); got != tt.want {
				t.Errorf("DelUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteAllSession(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteAllSession(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAllSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteAllSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteSession(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSession(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetSession(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSession() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetSession() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    user
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSeassion(t *testing.T) {
	type args struct {
		sessionKey   string
		email        string
		loginTime    string
		lastSeemtime string
	}
	tests := []struct {
		name string
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeassion(tt.args.sessionKey, tt.args.email, tt.args.loginTime, tt.args.lastSeemtime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeassion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		email          string
		hashedPassword []byte
		salt           []byte
		isEmail        bool
	}
	tests := []struct {
		name string
		args args
		want *user
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.email, tt.args.hashedPassword, tt.args.salt, tt.args.isEmail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserEmty(t *testing.T) {
	tests := []struct {
		name string
		want *user
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserEmty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserEmty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserId(t *testing.T) {
	type args struct {
		id             int64
		email          string
		hashedPassword []byte
		salt           []byte
		isEmail        bool
	}
	tests := []struct {
		name string
		args args
		want *user
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserId(tt.args.id, tt.args.email, tt.args.hashedPassword, tt.args.salt, tt.args.isEmail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetupConnection(t *testing.T) {
	type args struct {
		connectionType string
		socket         string
		user           string
		pass           string
		database       string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetupConnection(tt.args.connectionType, tt.args.socket, tt.args.user, tt.args.pass, tt.args.database)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetupConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		u user
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_GetEmail(t *testing.T) {
	type fields struct {
		id           int64
		email        string
		passwordHash []byte
		salt         []byte
		isEmail      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:           tt.fields.id,
				email:        tt.fields.email,
				passwordHash: tt.fields.passwordHash,
				salt:         tt.fields.salt,
				isEmail:      tt.fields.isEmail,
			}
			if got := u.GetEmail(); got != tt.want {
				t.Errorf("GetEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_GetHashedPassword(t *testing.T) {
	type fields struct {
		id           int64
		email        string
		passwordHash []byte
		salt         []byte
		isEmail      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:           tt.fields.id,
				email:        tt.fields.email,
				passwordHash: tt.fields.passwordHash,
				salt:         tt.fields.salt,
				isEmail:      tt.fields.isEmail,
			}
			if got := u.GetHashedPassword(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHashedPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_GetId(t *testing.T) {
	type fields struct {
		id           int64
		email        string
		passwordHash []byte
		salt         []byte
		isEmail      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:           tt.fields.id,
				email:        tt.fields.email,
				passwordHash: tt.fields.passwordHash,
				salt:         tt.fields.salt,
				isEmail:      tt.fields.isEmail,
			}
			if got := u.GetId(); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_GetSalt(t *testing.T) {
	type fields struct {
		id           int64
		email        string
		passwordHash []byte
		salt         []byte
		isEmail      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:           tt.fields.id,
				email:        tt.fields.email,
				passwordHash: tt.fields.passwordHash,
				salt:         tt.fields.salt,
				isEmail:      tt.fields.isEmail,
			}
			if got := u.GetSalt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
