package dto

type CreateOtpDto struct {
	UserReference string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
	Contact       string `json:"contact" bson:"contact" validate:"required,valid_contact"`
	Channel       string `json:"channel" bson:"channel" validate:"required,eq=sms|eq=email|in_app"`
	OtpType       string `json:"otp_type" bson:"otp_type" validate:"required"`
	Device        Device `json:"device" bson:"device" validate:"required"`
}
