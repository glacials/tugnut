package livesplit

import (
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()

	rc := Config{
		ParseSegments:       false,
		ParseSegmentHistory: false,
		ParseRunHistory:     false,
	}

	p := NewParser(ctx, rc)
	if p == nil {
		t.Errorf("got a nil parser back")
	}
}

func TestParse(t *testing.T) {
	ctx := context.Background()

	rc := Config{
		ParseSegments:       false,
		ParseSegmentHistory: false,
		ParseRunHistory:     false,
	}

	p := NewParser(ctx, rc)

	r := strings.NewReader(`<Run></Run>`)

	_, err := p.Parse(ctx, r)
	if err != nil {
		t.Errorf("can't parse: %s", err)
	}
}
