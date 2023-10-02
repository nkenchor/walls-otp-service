package publisher

import (
	"context"
	helper "walls-otp-service/internal/core/helper/event-helper"

	"github.com/redis/go-redis/v9"
)

type EventPublisher struct {
	redisClient *redis.Client
}

func NewPublisher(redisClient *redis.Client) *EventPublisher {
	return &EventPublisher{
		redisClient: redisClient,
	}
}

func (p *EventPublisher) PublishOtpCreatedEvent(ctx context.Context, event interface{},eventType ...string) error {

	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event,eventType...)

}

func (p *EventPublisher) PublishOtpValidatedEvent(ctx context.Context, event interface{}, eventType ...string) error {

	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event,eventType...)

}
