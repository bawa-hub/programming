package remote

type DeviceCommand interface {
	Execute()
	Undo()
	GetName() string
}