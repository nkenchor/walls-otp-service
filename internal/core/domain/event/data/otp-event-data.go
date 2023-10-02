package event

import (
	"walls-otp-service/internal/core/domain/entity"
)

type OtpCreateEventData struct {
	UserReference string        `json:"user_reference"`
	Contact       string        `json:"contact"`
	OtpType       string        `json:"otp_type"`
	Device        entity.Device `json:"device"`
	CreatedOn     string        `json:"created_on"`
}

type OtpValidateEventData struct {
	UserReference string        `json:"user_reference"`
	Otp           string        `json:"otp"`
	Contact       string        `json:"contact"`
	OtpType       string        `json:"otp_type"`
	Device        entity.Device `json:"device"`
}
