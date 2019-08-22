package bsmx

import (
	"github.com/leonb/go-beersmith/units"
)

type Grains []Grain

type Grain struct {
	MOD            units.Date       `xml:"_MOD_"`
	Name           String           `xml:"F_G_NAME"`
	Origin         string           `xml:"F_G_ORIGIN"`
	Supplier       string           `xml:"F_G_SUPPLIER"`
	Type           GrainType        `xml:"F_G_TYPE"`
	Use            GrainUse         `xml:"F_G_USE"`
	UseSet         int              `xml:"F_G_USE_SET"`
	AcidPercentage units.Percentage `xml:"F_G_ACID_PCT"`
	InRecipe       units.Bool       `xml:"F_G_IN_RECIPE"`
	Inventory      units.Weight     `xml:"F_G_INVENTORY"`
	Amount         units.Weight     `xml:"F_G_AMOUNT"`
	Color          units.Color      `xml:"F_G_COLOR"`
	// Fine crush, dry grain yield of the malt
	Yield units.Percentage `xml:"F_G_YIELD"`
	// Late Extract Boil Time
	LateExtract units.Duration `xml:"F_G_LATE_EXTRACT"`
	// Percentage in recipe
	Percent        units.Percentage `xml:"F_G_PERCENT"`
	NotFermentable units.Bool       `xml:"F_G_NOT_FERMENTABLE"`
	Order          int              `xml:"F_ORDER"`
	// Difference between the coarse and fine grain yield in percent.
	CoarseFineDiff units.Percentage `xml:"F_G_COARSE_FINE_DIFF"`
	// Moisture content for the grain in percent.
	Moisture units.Percentage `xml:"F_G_MOISTURE"`
	// Diastic power is a measure of how much enzymes the grain contributes for
	// converting sugars during the mash.
	DiastaticPower units.Lintner `xml:"F_G_DIASTATIC_POWER"`
	// Protein content of the grain in percent
	Protein units.Percentage `xml:"F_G_PROTEIN"`
	// Used for DME/LME: IBU's per Gallon Per Pound?
	IBUGalPerLb   units.Unknown    `xml:"F_G_IBU_GAL_PER_LB"`
	AddAfterBoil  units.Bool       `xml:"F_G_ADD_AFTER_BOIL"`
	RecommendMash units.Bool       `xml:"F_G_RECOMMEND_MASH"`
	MaxInBatch    units.Percentage `xml:"F_G_MAX_IN_BATCH"`
	Notes         String           `xml:"F_G_NOTES"`
	BoilTime      units.Duration   `xml:"F_G_BOIL_TIME"`
	Price         units.Price      `xml:"F_G_PRICE"`
	// Replace this grain with the following grain when converting recipes to
	// extract.
	ConvertGrain string `xml:"F_G_CONVERT_GRAIN"`
}

type GrainType int

const (
	GrainTypeGrain      GrainType = 0
	GrainTypeExtract    GrainType = 1
	GrainTypeSugar      GrainType = 2
	GrainTypeAdjunct    GrainType = 3
	GrainTypeDryExtract GrainType = 4
	GrainTypeFruit      GrainType = 5
	GrainTypeJuice      GrainType = 6
	GrainTypeHoney      GrainType = 7
)

func (gt GrainType) String() string {
	switch gt {
	case 0:
		return "Grain"
	case 1:
		return "Extract"
	case 2:
		return "Sugar"
	case 3:
		return "Adjunct"
	case 4:
		return "DryExtract"
	case 5:
		return "Fruit"
	case 6:
		return "Juice"
	case 7:
		return "Honey"
	}

	return ""
}

type GrainUse int

const (
	GrainUseMash      GrainUse = 0
	GrainUseSteep     GrainUse = 1
	GrainUseBoil      GrainUse = 2
	GrainUseWhirlpool GrainUse = 3
	GrainUsePrimary   GrainUse = 4
	GrainUseSeconday  GrainUse = 5
	GrainUseTertiary  GrainUse = 6
	GrainUseBottling  GrainUse = 7
)

func (gu GrainUse) String() string {
	switch gu {
	case 0:
		return "Mash"
	case 1:
		return "Steep"
	case 2:
		return "Boil"
	case 3:
		return "Whirlpool"
	case 4:
		return "Primary"
	case 5:
		return "Secondary"
	case 6:
		return "Tertiary"
	case 7:
		return "Bottling"
	}

	return ""
}
