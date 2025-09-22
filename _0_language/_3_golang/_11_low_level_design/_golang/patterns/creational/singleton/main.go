package main

import (
	"fmt"
	"sync"
)

type config struct { value string }

var (
	instance *config
	once sync.Once
)

func GetConfig() *config {
	once.Do(func() {
		instance = &config{value: "default"}
	})
	return instance
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	fmt.Println(c1 == c2, c1.value)
}
