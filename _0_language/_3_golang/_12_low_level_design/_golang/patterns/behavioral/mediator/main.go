package main

import "fmt"

type ChatRoom interface { Send(user, msg string) }

type SimpleChat struct{}
func (SimpleChat) Send(user, msg string) { fmt.Println(user+":", msg) }

type User struct { name string; chat ChatRoom }
func (u User) Send(msg string) { u.chat.Send(u.name, msg) }

func main() {
	chat := SimpleChat{}
	alice := User{name: "Alice", chat: chat}
	bob := User{name: "Bob", chat: chat}
	alice.Send("Hi")
	bob.Send("Hello")
}
