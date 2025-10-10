package validation

type PasswordValidationStrategy struct{}

func (pvs *PasswordValidationStrategy) Validate(data string) (bool, string) {
	if len(data) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range data {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case char >= 33 && char <= 126:
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return false, "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter"
	}
	if !hasDigit {
		return false, "Password must contain at least one digit"
	}
	if !hasSpecial {
		return false, "Password must contain at least one special character"
	}
	
	return true, "Valid password"
}

func (pvs *PasswordValidationStrategy) GetName() string {
	return "Password Validation"
}