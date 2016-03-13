package livesplit

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
)

// parser implements github.com/glacials/tugnut/parser.Parser
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
func NewParser(ctx context.Context, c Config) *parser {
	p := parser{
		c: c,
	}

	return &p
}

// Parse reads and parses a LiveSplit file.
func (p *parser) Parse(ctx context.Context, r io.Reader) (Run, error) {
	b := make([]byte, 1024*1024)
	bytesRead, err := r.Read(b)
	if err != nil {
		panic("Can't read")
	}

	log.Printf("LiveSplit parser read %d bytes", bytesRead)

	var (
		input  RunTag
		output Run
	)

	err = xml.Unmarshal(b[:bytesRead], &input)
	if err != nil {
		return Run{}, errors.New(fmt.Sprintf("can't parse LiveSplit file: %s", err))
	}

	p.parseGeneralInfo(ctx, &input, &output)

	if p.c.ParseRunHistory {
		p.parseRunHistory(ctx, &input, &output)
	}

	if p.c.ParseSegments {
		err := p.parseSegments(ctx, &input, &output)
		if err != nil {
			return Run{}, fmt.Errorf("can't parse segments: %s", err)
		}
	}

	if p.c.ParseSegmentHistory {
		err := p.parseSegmentHistory(ctx, &input, &output)
		if err != nil {
			return Run{}, fmt.Errorf("can't parse segment history: %s", err)
		}
	}

	return output, nil
}

func (p *parser) parseGeneralInfo(ctx context.Context, input *RunTag, output *Run) {
	output.Game = Game{
		Names: []string{input.Game},
		SRLInfo: SRLGameInfo{
			ID: "",
		},
		SRDCInfo: SRDCGameInfo{
			ID: "",
		},
	}

	output.Category = Category{
		Names: []string{input.Category},
		SRLInfo: SRLCategoryInfo{
			ID: "",
		},
		SRDCInfo: SRDCCategoryInfo{
			ID: "",
		},
	}

	output.Attempts = input.Attempts
}

func (p *parser) parseRunHistory(ctx context.Context, input *RunTag, output *Run) {
	// TODO
}

func (p *parser) parseSegments(ctx context.Context, input *RunTag, output *Run) error {
	output.Segments = make([]Segment, len(input.Segments.Segments))
	for i, s := range input.Segments.Segments {

		// The current segment's start time is equal to the previous segment's end time
		var startTime Duration
		if i == 0 {
			startTime = Duration{
				RealTime: time.Duration(0),
				GameTime: time.Duration(0),
			}
		} else {
			startTime = output.Segments[i-1].EndTime
		}

		realEndTime, err := parseTime(
			s.SplitTimes.SplitTimes[0].RealTime,
		)
		if err != nil {
			return fmt.Errorf("can't parse segment real time: %s", err)
		}
		gameEndTime, err := parseTime(
			s.SplitTimes.SplitTimes[0].GameTime,
		)
		if err != nil {
			return fmt.Errorf("can't parse segment game time: %s", err)
		}

		endTime := Duration{
			RealTime: realEndTime,
			GameTime: gameEndTime,
		}

		var realDuration time.Duration
		if endTime.RealTime == 0 {
			realDuration = 0
		} else {
			realDuration = endTime.RealTime - startTime.RealTime
		}

		var gameDuration time.Duration
		if endTime.GameTime == 0 {
			gameDuration = 0
		} else {
			gameDuration = endTime.GameTime - startTime.GameTime
		}

		duration := Duration{
			RealTime: realDuration,
			GameTime: gameDuration,
		}

		output.Segments[i] = Segment{
			Name:      s.Name,
			StartTime: startTime,
			EndTime:   endTime,
			Duration:  duration,
		}
	}
	return nil
}

func (p *parser) parseSegmentHistory(ctx context.Context, input *RunTag, output *Run) error {
	// TODO
	return nil
}
