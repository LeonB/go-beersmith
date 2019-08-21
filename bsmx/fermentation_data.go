package bsmx

import "github.com/leonb/go-beersmith/units"

type FermentationData struct {
	_MOD_                units.Date           `xml:"_MOD_"`
	Name                 string               `xml:"Name"`
	Type                 Type                 `xml:"Type"`
	Dirty                units.Bool           `xml:"Dirty"`
	Owndata              units.Bool           `xml:"Owndata"`
	TID                  int                  `xml:"TID"`
	Size                 int                  `xml:"Size"`
	_XName               string               `xml:"_XName"`
	Allocinc             int                  `xml:"Allocinc"`
	FermentationReadings FermentationReadings `xml:"Data>AgeData"`
	_TExpanded           units.Bool           `xml:"_TExpanded"`
}
