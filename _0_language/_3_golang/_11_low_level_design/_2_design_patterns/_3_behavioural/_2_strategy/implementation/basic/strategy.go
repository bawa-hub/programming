package basic

// Strategy interface
type Strategy interface {
	Execute(data interface{}) interface{}
	GetName() string
}