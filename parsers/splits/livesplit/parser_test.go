package livesplit

import (
	"strings"
	"testing"

	"golang.org/x/net/context"
)

func TestNewParser(t *testing.T) {
	c := Config{
		ParseSegments:       false,
		ParseSegmentHistory: false,
		ParseRunHistory:     false,
	}

	p := NewParser(c)
	if p == nil {
		t.Errorf("got a nil parser back")
	}
}

func TestParse(t *testing.T) {
	ctx := context.Background()

	c := Config{
		ParseSegments:       false,
		ParseSegmentHistory: false,
		ParseRunHistory:     false,
	}

	p := NewParser(c)

	r := strings.NewReader(`<Run></Run>`)

	_, err := p.Parse(ctx, r)
	if err != nil {
		t.Errorf("can't parse: %s", err)
	}
}
