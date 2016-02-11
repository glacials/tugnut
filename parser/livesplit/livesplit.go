package livesplit

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/glacials/tugnut/run"
)

// parser implements github.com/glacials/tugnut/parser.Parser
type parser struct {
	f io.Reader
	r RunTag
}

// NewParser constructs and returns a LiveSplit parser. No parsing is performed.
func NewParser(r io.Reader) *parser {
	p := parser{
		f: r,
	}

	return &p
}

func (p *parser) Parse() error {
	b := make([]byte, 1024*1024)
	bytesRead, err := p.f.Read(b)
	if err != nil {
		panic("Can't read")
	}

	log.Printf("LiveSplit parser read %d bytes", bytesRead)

	err = xml.Unmarshal(b[:bytesRead], &p.r)
	if err != nil {
		return errors.New(fmt.Sprintf("can't parse LiveSplit file: %s", err))
	}
	return nil
}

func (p *parser) Game() string {
	return p.r.Game
}

func (p *parser) Category() string {
	return p.r.Category
}

func (p *parser) Attempts() uint {
	return p.r.Attempts
}

func (p *parser) Segments() ([]run.Segment, error) {
	segments := make([]run.Segment, len(p.r.Segments.Segments))

	for i, s := range p.r.Segments.Segments {

		// The current segment's start time is equal to the previous segment's end time
		var startTime run.Duration
		if i == 0 {
			startTime = run.Duration{
				RealTime: time.Duration(0),
				GameTime: time.Duration(0),
			}
		} else {
			startTime = segments[i-1].EndTime
		}

		realEndTime, err := parseTime(
			s.SplitTimes.SplitTimes[0].RealTime,
		)
		if err != nil {
			return []run.Segment{}, fmt.Errorf("can't parse segment real time: %s", err)
		}
		gameEndTime, err := parseTime(
			s.SplitTimes.SplitTimes[0].GameTime,
		)
		if err != nil {
			return []run.Segment{}, fmt.Errorf("can't parse segment game time: %s", err)
		}

		endTime := run.Duration{
			RealTime: realEndTime,
			GameTime: gameEndTime,
		}

		segments[i] = run.Segment{
			Name:      s.Name,
			StartTime: startTime,
			EndTime:   endTime,
		}
		segments[i].Duration = run.Duration{
			RealTime: endTime.RealTime - startTime.RealTime,
			GameTime: endTime.GameTime - startTime.GameTime,
		}
	}
	return segments, nil
}
