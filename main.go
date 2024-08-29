package main

import (
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

type myCredentialStore struct {
	user     string
	password string
}

func (cs *myCredentialStore) Valid(user, password, userAddr string) bool {
	return user == cs.user && password == cs.password
}

func main() {

	auth := socks5.UserPassAuthenticator{
		Credentials: &myCredentialStore{user: "login", password: "password"},
	}

	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithAuthMethods([]socks5.Authenticator{auth}),
	)

	if err := server.ListenAndServe("tcp", ":10800"); err != nil {
		panic(err)
	}
}
