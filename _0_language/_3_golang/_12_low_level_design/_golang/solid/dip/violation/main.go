package main

import "fmt"

type MySQL struct{}

func (MySQL) Save(order string) { fmt.Println("MySQL save:", order) }

// High-level depends on concrete MySQL => DIP violation
// Changing DB requires modifying OrderService.
type OrderService struct{ db MySQL }

func (s OrderService) Place(order string) { s.db.Save(order) }

func main() {
	service := OrderService{db: MySQL{}}
	service.Place("order-1")
}
