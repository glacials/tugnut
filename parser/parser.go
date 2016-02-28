package parser

import (
	"io"

	"github.com/glacials/tugnut/parser/livesplit"
	"github.com/glacials/tugnut/run"
	"golang.org/x/net/context"
)

type Segment interface {
	Name() string
	Duration() uint
	StartTime() string
	EndTime() string

	ShortestDuration() uint
}

type Parser interface {
	Parse(context.Context, io.Reader) (run.Run, error)
}

func New(ctx context.Context, c run.Config) Parser {
	return Parser(livesplit.NewParser(ctx, c))
}
