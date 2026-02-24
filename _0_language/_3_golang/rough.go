package main

import "fmt"


func main() {
	s := "vikram"
     
	for i:=0;i<len(s);i++ {
		fmt.Println(s[i])
	}

	for _,ch := range s {
		fmt.Println("ch : ", ch)
	}
}