package dto

type ValidateOtpDto struct {
	UserReference string `json:"user_reference"`
	Otp           string `json:"otp" bson:"otp"`
	Contact       string `json:"contact" bson:"contact"`
	OtpType       string `json:"otp_type" bson:"otp_type"`
	Device        Device `json:"device" bson:"device"`
}
