package ports

import (
	"walls-otp-service/internal/core/domain/entity"
)

type OtpRepository interface {
	CreateOtp(otp entity.Otp) (interface{}, error)
	ValidateOtp(user_reference string, otp entity.Otp) (interface{}, error)
	GetLastOtp(user_reference string, imei_number string) (interface{}, error)
}
