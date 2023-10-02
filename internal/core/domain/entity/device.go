package entity

type Device struct {
	DeviceReference string `json:"device_reference" bson:"device_reference"`
	Imei            string `json:"imei" bson:"imei"`
	Type            string `json:"type" bson:"type" validate:"required,eq=mobile|eq=tablet|eq=desktop|eq=phablet|eq=smart_watch"`
	Brand           string `json:"brand" bson:"brand"`
	Model           string `json:"model" bson:"model"`
}
