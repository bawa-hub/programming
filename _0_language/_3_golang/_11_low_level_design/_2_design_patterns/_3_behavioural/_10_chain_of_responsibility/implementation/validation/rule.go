package validation

type ValidationRule interface {
	Validate(data interface{}) (bool, string)
	SetNext(rule ValidationRule)
}

type AbstractValidationRule struct {
	next ValidationRule
}

func (avr *AbstractValidationRule) SetNext(rule ValidationRule) {
	avr.next = rule
}

func (avr *AbstractValidationRule) Validate(data interface{}) (bool, string) {
	if avr.next != nil {
		return avr.next.Validate(data)
	}
	return true, ""
}