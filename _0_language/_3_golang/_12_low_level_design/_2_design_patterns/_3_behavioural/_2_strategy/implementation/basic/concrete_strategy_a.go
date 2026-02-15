package basic

import "fmt"

type ConcreteStrategyA struct{}

func (csa *ConcreteStrategyA) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy A processed: %v", data)
}

func (csa *ConcreteStrategyA) GetName() string {
	return "Strategy A"
}