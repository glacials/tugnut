package main

import (
	"encoding/json"

	"github.com/glacials/tugnut/parser"
	"github.com/glacials/tugnut/run"
)

type runResponse struct {
	Game     string        `json:"game"`
	Category string        `json:"category"`
	Attempts uint          `json:"attempts"`
	Segments []run.Segment `json:"segments"`
}

type runSegment struct {
	Name      string `json:"name"`
	Duration  uint   `json:"duration"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`

	ShortestDuration uint `json:"shortest_duration"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func jsonRun(p parser.Parser) []byte {
	segments, err := p.Segments()
	if err != nil {
		return jsonErr("Error reading run segments.", err)
	}

	res := runResponse{
		Game:     p.Game(),
		Category: p.Category(),
		Attempts: p.Attempts(),
		Segments: segments,
	}

	j, err := json.Marshal(res)
	if err != nil {
		return jsonErr("Run was valid, but encountered an error preparing it.", nil)
	}
	return j
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
