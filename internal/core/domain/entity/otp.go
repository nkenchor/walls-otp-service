package entity

type Otp struct {
	Reference     string `json:"reference" bson:"reference"`
	UserReference string `json:"user_reference" bson:"user_reference"`
	Otp           string `json:"otp" bson:"otp"`
	Contact       string `json:"contact" bson:"contact"`
	Channel       string `json:"channel" bson:"channel"`
	OtpType       string `json:"otp_type" bson:"otp_type"`
	Device        Device `json:"device" bson:"device"`
	Expired       bool   `json:"expired" bson:"expired"`
	ExpiresAt     string `json:"expired_at" bson:"expired_at"`
	CreatedOn     string `json:"created_at" bson:"created_at"`
	UpdatedOn     string `json:"updated_on" bson:"updated_on"`
}
