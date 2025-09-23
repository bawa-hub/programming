package handler


// Handler interface
type Handler interface {
	Handle(request Request) bool
	SetNext(handler Handler)
	CanHandle(request Request) bool
}

// Abstract Handler
type AbstractHandler struct {
	next Handler
}

func (ah *AbstractHandler) SetNext(handler Handler) {
	ah.next = handler
}

func (ah *AbstractHandler) Handle(request Request) bool {
	if ah.next != nil {
		return ah.next.Handle(request)
	}
	return false
}