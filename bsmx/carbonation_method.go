package bsmx

import "github.com/leonb/go-beersmith/units"

type CarbonationMethod struct {
	F_C_CARB_RATE   float64               `xml:"F_C_CARB_RATE"`
	F_C_NAME        string                `xml:"F_C_NAME"`
	F_C_NOTES       string                `xml:"F_C_NOTES"`
	F_C_PRIMER_NAME string                `xml:"F_C_PRIMER_NAME"`
	F_C_TEMPERATURE units.Temperature     `xml:"F_C_TEMPERATURE"`
	F_C_TYPE        CarbonationMethodType `xml:"F_C_TYPE"`
	_MOD_           units.Date            `xml:"_MOD_"`
}

type CarbonationMethodType int
