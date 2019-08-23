package bsmx

import "github.com/leonb/go-beersmith/units"

type CarbonationMethod struct {
	// Priming effectiveness (% relative to corn sugar)
	CarbRate    units.Percentage      `xml:"F_C_CARB_RATE"`
	Name        string                `xml:"F_C_NAME"`
	Notes       string                `xml:"F_C_NOTES"`
	PrimerName  string                `xml:"F_C_PRIMER_NAME"`
	Temperature units.Temperature     `xml:"F_C_TEMPERATURE"`
	Type        CarbonationMethodType `xml:"F_C_TYPE"`
	MOD         units.Date            `xml:"_MOD_"`
}

type CarbonationMethodType int

var (
	CarbonationMethodTypeBottle              CarbonationMethodType = 1
	CarbonationMethodTypeKeg                 CarbonationMethodType = 2
	CarbonationMethodTypeKegWithPrimingAgent CarbonationMethodType = 3
)
