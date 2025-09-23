package validation

import "fmt"

type LengthValidationRule struct {
	AbstractValidationRule
	minLength int
	maxLength int
}

func NewLengthValidationRule(minLength, maxLength int) *LengthValidationRule {
	return &LengthValidationRule{
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (lvr *LengthValidationRule) Validate(data interface{}) (bool, string) {
	text, ok := data.(string)
	if !ok {
		return false, "Data is not a string"
	}
	
	if len(text) < lvr.minLength {
		return false, fmt.Sprintf("Text too short (minimum %d characters)", lvr.minLength)
	}
	
	if len(text) > lvr.maxLength {
		return false, fmt.Sprintf("Text too long (maximum %d characters)", lvr.maxLength)
	}
	
	return lvr.AbstractValidationRule.Validate(data)
}
