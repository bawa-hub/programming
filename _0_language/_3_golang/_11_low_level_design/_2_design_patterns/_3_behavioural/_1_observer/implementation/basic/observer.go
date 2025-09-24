package basic

type Observer interface {
	Update(data interface{})
	GetID() string
}
