package bsmx

import "github.com/leonb/go-beersmith/units"

type Carbonation struct {
	Name        string            `xml:"F_C_NAME"`
	Temperature units.Temperature `xml:"F_C_TEMPERATURE"`
	Type        CarbonationType   `xml:"F_C_TYPE"`
	PrimerName  string            `xml:"F_C_PRIMER_NAME"`
	// effectiveness relative to corn sugar
	CarbRate units.Percentage `xml:"F_C_CARB_RATE"`
	Notes    string           `xml:"F_C_NOTES"`
}

type CarbonationType int

const (
	CarbonationTypeBottle              CarbonationType = 0
	CarbonationTypeKeg                 CarbonationType = 1
	CarbonationTypeKegWithPrimingAgent CarbonationType = 2
)

func (ct CarbonationType) String() string {
	switch ct {
	case 0:
		return "Bottle"
	case 1:
		return "keg"
	case 2:
		return "Keg with priming agent"
	}

	return ""
}
