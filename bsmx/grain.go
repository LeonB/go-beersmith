package bsmx

import "github.com/leonb/go-beersmith/units"

type Grain struct {
	AddAfterBoil units.Bool     `xml:"F_G_ADD_AFTER_BOIL"`
	Amount       units.Weight   `xml:"F_G_AMOUNT"`
	BoilTime     units.Duration `xml:"F_G_BOIL_TIME"`
	// Difference between the coarse and fine grain yield in percent.
	CoarseFineDiff units.Percentage `xml:"F_G_COARSE_FINE_DIFF"`
	Color          units.Color      `xml:"F_G_COLOR"`
	// Replace this grain with the following grain when converting recipes to
	// extract.
	ConvertGrain string `xml:"F_G_CONVERT_GRAIN"`
	// Diastic power is a measure of how much enzymes the grain contributes for
	// converting sugars during the mash.
	DiastaticPower units.Lintner    `xml:"F_G_DIASTATIC_POWER"`
	IBUGalPerLb    units.Unknown    `xml:"F_G_IBU_GAL_PER_LB"`
	Inventory      units.Weight     `xml:"F_G_INVENTORY"`
	InRecipe       units.Bool       `xml:"F_G_IN_RECIPE"`
	LateExtract    float64          `xml:"F_G_LATE_EXTRACT"`
	MaxInBatch     units.Percentage `xml:"F_G_MAX_IN_BATCH"`
	// Moisture content for the grain in percent.
	Moisture       units.Percentage `xml:"F_G_MOISTURE"`
	Name           string           `xml:"F_G_NAME"`
	Notes          string           `xml:"F_G_NOTES"`
	NotFermentable units.Bool       `xml:"F_G_NOT_FERMENTABLE"`
	Origin         string           `xml:"F_G_ORIGIN"`
	Percent        units.Percentage `xml:"F_G_PERCENT"`
	Price          units.Price      `xml:"F_G_PRICE"`
	// Protein content of the grain in percent
	Protein       units.Percentage `xml:"F_G_PROTEIN"`
	RecommendMash units.Bool       `xml:"F_G_RECOMMEND_MASH"`
	Supplier      string           `xml:"F_G_SUPPLIER"`
	Type          GrainType        `xml:"F_G_TYPE"`
	// Fine crush, dry grain yield of the malt
	Yield units.Percentage `xml:"F_G_YIELD"`
	Order int              `xml:"F_ORDER"`
	_MOD_ units.Date       `xml:"_MOD_"`
}

type GrainType int

var (
	GrainTypeGrain      GrainType = 0
	GrainTypeExtract    GrainType = 1
	GrainTypeSugar      GrainType = 2
	GrainTypeAdjunct    GrainType = 3
	GrainTypeDryExtract GrainType = 4
)
