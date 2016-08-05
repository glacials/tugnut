package livesplit

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/glacials/tugnut/run"
	"golang.org/x/net/context"
)

type parser struct {
	c Config
}

// Config specifies which parts of the run a parser should parse. It is moderately expensive to parse segments (n) and
// run history (m), and very expensive to parse segment history (n*m). You should tailor your parser's configuration
// with these facts in mind.
//
// General run info is always parsed.
type Config struct {
	ParseSegments       bool
	ParseSegmentHistory bool
	ParseRunHistory     bool
}

// NewParser constructs and returns a LiveSplit parser. No parsing is performed.
func NewParser(c Config) *parser {
	return &parser{
		c: c,
	}
}

// Parse reads and parses a LiveSplit file.
func (p *parser) Parse(ctx context.Context, r io.Reader) (run.Run, error) {
	b := make([]byte, 1024*1024)
	bytesRead, err := r.Read(b)
	if err != nil {
		return run.Run{}, fmt.Errorf("can't read LiveSplit splits: %s", err)
	}

	log.Printf("LiveSplit parser read %d bytes", bytesRead)

	var (
		input  RunTag
		output run.Run
	)

	err = xml.Unmarshal(b[:bytesRead], &input)
	if err != nil {
		return run.Run{}, errors.New(fmt.Sprintf("can't parse LiveSplit file: %s", err))
	}

	p.parseBasicInfo(ctx, &input, &output)

	if p.c.ParseRunHistory {
		p.parseRunHistory(ctx, &input, &output)
	}

	if p.c.ParseSegments {
		err := p.parseSegments(ctx, &input, &output)
		if err != nil {
			return run.Run{}, fmt.Errorf("can't parse segments: %s", err)
		}
	}

	if p.c.ParseSegmentHistory {
		err := p.parseSegmentHistory(ctx, &input, &output)
		if err != nil {
			return run.Run{}, fmt.Errorf("can't parse segment history: %s", err)
		}
	}

	return output, nil
}

func (p *parser) parseBasicInfo(ctx context.Context, input *RunTag, output *run.Run) {
	output.Game = run.Game{
		Names: []string{input.Game},
		SRLInfo: run.SRLGameInfo{
			ID: "",
		},
		SRDCInfo: run.SRDCGameInfo{
			ID: "",
		},
	}

	output.Category = run.Category{
		Names: []string{input.Category},
		SRLInfo: run.SRLCategoryInfo{
			ID: "",
		},
		SRDCInfo: run.SRDCCategoryInfo{
			ID: "",
		},
	}

	output.Attempts = input.Attempts
}

func (p *parser) parseRunHistory(ctx context.Context, input *RunTag, output *run.Run) {
	// TODO
}

func (p *parser) parseSegments(ctx context.Context, input *RunTag, output *run.Run) error {
	output.Segments = make([]run.Segment, len(input.Segments.Segments))
	for i, s := range input.Segments.Segments {

		// The current segment's start time is equal to the previous segment's end time
		var start, end run.Duration
		if i == 0 {
			start = run.Duration{
				RealTime: time.Duration(0),
				GameTime: time.Duration(0),
			}
		} else {
			start = output.Segments[i-1].End
		}

		realEnd, err := parseTime(
			s.SplitTimes.SplitTimes[0].RealTime,
		)
		if err != nil {
			return fmt.Errorf("can't parse segment real time: %s", err)
		}
		gameEnd, err := parseTime(
			s.SplitTimes.SplitTimes[0].GameTime,
		)
		if err != nil {
			return fmt.Errorf("can't parse segment game time: %s", err)
		}

		end = run.Duration{
			RealTime: realEnd,
			GameTime: gameEnd,
		}

		var realDuration time.Duration
		if end.RealTime == 0 {
			realDuration = 0
		} else {
			realDuration = end.RealTime - start.RealTime
		}

		var gameDuration time.Duration
		if end.GameTime == 0 {
			gameDuration = 0
		} else {
			gameDuration = end.GameTime - start.GameTime
		}

		duration := run.Duration{
			RealTime: realDuration,
			GameTime: gameDuration,
		}

		output.Segments[i] = run.Segment{
			Name:     s.Name,
			Start:    start,
			End:      end,
			Duration: duration,
		}
	}
	return nil
}

func (p *parser) parseSegmentHistory(ctx context.Context, input *RunTag, output *run.Run) error {
	// TODO
	return nil
}
