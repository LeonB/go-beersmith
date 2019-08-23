package bsmx

import "github.com/leonb/go-beersmith/units"

type Ingredients struct {
	MOD      units.Date `xml:"_MOD_"`
	Name     string     `xml:"Name"`
	Type     Type       `xml:"Type"`
	Dirty    units.Bool `xml:"Dirty"`
	Owndata  units.Bool `xml:"Owndata"`
	TID      int        `xml:"TID"`
	Size     int        `xml:"Size"`
	XName    String     `xml:"_XName"`
	Allocinc int        `xml:"Allocinc"`
	// Data       *Data      `xml:"Data"`
	Grains    Grains     `xml:"Data>Grain"`
	Hops      Hops       `xml:"Data>Hops"`
	Yeasts    Yeasts     `xml:"Data>Yeast"`
	Miscs     Miscs      `xml:"Data>Misc"`
	TExpanded units.Bool `xml:"_TExpanded"`
}
