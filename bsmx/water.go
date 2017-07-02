package bsmx

import "github.com/leonb/go-beersmith/units"

type Water struct {
	F_ORDER           int          `xml:"F_ORDER"`
	F_W_ADDITIONS_VOL units.Volume `xml:"F_W_ADDITIONS_VOL"`
	F_W_AMOUNT        units.Volume `xml:"F_W_AMOUNT"`
	F_W_BICARB        units.Parts  `xml:"F_W_BICARB"`
	F_W_CACL          units.Parts  `xml:"F_W_CACL"`
	F_W_CALCIUM       units.Parts  `xml:"F_W_CALCIUM"`
	F_W_CHALK         units.Weight `xml:"F_W_CHALK"`
	F_W_CHLORIDE      units.Parts  `xml:"F_W_CHLORIDE"`
	F_W_EPSOM         units.Weight `xml:"F_W_EPSOM"`
	F_W_GYPSUM        units.Weight `xml:"F_W_GYPSUM"`
	F_W_INVENTORY     units.Volume `xml:"F_W_INVENTORY"`
	F_W_IN_RECIPE     units.Bool   `xml:"F_W_IN_RECIPE"`
	F_W_MAGNESIUM     units.Parts  `xml:"F_W_MAGNESIUM"`
	F_W_NAME          string       `xml:"F_W_NAME"`
	F_W_NOTES         string       `xml:"F_W_NOTES"`
	F_W_PH            units.PH     `xml:"F_W_PH"`
	F_W_PRICE         units.Price  `xml:"F_W_PRICE"`
	F_W_SALT          units.Weight `xml:"F_W_SALT"`
	F_W_SODA          units.Weight `xml:"F_W_SODA"`
	F_W_SODIUM        units.Parts  `xml:"F_W_SODIUM"`
	F_W_SULFATE       units.Parts  `xml:"F_W_SULFATE"`
	_MOD_             units.Date   `xml:"_MOD_"`
}
