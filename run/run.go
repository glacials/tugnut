package run

import (
	"time"
)

type Segment struct {
	Name      string   `json:"name"`
	Duration  Duration `json:"duration"`
	StartTime Duration `json:"start_time"`
	EndTime   Duration `json:"end_time"`

	ShortestDuration time.Duration `json:"shortest_duration"`
}

type Duration struct {
	RealTime time.Duration `json:"real"`
	GameTime time.Duration `json:"game"`
}
