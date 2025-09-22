package main

import "fmt"

type Authenticator interface {
	Authenticate(user, pass string) bool
}

type BasicAuth struct{}
func (BasicAuth) Authenticate(user, pass string) bool { return user == "admin" && pass == "secret" }

type OAuth struct{}
func (OAuth) Authenticate(user, pass string) bool { return user != "" }

func Login(a Authenticator, user, pass string) bool { return a.Authenticate(user, pass) }

func main() {
	fmt.Println(Login(BasicAuth{}, "admin", "secret"))
	fmt.Println(Login(OAuth{}, "alice", "token"))
}
