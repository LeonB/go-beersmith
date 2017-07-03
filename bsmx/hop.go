package bsmx

import "github.com/leonb/go-beersmith/units"

type Hop struct {
	Alpha      units.Percentage `xml:"F_H_ALPHA"`
	Amount     units.Weight     `xml:"F_H_AMOUNT"`
	Beta       units.Percentage `xml:"F_H_BETA"`
	BoilTime   units.Duration   `xml:"F_H_BOIL_TIME"`
	DryHopTime units.Duration   `xml:"F_H_DRY_HOP_TIME"`
	Form       HopForm          `xml:"F_H_FORM"`
	// Hop storage index provices percent of alpha acid lost in 6 months.
	HSI        units.Percentage `xml:"F_H_HSI"`
	IBUContrib units.Percentage `xml:"F_H_IBU_CONTRIB"`
	Inventory  units.Weight     `xml:"F_H_INVENTORY"`
	InRecipe   units.Bool       `xml:"F_H_IN_RECIPE"`
	Name       string           `xml:"F_H_NAME"`
	Notes      string           `xml:"F_H_NOTES"`
	Origin     string           `xml:"F_H_ORIGIN"`
	Percent    units.Percentage `xml:"F_H_PERCENT"`
	Price      units.Price      `xml:"F_H_PRICE"`
	Type       HopType          `xml:"F_H_TYPE"`
	Use        HopUse           `xml:"F_H_USE"`
	Order      int              `xml:"F_ORDER"`
	_MOD_      units.Date       `xml:"_MOD_"`
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
