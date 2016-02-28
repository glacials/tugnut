package parser

import (
	"strings"
	"testing"

	"github.com/glacials/tugnut/run"
	"golang.org/x/net/context"
)

func TestNew(t *testing.T) {
	ctx := context.Background()

	rc := run.Config{
		Parsables: map[run.Parsable]struct{}{
			run.History:        struct{}{},
			run.Segments:       struct{}{},
			run.SegmentHistory: struct{}{},
		},
	}

	p := New(ctx, rc)
	if p == nil {
		t.Errorf("expected non-nil parser")
	}

	r := strings.NewReader(`<Run></Run>`)

	p.Parse(ctx, r)
}
