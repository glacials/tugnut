package parser

import (
	"time"
)

type Attempt struct {
	Duration
}

type Duration struct {
	Realtime time.Duration
	Gametime time.Duration
}
