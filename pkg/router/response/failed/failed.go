package failed

import (
	"encoding/json"
)

type ErrorResponse struct {
	Cause   string `json:"cause"`   // cause of an error
	Message string `json:"message"` // error message
}

// Prepares error response 404
func ProvideNotFound(cause string) []byte {
	result, _ := json.Marshal(
		ErrorResponse{
			Cause:   cause,
			Message: "There is no such shortened link",
		})

	return result
}

// Prepares error response 400
func ProvideBadRequest(cause string) []byte {
	result, _ := json.Marshal(
		ErrorResponse{
			Cause:   cause,
			Message: "The data was entered incorrectly",
		})

	return result
}
