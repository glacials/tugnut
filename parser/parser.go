package parser

import (
	"io"

	"github.com/glacials/tugnut/parser/livesplit"
	"github.com/glacials/tugnut/run"
)

type Segment interface {
	Name() string
	Duration() uint
	StartTime() string
	EndTime() string

	ShortestDuration() uint
}

type Parser interface {
	Parse() error
	Game() string
	Category() string
	Attempts() uint
	Segments() ([]run.Segment, error)
}

func New(r io.Reader) Parser {
	return Parser(livesplit.NewParser(r))
}
