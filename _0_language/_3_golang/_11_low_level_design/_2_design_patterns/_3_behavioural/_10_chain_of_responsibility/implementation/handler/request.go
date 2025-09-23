package handler


// Request interface
type Request interface {
	GetType() string
	GetData() interface{}
	GetPriority() int
	IsProcessed() bool
	SetProcessed(processed bool)
}

// Concrete Request
type ConcreteRequest struct {
	requestType string
	data        interface{}
	priority    int
	processed   bool
}

func NewConcreteRequest(requestType string, data interface{}, priority int) *ConcreteRequest {
	return &ConcreteRequest{
		requestType: requestType,
		data:        data,
		priority:    priority,
		processed:   false,
	}
}

func (cr *ConcreteRequest) GetType() string {
	return cr.requestType
}

func (cr *ConcreteRequest) GetData() interface{} {
	return cr.data
}

func (cr *ConcreteRequest) GetPriority() int {
	return cr.priority
}

func (cr *ConcreteRequest) IsProcessed() bool {
	return cr.processed
}

func (cr *ConcreteRequest) SetProcessed(processed bool) {
	cr.processed = processed
}