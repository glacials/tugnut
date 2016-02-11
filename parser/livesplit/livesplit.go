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
	c run.Config
}

// NewParser constructs and returns a LiveSplit parser. No parsing is performed.
func NewParser(c run.Config) *parser {
	p := parser{
		c: c,
	}

	return &p
}

func (p *parser) Parse(r io.Reader) (run.Run, error) {
	b := make([]byte, 1024*1024)
	bytesRead, err := r.Read(b)
	if err != nil {
		panic("Can't read")
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

	p.parseGeneralInfo(&input, &output)

	if _, ok := p.c.Parsables[run.History]; ok {
		p.parseHistory(&input, &output)
	}

	if _, ok := p.c.Parsables[run.Segments]; ok {
		err := p.parseSegments(&input, &output)
		if err != nil {
			return run.Run{}, fmt.Errorf("can't parse segments: %s", err)
		}
	}

	return output, nil
}

func (p *parser) parseGeneralInfo(input *RunTag, output *run.Run) {
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

func (p *parser) parseHistory(input *RunTag, output *run.Run) {
	// TODO
}

func (p *parser) parseSegments(input *RunTag, output *run.Run) error {
	output.Segments = make([]run.Segment, len(input.Segments.Segments))
	for i, s := range input.Segments.Segments {

		// The current segment's start time is equal to the previous segment's end time
		var startTime run.Duration
		if i == 0 {
			startTime = run.Duration{
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

		endTime := run.Duration{
			RealTime: realEndTime,
			GameTime: gameEndTime,
		}

		output.Segments[i] = run.Segment{
			Name:      s.Name,
			StartTime: startTime,
			EndTime:   endTime,
			Duration: run.Duration{
				RealTime: endTime.RealTime - startTime.RealTime,
				GameTime: endTime.GameTime - startTime.GameTime,
			},
		}
	}
	return nil
}
