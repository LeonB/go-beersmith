package bsmx

import "github.com/leonb/go-beersmith/units"

type AgeDataList struct {
	_MOD_      units.Date `xml:"_MOD_"`
	Name       string     `xml:"Name"`
	Type       Type       `xml:"Type"`
	Dirty      units.Bool `xml:"Dirty"`
	Owndata    units.Bool `xml:"Owndata"`
	TID        int        `xml:"TID"`
	Size       int        `xml:"Size"`
	_XName     _XName     `xml:"_XName"`
	Allocinc   int        `xml:"Allocinc"`
	AgeDatas   []AgeData  `xml:"Data>AgeData"`
	_TExpanded units.Bool `xml:"_TExpanded"`
}

type AgeData struct {
	Date         units.Time  `xml:"Date"`
	F_AD_GRAVITY float64     `xml:"F_AD_GRAVITY"`
	F_AD_TEMP    float64     `xml:"F_AD_TEMP"`
	F_AD_TYPE    AgeDateType `xml:"F_AD_TYPE"`
}
