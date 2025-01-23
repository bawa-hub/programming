package main

import "fmt"

func main() {
  var shortGolang = "Watch Go crash course"
  var fullGolang = "Watch Nana's Golang Full Course"
  var rewardDessert = "Reward myself with a donut"

  var taskItems = []string{shortGolang, fullGolang, rewardDessert}

  fmt.Println("###### Welcome to our Todolist App! ######")
  fmt.Println("List of my Todos")

  for index, task := range taskItems {
    fmt.Printf("%d: %s\n", index+1, task)
  }
}
