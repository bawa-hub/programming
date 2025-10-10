package basic

import "fmt"

type ConcreteStrategyB struct{}

func (csb *ConcreteStrategyB) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy B processed: %v", data)
}

func (csb *ConcreteStrategyB) GetName() string {
	return "Strategy B"
}