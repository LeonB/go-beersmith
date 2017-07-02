package bsmx

import (
	"net/url"

	"github.com/leonb/go-beersmith/units"
)

type Style struct {
	F_S_CATEGORY    string                  `xml:"F_S_CATEGORY"`
	F_S_DESCRIPTION string                  `xml:"F_S_DESCRIPTION"`
	F_S_EXAMPLES    string                  `xml:"F_S_EXAMPLES"`
	F_S_GUIDE       string                  `xml:"F_S_GUIDE"`
	F_S_INGREDIENTS string                  `xml:"F_S_INGREDIENTS"`
	F_S_LETTER      int                     `xml:"F_S_LETTER"`
	F_S_MAX_ABV     units.ABV               `xml:"F_S_MAX_ABV"`
	F_S_MAX_CARB    units.CarbonationVolume `xml:"F_S_MAX_CARB"`
	F_S_MAX_COLOR   units.Color             `xml:"F_S_MAX_COLOR"`
	F_S_MAX_FG      units.Gravity           `xml:"F_S_MAX_FG"`
	F_S_MAX_IBU     units.IBU               `xml:"F_S_MAX_IBU"`
	F_S_MAX_OG      units.Gravity           `xml:"F_S_MAX_OG"`
	F_S_MIN_ABV     units.ABV               `xml:"F_S_MIN_ABV"`
	F_S_MIN_CARB    units.CarbonationVolume `xml:"F_S_MIN_CARB"`
	F_S_MIN_COLOR   units.Color             `xml:"F_S_MIN_COLOR"`
	F_S_MIN_FG      units.Gravity           `xml:"F_S_MIN_FG"`
	F_S_MIN_IBU     units.IBU               `xml:"F_S_MIN_IBU"`
	F_S_MIN_OG      units.Gravity           `xml:"F_S_MIN_OG"`
	F_S_NAME        string                  `xml:"F_S_NAME"`
	F_S_NUMBER      int                     `xml:"F_S_NUMBER"`
	F_S_PROFILE     string                  `xml:"F_S_PROFILE"`
	F_S_TYPE        StyleType               `xml:"F_S_TYPE"`
	F_S_WEB_LINK    url.URL                 `xml:"F_S_WEB_LINK"`
	_MOD_           units.Date              `xml:"_MOD_"`
}

type StyleType int
