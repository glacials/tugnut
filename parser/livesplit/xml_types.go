package livesplit

import (
	"encoding/xml"
)

type RunTag struct {
	XMLName        xml.Name          `xml:"Run,name"`
	Game           string            `xml:"GameName"`
	Category       string            `xml:"CategoryName"`
	Metadata       MetadataTag       `xml:"Metadata"`
	Offset         string            `xml:"Offset"`
	Attempts       uint              `xml:"AttemptCount"`
	AttemptHistory AttemptHistoryTag `xml:"AttemptHistory"`
	Segments       SegmentsTag       `xml:"Segments"`
}

type MetadataTag struct {
	XMLName xml.Name `xml:"Metadata",name`
	Run     string   `xml:"Run"`
}

type AttemptHistoryTag struct {
	XMLName  xml.Name     `xml:"AttemptHistory,name"`
	Attempts []AttemptTag `xml:"Attempt"`
}

type AttemptTag struct {
	XMLName  xml.Name `xml:"Attempt,name"`
	RealTime string   `xml:"RealTime"`
	GameTime string   `xml:"GameTime"`
}

type SegmentsTag struct {
	XMLName  xml.Name      `xml:"Segments,name"`
	Segments []*SegmentTag `xml:"Segment"`
}

type SegmentTag struct {
	XMLName         xml.Name           `xml:"Segment,name"`
	Name            string             `xml:"name"`
	SplitTimes      SplitTimesTag      `xml:"SplitTimes"`
	BestSegmentTime BestSegmentTimeTag `xml:"BestSegmentTime"`
	SegmentHistory  SegmentHistoryTag  `xml:"SegmentHistory"`
}

type SplitTimesTag struct {
	XMLName    xml.Name       `xml:"SplitTimes,name"`
	SplitTimes []SplitTimeTag `xml:"SplitTime"`
}

type SplitTimeTag struct {
	XMLName  xml.Name `xml:"SplitTime,name"`
	RealTime string   `xml:"RealTime"`
	GameTime string   `xml:"GameTime"`
}

type BestSegmentTimeTag struct {
	XMLName  xml.Name `xml:"BestSegmentTime,name"`
	RealTime string   `xml:"RealTime"`
	GameTime string   `xml:"GameTime"`
}

type SegmentHistoryTag struct {
	XMLName xml.Name  `xml:"SegmentHistory,name"`
	Time    []TimeTag `xml:"Time"`
}

type TimeTag struct {
	XMLName  xml.Name `xml:"Time,name"`
	RealTime string   `xml:"RealTime"`
	GameTime string   `xml:"GameTime"`
}
