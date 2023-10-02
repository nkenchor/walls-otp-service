package api

import (
	ports "walls-otp-service/internal/port"
)

// Httphander for the api
type HTTPHandler struct {
	otpService ports.OtpService
}

func NewHTTPHandler(
	countryService ports.OtpService) *HTTPHandler {
	return &HTTPHandler{
		otpService: countryService,
	}
}
