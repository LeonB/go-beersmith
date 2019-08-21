package bsmx

import "github.com/leonb/go-beersmith/units"

type Archives []Archive

type Archive struct {
	Name      string     `xml:"F_AR_NAME"`
	Date      units.Time `xml:"Date"`
	Action    String     `xml:"F_AR_ACTION"`
	Directory String     `xml:"F_AR_DIRECTORY"`
	File      String     `xml:"F_AR_FILE"`
	_MOD_     units.Date `xml:"_MOD_"`
}
