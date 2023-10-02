package ports

import (
	"walls-otp-service/internal/core/domain/dto"
)

type OtpService interface {
	CreateOtp(createOtpDto dto.CreateOtpDto) (interface{}, error)
	ValidateOtp(user_reference string, validateOtp dto.ValidateOtpDto) (interface{}, error)
	GetLastOtp(user_reference string, imei_number string) (interface{}, error)
	CreateKey() (interface{}, error)
}
