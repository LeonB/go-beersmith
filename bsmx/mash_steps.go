package bsmx

import "github.com/leonb/go-beersmith/units"

type MashSteps struct {
	_MOD_    units.Date `xml:"_MOD_"`
	Name     string     `xml:"Name"`
	Type     Type       `xml:"Type"`
	Dirty    units.Bool `xml:"Dirty"`
	Owndata  units.Bool `xml:"Owndata"`
	TID      int        `xml:"TID"`
	Size     int        `xml:"Size"`
	_XName   *_XName    `xml:"_XName"`
	Allocinc int        `xml:"Allocinc"`
	// Data       *Data      `xml:"Data"`
	MashSteps  []MashStep `xml:"Data>MashStep"`
	_TExpanded units.Bool `xml:"_TExpanded"`
}
