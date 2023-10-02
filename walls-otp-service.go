package main

import (
	"context"
	"walls-otp-service/internal/adapter/events/subscriber"
	extensions "walls-otp-service/internal/adapter/extensions"
	mongoRepository "walls-otp-service/internal/adapter/repository/mongodb"

	"fmt"
	"walls-otp-service/internal/adapter/routes"
	channel "walls-otp-service/internal/core/domain/event/channel"
	configuration "walls-otp-service/internal/core/helper/configuration-helper"
	logger "walls-otp-service/internal/core/helper/log-helper"
	message "walls-otp-service/internal/core/helper/message-helper"
)

func main() {
	//Initialize request Log
	logger.InitializeLog()
	//Start DB Connection
	mongoRepo := extensions.StartDatabase("mongodb")

	logger.LogEvent("INFO", "MongoDB Connected and Initialized!")

	logger.LogEvent("INFO", message.StartingRedis)
	redisClient := extensions.StartEventBus("redis")
	ctx := context.Background()

	//Set up routes
	router := routes.SetupRouter(mongoRepo.(mongoRepository.MongoRepositories).Otp, redisClient)

	config := configuration.ServiceConfiguration

	go func() {
		logger.LogEvent("INFO", message.StartingServer)
		err := router.Run(":" + config.ServicePort)
		//api.SetConfiguration
		if err != nil {
			fmt.Println(err)
			logger.LogEvent("ERROR", "Error Starting Server : "+err.Error())
		}
	}()

	// Initialize the event subscriber
	eventSubscriber := subscriber.NewEventSubscriber(redisClient)
	// Run the subscription code in a Goroutine
	go func() {
		eventSubscriber.SubscribeToCreateOtpEvent(ctx, channel.CreateOtpEvent)
	}()

	go func() {
		eventSubscriber.SubscribeToValidateOtpEvent(ctx, channel.ValidateOtpEvent)
	}()

	select {}
}
