package helper

import (
	"fmt"
	"regexp"
	"time"
	helper "walls-otp-service/internal/core/helper/configuration-helper"

	"github.com/pquerna/otp/totp"
)

func CreateKey() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      helper.ServiceConfiguration.AppName,
		AccountName: helper.ServiceConfiguration.Account,
	})
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func CreateOtp() (string, error) {
	otp, err := totp.GenerateCode(helper.ServiceConfiguration.Key, time.Now())
	if err != nil {
		return "", err
	}

	return otp, nil
}

func ValidateOtp(otp string) bool {
	valid := totp.Validate(otp, helper.ServiceConfiguration.Key)
	return valid
}

func isValidEmail(email string) bool {
	// Regular expression pattern to validate email format
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Regular expression pattern to validate phone number format
	// This example includes support for phone numbers with country code +1 (United States)
	phonePattern := `^\+1\d{10}$`
	match, _ := regexp.MatchString(phonePattern, phoneNumber)
	return match
}

func ValidateEmailOrPhoneNumber(input string) error {
	if isValidEmail(input) {
		return nil // Valid email address
	} else if isValidPhoneNumber(input) {
		return nil // Valid phone number
	} else {
		return fmt.Errorf("%s is neither a valid email address nor a valid phone number", input)
	}
}
