package handlers

import (
	"context"

	"fmt"
	"walls-otp-service/internal/core/domain/dto"

	"walls-otp-service/internal/core/services"

	events "walls-otp-service/internal/core/domain/event/data"

	validation "walls-otp-service/internal/core/helper/validation-helper"
	extraction "walls-otp-service/internal/adapter/handlers/extraction"
)

// Event handler function
func OtpCreateEventHandler(ctx context.Context, event interface{}) {
	event,data, err := extraction.ExtractEventData(event,events.OtpCreateEventData{})
	if err != nil {
		fmt.Println("extracting event:", err)
		return
	}

	iEventData := data.(events.OtpCreateEventData)
	channel, _ := validation.GetChannel(iEventData.Contact)

	createOtpDto := dto.CreateOtpDto{
		UserReference: iEventData.UserReference,
		Contact:       iEventData.Contact,
		Channel:       channel,
		OtpType:       iEventData.OtpType,
		Device: dto.Device{
			DeviceReference: iEventData.Device.DeviceReference,
			Imei:            iEventData.Device.Imei,
			Brand:           iEventData.Device.Brand,
			Model:           iEventData.Device.Model,
			Type:            iEventData.Device.Type,
		},
	}

	services.OtpService.CreateOtp(createOtpDto)

}


