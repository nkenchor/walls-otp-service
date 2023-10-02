package repository

import (
	"context"
	"walls-otp-service/internal/core/domain/entity"

	logger "walls-otp-service/internal/core/helper/log-helper"
	ports "walls-otp-service/internal/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OtpInfra struct {
	Collection *mongo.Collection
}

func NewOtp(Collection *mongo.Collection) *OtpInfra {
	return &OtpInfra{Collection}
}

// UserRepo implements the repository.UserRepository interface
var _ ports.OtpRepository = &OtpInfra{}

func (r *OtpInfra) CreateOtp(otp entity.Otp) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting otp with reference: "+otp.Reference)

	_, err := r.Collection.InsertOne(context.TODO(), otp)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Persisting otp with reference: "+otp.Reference+" completed successfully...")
	return otp.Reference, nil
}

func (r *OtpInfra) ValidateOtp(user_reference string, otp entity.Otp) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting otp with user reference: "+user_reference)
	validatedOtp := entity.Otp{}
	err := r.Collection.FindOneAndReplace(
		context.TODO(),
		bson.M{"user_reference": bson.M{"$eq": user_reference}}, otp).Decode(validatedOtp)
	if err != nil || validatedOtp == (entity.Otp{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Persisting otp with user reference: "+user_reference+" completed successfully. ")
	return otp.Reference, nil
}

func (r *OtpInfra) GetLastOtp(user_reference string, device_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving last otp with user reference: "+user_reference+" and device reference: "+device_reference)
	otp := entity.Otp{}
	filter := bson.M{"user_reference": user_reference, "device.reference": device_reference}
	//Sort in descending order
	opts := options.FindOne().SetSort(bson.M{"_id": -1})
	err := r.Collection.FindOne(context.TODO(), filter, opts).Decode(&otp)
	if err != nil || otp == (entity.Otp{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving last otp with user reference: "+user_reference+" and device reference: "+device_reference+" completed successfully. ")
	return otp, nil
}
