package livesplit

import (
	"strings"
	"testing"

	"github.com/glacials/tugnut/run"
	"golang.org/x/net/context"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()

	rc := run.Config{
		Parsables: map[run.Parsable]struct{}{
			run.History:        struct{}{},
			run.Segments:       struct{}{},
			run.SegmentHistory: struct{}{},
		},
	}

	p := NewParser(ctx, rc)
	if p == nil {
		t.Errorf("got a nil parser back")
	}
}

func TestParse(t *testing.T) {
	ctx := context.Background()

	rc := run.Config{
		Parsables: map[run.Parsable]struct{}{
			run.History:        struct{}{},
			run.Segments:       struct{}{},
			run.SegmentHistory: struct{}{},
		},
	}

	p := NewParser(ctx, rc)

	r := strings.NewReader(`<Run></Run>`)

	_, err := p.Parse(ctx, r)
	if err != nil {
		t.Errorf("can't parse: %s", err)
	}
}
