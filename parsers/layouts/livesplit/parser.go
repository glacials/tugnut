package livesplit

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
)

type parser struct{}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Parse(ctx context.Context, r io.Reader) (Layout, error) {
	b := make([]byte, 1024*1024)
	bytesRead, err := r.Read(b)
	if err != nil {
		return Layout{}, fmt.Errorf("can't read LiveSplit layout: %s")
	}

	log.Printf("LiveSplit layout parser read %d bytes", bytesRead)

	var (
		input  LayoutTag
		output Layout
	)

	err = xml.Unmarshal(b[:bytesRead], &input)
	if err != nil {
		return Layout{}, errors.New(fmt.Sprintf("can't parse LiveSplit file: %s", err))
	}

	p.parseBasicInfo(ctx, &input, &output)

	return output, nil
}

func (p *parser) parseBasicInfo(ctx context.Context, input *LayoutTag, output *Layout) {
}
