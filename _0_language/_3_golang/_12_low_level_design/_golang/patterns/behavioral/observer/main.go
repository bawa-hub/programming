package main

import "fmt"

type Observer interface { Update(data string) }

type Subject interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(data string)
}

type News struct{ subs []Observer }
func (n *News) Attach(o Observer) { n.subs = append(n.subs, o) }
func (n *News) Detach(o Observer) {
	for i, s := range n.subs { if s == o { n.subs = append(n.subs[:i], n.subs[i+1:]...); break } }
}
func (n *News) Notify(data string) { for _, s := range n.subs { s.Update(data) } }

type User struct{ name string }
func (u User) Update(data string) { fmt.Println(u.name, "received:", data) }

func main() {
	n := &News{}
	a := User{name: "Alice"}
	b := User{name: "Bob"}
	n.Attach(a); n.Attach(b)
	n.Notify("Breaking")
}
