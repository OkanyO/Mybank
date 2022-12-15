package users

import (
	"nuel/go-bank/interfaces"
	"reflect"
	"testing"
)

func Test_prepareToken(t *testing.T) {
	type args struct {
		user *interfaces.User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareToken(tt.args.user); got != tt.want {
				t.Errorf("prepareToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareResponse(t *testing.T) {
	type args struct {
		user      *interfaces.User
		accounts  []interfaces.ResponseAccount
		withToken bool
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareResponse(tt.args.user, tt.args.accounts, tt.args.withToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		username string
		pass     string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(tt.args.username, tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		username string
		email    string
		pass     string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Register(tt.args.username, tt.args.email, tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		id  string
		jwt string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUser(tt.args.id, tt.args.jwt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
