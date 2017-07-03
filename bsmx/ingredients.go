package bsmx

import "github.com/leonb/go-beersmith/units"

type Ingredients struct {
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
	Grains     []Grain    `xml:"Data>Grain"`
	Hops       Hops       `xml:"Data>Hop"`
	Yeasts     []Yeast    `xml:"Data>Yeast"`
	Miscs      []Misc     `xml:"Data>Misc"`
	_TExpanded units.Bool `xml:"_TExpanded"`
}