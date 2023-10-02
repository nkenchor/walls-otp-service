package mapper

import (
	"strconv"
	"time"
	"walls-otp-service/internal/core/domain/dto"
	"walls-otp-service/internal/core/domain/entity"
	config "walls-otp-service/internal/core/helper/configuration-helper"
	logger "walls-otp-service/internal/core/helper/log-helper"
	helper "walls-otp-service/internal/core/helper/otp-helper"

	"github.com/google/uuid"
)

func CreateOtpDtoToOtp(otpDto dto.CreateOtpDto) entity.Otp {
	otpExpiry, _ := strconv.Atoi(config.ServiceConfiguration.OtpExpiry)
	otp, err := helper.CreateOtp()

	if err != nil {
		logger.LogEvent("ERROR", "Error while creating OTP")
	}
	otpMap := entity.Otp{
		Reference:     uuid.New().String(),
		UserReference: otpDto.UserReference,
		Otp:           otp,
		Contact:       otpDto.Contact,
		Channel:       otpDto.Channel,
		OtpType:       otpDto.OtpType,
		Device: entity.Device{
			DeviceReference: uuid.New().String(),
			Imei:            otpDto.Device.Imei,
			Type:            otpDto.Device.Type,
			Model:           otpDto.Device.Model,
			Brand:           otpDto.Device.Brand,
		},
		Expired:   false,
		ExpiresAt: time.Now().Add(time.Duration(otpExpiry) * time.Second).Format(time.RFC3339),
		CreatedOn: time.Now().Format(time.RFC3339),
		UpdatedOn: time.Now().Format(time.RFC3339),
	}
	return otpMap
}
