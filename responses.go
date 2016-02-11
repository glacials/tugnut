package main

import (
	"encoding/json"
)

type errorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func jsonErr(msg string, err error) []byte {
	j, err := json.Marshal(errorResponse{
		Error:   msg,
		Details: err.Error(),
	})

	if err != nil {
		return []byte("{\"error\": \"Couldn't even create a proper error message D:\"}")
	}

	return j
}
