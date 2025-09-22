package main

import "fmt"

type OrderStore interface{ Save(order string) }

type MySQL struct{}
func (MySQL) Save(order string) { fmt.Println("MySQL save:", order) }

type Postgres struct{}
func (Postgres) Save(order string) { fmt.Println("Postgres save:", order) }

type OrderService struct{ store OrderStore }

func NewOrderService(store OrderStore) OrderService { return OrderService{store: store} }

func (s OrderService) Place(order string) { s.store.Save(order) }

func main() {
	service := NewOrderService(MySQL{})
	service.Place("order-2")

	service = NewOrderService(Postgres{})
	service.Place("order-3")
}
