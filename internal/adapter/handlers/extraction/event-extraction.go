package handlers

import (
	"encoding/json"
	"fmt"

	eto "walls-otp-service/internal/core/helper/event-helper/eto"
)

// Event handler function
// extractEventData takes in an event and extracts the otpValidatedEventData from it.
func ExtractEventData(event interface{}, data interface{}) (interface{},interface {}, error) {
	var iEvent eto.Event
	err := convertEvent(event, &iEvent)
	if err != nil {
		return nil,nil, fmt.Errorf("error converting event to validated event: %v", err)
	}

	var iEventData interface{}
	err = convertEvent(iEvent.EventData, &iEventData)
	if err != nil {
		return nil, nil, fmt.Errorf("error converting event data to Data: %v", err)
	}

	return iEvent,iEventData, nil
}

func convertEvent(event interface{}, outputEvent interface{}) error {
	// Convert interface{} to byte array
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Deserialize JSON to outputEvent
	err = json.Unmarshal(jsonBytes, outputEvent)
	if err != nil {
		return err
	}

	return nil
}
