package helpers

import (
	"banklampung-core/errors"
	"encoding/json"
)

func InterfaceToStruct(from interface{}, to interface{}) error {
	marshaledDoc, _ := json.Marshal(from)

	if err := json.Unmarshal(marshaledDoc, to); err != nil {
		return errors.InternalServerError("Cannot parsing string to struct")
	}

	return nil
}

func JsonStringToStruct(jsonString string, to interface{}) error {
	if err := json.Unmarshal([]byte(jsonString), to); err != nil {
		return errors.InternalServerError("Cannot parsing string to struct")
	}

	return nil
}
