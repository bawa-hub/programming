package validation

type ValidationStrategy interface {
	Validate(data string) (bool, string)
	GetName() string
}