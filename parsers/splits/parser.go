package splits

import (
	"errors"
	"io"

	"github.com/glacials/tugnut/parsers/splits/livesplit"
	"github.com/glacials/tugnut/run"
	"golang.org/x/net/context"
)

type Parser interface {
	Parse(context.Context, io.Reader) (run.Run, error)
}

type timer uint

const (
	LiveSplit timer = iota
)

func NewParser(t timer) (Parser, error) {
	if t == LiveSplit {
		return livesplit.NewParser(livesplit.Config{
			ParseSegments:       true,
			ParseSegmentHistory: false,
			ParseRunHistory:     false,
		}), nil
	}

	return nil, errors.New("invalid timer")
}
