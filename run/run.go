package run

import (
	"time"
)

// Run is a single file from a timer, parsed.
type Run struct {
	Game     Game      `json:"game"`
	Category Category  `json:"category"`
	Attempts uint      `json:"attempts"`
	Segments []Segment `json:"segments"`
}

// Game is the game recorded in the timer's output.
type Game struct {
	Names    []string     `json:"names"`
	SRLInfo  SRLGameInfo  `json:"srl"`
	SRDCInfo SRDCGameInfo `json:"srdc"`
}

// SRLGameInfo contains SpeedRunsLive information about the game, if the timer made it available.
type SRLGameInfo struct {
	ID        string `json:"id"`
	shortname string `json:"shortname"`
}

// SRDCGameInfo contains Speedrun.com information about the game, if the timer made it available.
type SRDCGameInfo struct {
	ID        string `json:"id"`
	shortname string `json:"shortname"`
}

// Category is the category recorded in the timer's output.
type Category struct {
	Names    []string         `json:"names"`
	SRLInfo  SRLCategoryInfo  `json:"srl"`
	SRDCInfo SRDCCategoryInfo `json:"srdc"`
}

// SRLCategoryInfo contains SpeedRunsLive information about the category, if the timer made it available.
type SRLCategoryInfo struct {
	ID string `json:"id"`
}

// SRDCCategoryInfo contains Speedrun.com information about the category, if the timer made it available.
type SRDCCategoryInfo struct {
	ID string `json:"id"`
}

// Segment is the area of time between two split events, or between the start of the run and the first split. Duration,
// StartTime, and EndTime are always available if known, or if they can be calculated from other data in the timer's
// output. ShortestDuration is only available if the timer includes it or supplies full history for each segment.
type Segment struct {
	Name      string   `json:"name"`
	Duration  Duration `json:"duration"`
	StartTime Duration `json:"start_time"`
	EndTime   Duration `json:"end_time"`

	ShortestDuration time.Duration `json:"shortest_duration"`
}

// Duration is a wrapper around time.Duration that includes both a real-world time and a game-world time. Each of these
// is only available if supplied by the timer.
type Duration struct {
	RealTime time.Duration `json:"real"`
	GameTime time.Duration `json:"game"`
}
