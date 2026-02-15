package validation

import "strings"

type RequiredFieldValidationRule struct {
	AbstractValidationRule
}

func NewRequiredFieldValidationRule() *RequiredFieldValidationRule {
	return &RequiredFieldValidationRule{}
}

func (rfvr *RequiredFieldValidationRule) Validate(data interface{}) (bool, string) {
	if data == nil {
		return false, "Field is required"
	}
	
	text, ok := data.(string)
	if ok && strings.TrimSpace(text) == "" {
		return false, "Field cannot be empty"
	}
	
	return rfvr.AbstractValidationRule.Validate(data)
}