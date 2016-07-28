package livesplit

import (
	"io"

	"golang.org/x/net/context"
)

type parser struct{}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Parse(ctx context.Context, r io.Reader) (Layout, error) {
	return Layout{}, nil
}
