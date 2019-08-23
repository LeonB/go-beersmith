package bsmx

import (
	"net/url"

	"github.com/leonb/go-beersmith/units"
)

type Styles []Style

type Style struct {
	Category    string                  `xml:"F_S_CATEGORY"`
	Description string                  `xml:"F_S_DESCRIPTION"`
	Examples    string                  `xml:"F_S_EXAMPLES"`
	Guide       string                  `xml:"F_S_GUIDE"`
	Ingredients string                  `xml:"F_S_INGREDIENTS"`
	Letter      int                     `xml:"F_S_LETTER"`
	MaxABV      units.ABV               `xml:"F_S_MAX_ABV"`
	MaxCarb     units.CarbonationVolume `xml:"F_S_MAX_CARB"`
	MaxColor    units.Color             `xml:"F_S_MAX_COLOR"`
	MaxFG       units.Gravity           `xml:"F_S_MAX_FG"`
	MaxIBU      units.IBU               `xml:"F_S_MAX_IBU"`
	MaxOg       units.Gravity           `xml:"F_S_MAX_OG"`
	MinABV      units.ABV               `xml:"F_S_MIN_ABV"`
	MinCarb     units.CarbonationVolume `xml:"F_S_MIN_CARB"`
	MinColor    units.Color             `xml:"F_S_MIN_COLOR"`
	MinFG       units.Gravity           `xml:"F_S_MIN_FG"`
	MinIBU      units.IBU               `xml:"F_S_MIN_IBU"`
	MinOG       units.Gravity           `xml:"F_S_MIN_OG"`
	Name        string                  `xml:"F_S_NAME"`
	Number      string                  `xml:"F_S_NUMBER"`
	Profile     string                  `xml:"F_S_PROFILE"`
	Type        StyleType               `xml:"F_S_TYPE"`
	WebLink     url.URL                 `xml:"F_S_WEB_LINK"`
	MOD         units.Date              `xml:"_MOD_"`
}

type StyleType int

var (
	StyleTypeAle   StyleType = 0
	StyleTypeLager StyleType = 1
	StyleTypeMixed StyleType = 2
	StyleTypeMead  StyleType = 3
	StyleTypeCider StyleType = 4
	StyleTypeWheat StyleType = 5
)

func (st StyleType) String() string {
	switch st {
	case 0:
		return "Ale"
	case 1:
		return "Lager"
	case 2:
		return "Mixed"
	case 3:
		return "Mead"
	case 4:
		return "Cider"
	case 5:
		return "Wheat"
	}

	return ""
}
