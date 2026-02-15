package validation

import "strconv"


type PhoneValidationStrategy struct{}

func (pvs *PhoneValidationStrategy) Validate(data string) (bool, string) {
	// Simple phone validation (10 digits)
	if len(data) == 10 {
		if _, err := strconv.Atoi(data); err == nil {
			return true, "Valid phone number"
		}
	}
	return false, "Invalid phone number format"
}

func (pvs *PhoneValidationStrategy) GetName() string {
	return "Phone Validation"
}