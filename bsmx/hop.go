package bsmx

import "github.com/leonb/go-beersmith/units"

type Hops []Hop

type Hop struct {
	MOD       units.Date       `xml:"_MOD_"`
	Name      string           `xml:"F_H_NAME"`
	Origin    string           `xml:"F_H_ORIGIN"`
	Type      HopType          `xml:"F_H_TYPE"`
	Form      HopForm          `xml:"F_H_FORM"`
	Alpha     units.Percentage `xml:"F_H_ALPHA"`
	Beta      units.Percentage `xml:"F_H_BETA"`
	Percent   units.Percentage `xml:"F_H_PERCENT"`
	Inventory units.Weight     `xml:"F_H_INVENTORY"`
	Amount    units.Weight     `xml:"F_H_AMOUNT"`
	// Hop storage index provices percent of alpha acid lost in 6 months.
	HSI                  units.Percentage  `xml:"F_H_HSI"`
	BoilTime             units.Duration    `xml:"F_H_BOIL_TIME"`
	DryHopTime           units.Duration    `xml:"F_H_DRY_HOP_TIME"`
	Notes                string            `xml:"F_H_NOTES"`
	WhirlpoolTemperature units.Temperature `xml:"F_H_WHIRLPOOL_TEMP"`
	IBUContrib           units.Percentage  `xml:"F_H_IBU_CONTRIB"`
	Order                int               `xml:"F_ORDER"`
	Use                  HopUse            `xml:"F_H_USE"`
	InRecipe             units.Bool        `xml:"F_H_IN_RECIPE"`
	Price                units.Price       `xml:"F_H_PRICE"`
}

type HopForm int

var (
	HopFormPellet               HopForm = 0
	HopFormPlug                 HopForm = 1
	HopFormLeaf                 HopForm = 2
	HopFormHopEctractCO2        HopForm = 3
	HopFormHopEctractIsomerized HopForm = 4
)

type HopType int

var (
	HopTypeBittering HopType = 0
	HopTypeAroma     HopType = 1
	HopTypeBOth      HopType = 2
)

type HopUse int

var (
	HopUseBoil      HopUse = 0
	HopUseDryHop    HopUse = 1
	HopUseMash      HopUse = 2
	HopUseFirstWort HopUse = 3
	HopUseWhirlpool HopUse = 4
)
