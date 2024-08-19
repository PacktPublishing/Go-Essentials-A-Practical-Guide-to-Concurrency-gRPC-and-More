package validation

import (
	"regexp"
)

func IsValidPhoneNumber(phoneNumber string) bool {
	// "123-456-7890" or "1234567890"
	phoneNumberPattern := `^\d{3}[- ]?\d{3}[- ]?\d{4}$`

	matched, err := regexp.MatchString(phoneNumberPattern, phoneNumber)

	if err != nil {
		return false
	}

	return matched
}

func IsValidEmail(email string) bool {
	emailPattern := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailPattern, email)

	if err != nil {
		return false
	}

	return matched

}