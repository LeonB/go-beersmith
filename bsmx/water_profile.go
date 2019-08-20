package bsmx

import "github.com/leonb/go-beersmith/units"

type WaterProfiles []WaterProfile

type WaterProfile struct {
	Order        int          `xml:"F_ORDER"`
	AdditionsVol units.Volume `xml:"F_W_ADDITIONS_VOL"`
	Amount       units.Volume `xml:"F_W_AMOUNT"`
	Bicarb       units.Parts  `xml:"F_W_BICARB"`
	Cacl         units.Parts  `xml:"F_W_CACL"`
	Calcium      units.Parts  `xml:"F_W_CALCIUM"`
	Chalk        units.Weight `xml:"F_W_CHALK"`
	Chloride     units.Parts  `xml:"F_W_CHLORIDE"`
	Epsom        units.Weight `xml:"F_W_EPSOM"`
	Gypsum       units.Weight `xml:"F_W_GYPSUM"`
	Inventory    units.Volume `xml:"F_W_INVENTORY"`
	InRecipe     units.Bool   `xml:"F_W_IN_RECIPE"`
	Magnesium    units.Parts  `xml:"F_W_MAGNESIUM"`
	Name         string       `xml:"F_W_NAME"`
	Notes        string       `xml:"F_W_NOTES"`
	PH           units.PH     `xml:"F_W_PH"`
	Price        units.Price  `xml:"F_W_PRICE"`
	Salt         units.Weight `xml:"F_W_SALT"`
	Soda         units.Weight `xml:"F_W_SODA"`
	Sodium       units.Parts  `xml:"F_W_SODIUM"`
	Sulfate      units.Parts  `xml:"F_W_SULFATE"`
	_MOD_        units.Date   `xml:"_MOD_"`
}
