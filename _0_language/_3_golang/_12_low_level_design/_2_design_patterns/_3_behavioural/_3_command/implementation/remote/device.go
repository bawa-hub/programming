package remote

type Device interface {
	On()
	Off()
	GetName() string
	GetState() string
}