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
	Parse(io.Reader) (run.Run, error)
}

func New(c run.Config) Parser {
	return Parser(livesplit.NewParser(c))
}
