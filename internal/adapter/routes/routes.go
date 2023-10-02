package routes

import (
	docs "walls-otp-service/docs"
	"walls-otp-service/internal/adapter/api"
	configuration "walls-otp-service/internal/core/helper/configuration-helper"
	errorhelper "walls-otp-service/internal/core/helper/error-helper"
	logger "walls-otp-service/internal/core/helper/log-helper"
	message "walls-otp-service/internal/core/helper/message-helper"
	"walls-otp-service/internal/core/middleware"
	"walls-otp-service/internal/core/services"
	ports "walls-otp-service/internal/port"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(otpRepository ports.OtpRepository, redisClient *redis.Client) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	otpService := services.NewOtp(otpRepository, redisClient)

	handler := api.NewHTTPHandler(otpService)

	logger.LogEvent("INFO", "Configuring Routes!")
	router.Use(middleware.LogRequest)

	corrs_config := cors.DefaultConfig()
	corrs_config.AllowAllOrigins = true

	router.Use(cors.New(corrs_config))
	//router.Use(middleware.SetHeaders)

	docs.SwaggerInfo.Description = "Walls OTP Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = configuration.ServiceConfiguration.ServiceName

	router.POST("/api/otp", handler.CreateOtp)
	router.PUT("/api/otp/:user_reference", handler.ValidateOtp)
	router.GET("/api/otp/:user_reference/:device_reference", handler.GetLastOtp)
	router.POST("/api/otp/key", handler.CreateKey)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			errorhelper.ErrorMessage(errorhelper.NoResourceError, message.NoResourceFound))
	})

	return router
}
