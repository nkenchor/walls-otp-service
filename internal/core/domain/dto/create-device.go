package dto

type Device struct {
	DeviceReference string `json:"device_reference" validate:"required,uuid4"`
	Imei            string `json:"imei" bson:"imei" validate:"required,imei,min=10,max=50"`
	Type            string `json:"type" bson:"type" validate:"required,eq=mobile|eq=tablet|eq=desktop|eq=phablet|eq=smart_watch"`
	Brand           string `json:"brand" bson:"brand" validate:"required"`
	Model           string `json:"model" bson:"model" validate:"required"`
}
