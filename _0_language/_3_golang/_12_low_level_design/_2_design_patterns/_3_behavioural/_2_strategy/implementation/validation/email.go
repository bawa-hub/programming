package validation

import "strings"

type EmailValidationStrategy struct{}

func (evs *EmailValidationStrategy) Validate(data string) (bool, string) {
	if strings.Contains(data, "@") && strings.Contains(data, ".") {
		return true, "Valid email format"
	}
	return false, "Invalid email format"
}

func (evs *EmailValidationStrategy) GetName() string {
	return "Email Validation"
}