package livesplit

import (
	"encoding/xml"
)

type LayoutTag struct {
	XMLName          xml.Name      `xml:"Layout,name"`
	Mode             string        `xml:"Mode"`
	X                int           `xml:"X"`
	Y                int           `xml:"Y"`
	VerticalWidth    int           `xml:"VerticalWidth"`
	VerticalHeight   int           `xml:"VerticalHeight"`
	HorizontalWidth  int           `xml:"HorizontalWidth"`
	HorizontalHeight int           `xml:"HorizontalHeight"`
	Settings         SettingsTag   `xml:"Settings"`
	Components       ComponentsTag `xml:"Components"`
}

type SettingsTag struct {
	XMLName                xml.Name `xml:"Settings,name"`
	TextColor              string   `xml:"TextColor"`
	BackgroundColor        string   `xml:"BackgroundColor"`
	BackgroundColor2       string   `xml:"BackgroundColor2"`
	ThinSeparatorsColor    string   `xml:"ThinSeparatorsColor"`
	SeparatorsColor        string   `xml:"SeparatorsColor"`
	PersonalBestColor      string   `xml:"PersonalBestColor"`
	AheadGainingTimeColor  string   `xml:"AheadGainingTimeColor"`
	AheadLosingTimeColor   string   `xml:"AheadLosingTimeColor"`
	BehindGainingTimeColor string   `xml:"BehindGainingTimeColor"`
	BehindLosingTimeColor  string   `xml:"BehindLosingTimeColor"`
	BestSegmentColor       string   `xml:"BestSegmentColor"`
	UseRainbowColor        string   `xml:"UseRainbowColor"`
	NotRunningColor        string   `xml:"NotRunningColor"`
	PausedColor            string   `xml:"PausedColor"`
	ShadowsColor           string   `xml:"ShadowsColor"`
	TimesFont              string   `xml:"TimesFont"`
	TimerFont              string   `xml:"TimerFont"`
	TextFont               string   `xml:"TextFont"`
	AlwaysOnTop            bool     `xml:"AlwaysOnTop"`
	ShowBestSegments       bool     `xml:"ShowBestSegments"`
	AntiAliasing           bool     `xml:"AntiAliasing"`
	DropShadows            bool     `xml:"DropShadows"`
	BackgroundType         string   `xml:"BackgroundType"`
	BackgroundImage        string   `xml:"BackgroundImage"`
	ImageOpacity           float64  `xml:"ImageOpacity"`
	ImageBlur              float64  `xml:"ImageBlur"`
	Opacity                float64  `xml:"Opacity"`
}

type ComponentsTag struct {
	XMLName    xml.Name       `xml:"Components,name"`
	Components []ComponentTag `xml:"Component"`
}

type ComponentTag struct {
	XMLName  xml.Name             `xml:"Component,name"`
	Settings ComponentSettingsTag `xml:"Settings"`
}

type ComponentSettingsTag struct {
	XMLName xml.Name `xml:"Settings,name"`
	Version string   `xml:"Version"`

	// Common
	BackgroundColor    string `xml:"BackgroundColor"`
	BackgroundColor2   string `xml:"BackgroundColor2"`
	BackgroundGradient string `xml:"BackgroundGradient"`

	// LiveSplit.Timer.dll
	TimerHeight         int    `xml:"TimerHeight"`
	TimerWidth          int    `xml:"TimerWidth"`
	TimerFormat         string `xml:"TimerFormat"`
	OverrideSplitColors bool   `xml:"OverrideSplitColors"`
	ShowGradient        bool   `xml:"ShowGradient"`
	TimerColor          string `xml:"TimerColor"`
	CenterTimer         bool   `xml:"CenterTimer"`
	TimingMethod        string `xml:"TimingMethod"`
	DecimalsSize        int    `xml:"DecimalsSize"`

	// LiveSplit.RunPrediction.dll
	TextColor         string `xml:"TextColor"`
	OverrideTextColor bool   `xml:"OverrideTextColor"`
	TimeColor         string `xml:"TimeColor"`
	OverrideTimeColor bool   `xml:"OverrideTimeColor"`
	Accuracy          string `xml:"Accuracy"`
	Comparison        string `xml:"Comparison"`
	Display2Rows      bool   `xml:"Display2Rows"`
}
