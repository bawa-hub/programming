package validation

import "strings"

type EmailValidationRule struct {
	AbstractValidationRule
}

func NewEmailValidationRule() *EmailValidationRule {
	return &EmailValidationRule{}
}

func (evr *EmailValidationRule) Validate(data interface{}) (bool, string) {
	email, ok := data.(string)
	if !ok {
		return false, "Data is not a string"
	}
	
	if !strings.Contains(email, "@") {
		return false, "Invalid email format"
	}
	
	return evr.AbstractValidationRule.Validate(data)
}