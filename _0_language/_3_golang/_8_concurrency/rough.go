package main

import (
	"fmt"
	"time"
)

type TODO struct {
	id int
	todo string
	created_at time.Time
}


func main() {
	todos := []TODO{}

	todo := TODO{id: 1, todo: "Learn Golang", created_at: time.Now()}
	todos = append(todos,todo)

	for i:=0;i<len(todos);i++ {
		fmt.Println(todos[i].todo, todos[i].created_at.Format(time.DateOnly))
	}
}