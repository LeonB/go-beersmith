package bsmx

import "github.com/leonb/go-beersmith/units"

type CarbonationProfiles []CarbonationProfile

type CarbonationProfile struct {
	MOD         units.Date             `xml:"_MOD_"`
	Name        string                 `xml:"F_C_NAME"`
	Temperature units.Temperature      `xml:"F_C_TEMPERATURE"`
	Type        CarbonationProfileType `xml:"F_C_TYPE"`
	PrimerName  string                 `xml:"F_C_PRIMER_NAME"`
	// effectiveness relative to corn sugar
	CarbRate units.Percentage `xml:"F_C_CARB_RATE"`
	Notes    string           `xml:"F_C_NOTES"`
}

type CarbonationProfileType int

const (
	CarbonationProfileTypeBottle              CarbonationProfileType = 0
	CarbonationProfileTypeKeg                 CarbonationProfileType = 1
	CarbonationProfileTypeKegWithPrimingAgent CarbonationProfileType = 2
)

func (cpt CarbonationProfileType) String() string {
	switch cpt {
	case 0:
		return "Bottle"
	case 1:
		return "keg"
	case 2:
		return "Keg with priming agent"
	}

	return ""
}
