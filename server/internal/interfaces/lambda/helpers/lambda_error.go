package helpers

import (
	"encoding/json"
	"log"
)

type LambdaError struct {
	Code        string
	Message     string
	Description error
}

func (e LambdaError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b[:])
}

func (e LambdaError) MarshalJSON() []byte {
	bytes, err := json.Marshal(&struct {
		Code             string `json:"code"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}{
		Code:             e.Code,
		Error:            e.Message,
		ErrorDescription: e.Description.Error(),
	})

	if err != nil {
		log.Fatal("failed to marshal error")
		return nil
	}

	return bytes
}
