package handlers

import (
	"context"
	"fmt"
	"walls-otp-service/internal/core/domain/dto"
	"walls-otp-service/internal/core/services"

	events "walls-otp-service/internal/core/domain/event/data"
	

	extraction "walls-otp-service/internal/adapter/handlers/extraction"

)


func OtpValidateEventHandler(ctx context.Context, event interface{}) {
	event,data, err := extraction.ExtractEventData(event,events.OtpValidateEventData{})
	if err != nil {
		fmt.Println("extracting event:", err)
		return
	}

	iEventData := data.(events.OtpValidateEventData)
	

	validateOtpDto := dto.ValidateOtpDto{
		UserReference: iEventData.UserReference,
		Contact:       iEventData.Contact,
		Otp:           iEventData.Otp,
		OtpType:       iEventData.OtpType,
		Device: dto.Device{
			DeviceReference: iEventData.Device.DeviceReference,
			Imei:            iEventData.Device.Imei,
			Brand:           iEventData.Device.Brand,
			Model:           iEventData.Device.Model,
			Type:            iEventData.Device.Type,
		},
	}

	// Create an instance of the OtpService
	services.OtpService.ValidateOtp(iEventData.UserReference, validateOtpDto)

}




