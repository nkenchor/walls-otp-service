package helper

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateValidContact(fl validator.FieldLevel) bool {
	contact := fl.Field().String()

	// Regular expression patterns for email and phone number
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{1,3}\d{4,}$`

	// Check if the contact matches the email pattern
	isEmail, _ := regexp.MatchString(emailPattern, contact)

	// Check if the contact matches the phone number pattern
	isPhone, _ := regexp.MatchString(phonePattern, contact)

	return isEmail || isPhone
}

func ValidateGUID(fl validator.FieldLevel) bool {
	guid := fl.Field().String()

	// Define the regular expression pattern for a GUID-like string
	// Adjust the pattern according to the specific format you expect
	pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

	// Match the GUID string against the regular expression pattern
	match, _ := regexp.MatchString(pattern, guid)

	return match
}

// ValidateContact validate the contact fields in incoming payload on
// validate otp endpoint
func ValidateContact(contact string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailPattern, contact)

	phonePattern := `^\+\d{1,3}\d{4,}$`
	phoneMatch, _ := regexp.MatchString(phonePattern, contact)

	return emailMatch || phoneMatch
}

func ValidateIMEI(fl validator.FieldLevel) bool {
	imei := fl.Field().String()

	// Define the regular expression pattern for an IMEI number
	// Adjust the pattern according to the specific format you expect
	pattern := `^\d{15}$`

	// Match the IMEI number against the regular expression pattern
	match, _ := regexp.MatchString(pattern, imei)

	return match
}

func GetContactType(contact string) (string, error) {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailPattern, contact)

	phonePattern := `^\+\d{1,3}\d{4,}$`
	phoneMatch, _ := regexp.MatchString(phonePattern, contact)

	if emailMatch {
		return "email", nil
	}
	if phoneMatch {
		return "phone", nil
	}
	return "", errors.New("invalid contact type")
}

// IdentifyCommunicationChannel determines the appropriate channel (sms/email) based on the contact.
func GetChannel(contact string) (string, error) {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailPattern, contact)

	phonePattern := `^\+\d{1,3}\d{4,}$`
	phoneMatch, _ := regexp.MatchString(phonePattern, contact)

	if emailMatch {
		return "email", nil
	}
	if phoneMatch {
		return "sms", nil
	}
	return "", errors.New("invalid contact type")
}
