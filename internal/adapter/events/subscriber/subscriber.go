package subscriber

import (
	"context"

	"walls-otp-service/internal/adapter/handlers"
	helper "walls-otp-service/internal/core/helper/event-helper"

	"github.com/redis/go-redis/v9"
)

type EventSubscriber struct {
	redisClient *redis.Client
}

func NewEventSubscriber(redisClient *redis.Client) *EventSubscriber {
	return &EventSubscriber{
		redisClient: redisClient,
	}
}

func (s *EventSubscriber) SubscribeToCreateOtpEvent(ctx context.Context, event interface{}) error {
	redisHelper := helper.NewRedisClient(s.redisClient)
	return redisHelper.SubscribeToEvent(ctx, event, handlers.OtpCreateEventHandler)
}
func (s *EventSubscriber) SubscribeToValidateOtpEvent(ctx context.Context, event interface{}) error {
	redisHelper := helper.NewRedisClient(s.redisClient)
	return redisHelper.SubscribeToEvent(ctx, event, handlers.OtpCreateEventHandler)
}
