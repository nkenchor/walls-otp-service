package api

import (
	"fmt"
	"walls-otp-service/internal/core/domain/dto"
	errorhelper "walls-otp-service/internal/core/helper/error-helper"
	validation "walls-otp-service/internal/core/helper/validation-helper"

	"github.com/gin-gonic/gin"
)

// @Summary Create OTP
// @Description Create an OTP
// @Tags OTP
// @Accept json
// @Produce json
// @Success 200 {string} entity.Otp.Reference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CreateOtpDto true "OTP request body"
// @Router /api/otp [post]
// Need to add a channel and contact field to the swagger docs
func (hdl *HTTPHandler) CreateOtp(c *gin.Context) {
	body := dto.CreateOtpDto{}
	_ = c.BindJSON(&body)

	err := validation.Validate(&body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	otp, err := hdl.otpService.CreateOtp(body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"reference:": otp})
}

// @Summary Validate OTP
// @Description Validate an OTP
// @Tags OTP
// @Accept json
// @Produce json
// @Success 200
// @Success 200 {string} entity.Otp.Reference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.ValidateOtpDto true "OTP request body"
// @Router /api/otp/{user-reference} [PUT]
func (hdl *HTTPHandler) ValidateOtp(c *gin.Context) {
	body := dto.ValidateOtpDto{}
	_ = c.BindJSON(&body)
	if err := validation.Validate(&body); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	otp, err := hdl.otpService.ValidateOtp(c.Param("user_reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.InvalidOtp, err.Error()))
		return

	}
	c.JSON(200, otp)
}

// @Summary Get the last OTP
// @Description Retrieve the last OTP for a user and device
// @Tags OTP
// @Accept json
// @Produce json
// @Success 200 {string} entity.Otp.Otp "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param device_reference path string true "Device reference"
// @Router /api/otp/{user_reference}/{device_reference} [get]
func (hdl *HTTPHandler) GetLastOtp(c *gin.Context) {
	otp, err := hdl.otpService.GetLastOtp(c.Param("user_reference"), c.Param("device_reference"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.NoRecordError, err.Error()))
		return
	}

	fmt.Println("last otp:", otp)

	c.JSON(200, otp)
}

// @Summary Create OTP Key
// @Description Generate a new OTP key
// @Tags OTP
// @Accept json
// @Produce json
// @Success 200 {string} entity.Otp.Reference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/otp/key [post]
func (hdl *HTTPHandler) CreateKey(c *gin.Context) {
	key, err := hdl.otpService.CreateKey()

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.InvalidKey, err.Error()))
		return
	}

	c.JSON(200, gin.H{"key": key})
}
