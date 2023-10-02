package services

import (
	"context"
	"errors"

	"time"
	publisher "walls-otp-service/internal/adapter/events/publisher"
	"walls-otp-service/internal/core/domain/dto"
	"walls-otp-service/internal/core/domain/entity"
	event "walls-otp-service/internal/core/domain/event/eto"
	"walls-otp-service/internal/core/domain/mapper"
	configuration "walls-otp-service/internal/core/helper/configuration-helper"
	eto "walls-otp-service/internal/core/helper/event-helper/eto"
	logger "walls-otp-service/internal/core/helper/log-helper"
	helper "walls-otp-service/internal/core/helper/otp-helper"
	ports "walls-otp-service/internal/port"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var OtpService = &otpService{}

type otpService struct {
	otpRepository ports.OtpRepository
	redisClient   *redis.Client
}

func NewOtp(otpRepository ports.OtpRepository, redisClient *redis.Client) *otpService {
	OtpService = &otpService{
		otpRepository: otpRepository,
		redisClient:   redisClient,
	}
	return OtpService
}

func (service *otpService) CreateOtp(createOtpDto dto.CreateOtpDto) (interface{}, error) {
	logger.LogEvent("INFO", "Creating OTP")

	mappedOtp := mapper.CreateOtpDtoToOtp(createOtpDto)

	otpCreatedEvent := event.OtpCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "OTPCREATEDEVENT",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          createOtpDto.OtpType,
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: createOtpDto.UserReference,
			EventData:          mappedOtp,
		},
	}
	result, err := service.otpRepository.CreateOtp(mappedOtp)
	if err != nil {
		logger.LogEvent("ERROR", "Unable to create OTP")
		return nil, errors.New("unable to create OTP")
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	ctx := context.Background()
	eventPublisher.PublishOtpCreatedEvent(ctx, otpCreatedEvent, otpCreatedEvent.EventType)

	return result, err
}

func (service *otpService) ValidateOtp(user_reference string, validateOtpDto dto.ValidateOtpDto) (interface{}, error) {

	logger.LogEvent("INFO", "Validating OTP for "+user_reference)

	validOtp, _ := service.GetLastOtp(user_reference, validateOtpDto.Device.Imei)

	if validOtp == nil {
		logger.LogEvent("ERROR", "Invalid OTP for "+user_reference)
		return nil, errors.New("invalid OTP. Please stand advised")
	}

	valid := helper.ValidateOtp(validateOtpDto.Otp)

	if !valid {
		logger.LogEvent("ERROR", "Invalid OTP for "+user_reference)
		return nil, errors.New("invalid OTP. Please stand advised")
	}

	otp := validOtp.(entity.Otp)

	otp.Expired = true

	result, err := service.otpRepository.ValidateOtp(user_reference, otp)
	if err != nil {
		logger.LogEvent("ERROR", "Unable to validate OTP")
		return nil, errors.New("unable to validate OTP")
	}

	validatedOtp := struct {
		UserReference string        `json:"user_reference"`
		Contact       string        `json:"contact"`
		Device        entity.Device `json:"device"`
	}{
		UserReference: otp.UserReference,
		Contact:       otp.Contact,
		Device:        otp.Device,
	}

	otpValidatedEvent := event.OtpValidatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "OTPVALIDATEDEVENT",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          validateOtpDto.OtpType,
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: validateOtpDto.UserReference,
			EventData:          validatedOtp,
		},
	}
	eventPublisher := publisher.NewPublisher(service.redisClient)
	ctx := context.Background()

	eventPublisher.PublishOtpValidatedEvent(ctx, otpValidatedEvent, otpValidatedEvent.EventType)
	return result, err
}
func (service *otpService) GetLastOtp(user_reference string, device_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting the last OTP for user "+user_reference+" and device "+device_reference)
	return service.otpRepository.GetLastOtp(user_reference, device_reference)
}

func (service *otpService) CreateKey() (interface{}, error) {
	return helper.CreateKey()
}
