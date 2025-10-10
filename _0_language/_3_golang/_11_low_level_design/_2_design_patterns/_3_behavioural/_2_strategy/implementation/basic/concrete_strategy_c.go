package basic

import "fmt"

type ConcreteStrategyC struct{}

func (csc *ConcreteStrategyC) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy C processed: %v", data)
}

func (csc *ConcreteStrategyC) GetName() string {
	return "Strategy C"
}