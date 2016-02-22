package responses

import (
	"encoding/json"
)

type errorResponse struct {
	Error   string `json:"error"`
	Details error  `json:"details,omitempty"`
}

func JSONErr(msg string, err error) []byte {
	j, err := json.Marshal(errorResponse{
		Error:   msg,
		Details: err,
	})

	if err != nil {
		return []byte("{\"error\": \"Couldn't even create a proper error message D:\"}")
	}

	return j
}
