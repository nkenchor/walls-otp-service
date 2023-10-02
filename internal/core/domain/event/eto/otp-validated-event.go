package event

import (
	"walls-otp-service/internal/core/helper/event-helper/eto"
)

type OtpValidatedEvent struct {
	eto.Event
}
