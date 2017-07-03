package bsmx

import "github.com/leonb/go-beersmith/units"

type Yeast struct {
	Order          int        `xml:"F_ORDER"`
	AddToSecondary units.Bool `xml:"F_Y_ADD_TO_SECONDARY"`
	// The percent of initial yeast cells lost each month as the yeast ages.
	AgeRate        units.Percentage  `xml:"F_Y_AGE_RATE"`
	Amount         units.Amount      `xml:"F_Y_AMOUNT"`
	BestFor        string            `xml:"F_Y_BEST_FOR"`
	BrewDate       units.Date        `xml:"F_Y_BREW_DATE"`
	Cells          float64           `xml:"F_Y_CELLS"`
	CultureDate    units.Date        `xml:"F_Y_CULTURE_DATE"`
	Flocculation   YeastFlocculation `xml:"F_Y_FLOCCULATION"`
	Form           YeastForm         `xml:"F_Y_FORM"`
	Inventory      units.Amount      `xml:"F_Y_INVENTORY"`
	InRecipe       units.Bool        `xml:"F_Y_IN_RECIPE"`
	Lab            string            `xml:"F_Y_LAB"`
	MaxAttenuation units.Percentage  `xml:"F_Y_MAX_ATTENUATION"`
	MaxReuse       int               `xml:"F_Y_MAX_REUSE"`
	MaxTemp        units.Temperature `xml:"F_Y_MAX_TEMP"`
	MinAttenuation units.Percentage  `xml:"F_Y_MIN_ATTENUATION"`
	MinTemp        units.Temperature `xml:"F_Y_MIN_TEMP"`
	Name           string            `xml:"F_Y_NAME"`
	Notes          string            `xml:"F_Y_NOTES"`
	PkgDate        units.Date        `xml:"F_Y_PKG_DATE"`
	Price          units.Price       `xml:"F_Y_PRICE"`
	ProductId      string            `xml:"F_Y_PRODUCT_ID"`
	StarterSize    float64           `xml:"F_Y_STARTER_SIZE"`
	TimesCultured  int               `xml:"F_Y_TIMES_CULTURED"`
	Type           YeastType         `xml:"F_Y_TYPE"`
	UseStarter     units.Bool        `xml:"F_Y_USE_STARTER"`
	_MOD_          units.Date        `xml:"_MOD_"`
}

type YeastFlocculation int

var (
	YeastFlocculationLow      YeastFlocculation = 0
	YeastFlocculationMedium   YeastFlocculation = 1
	YeastFlocculationHigh     YeastFlocculation = 2
	YeastFlocculationVeryHigh YeastFlocculation = 3
)

type YeastForm int

var (
	YeastFormLiquid  YeastForm = 0
	YeastFormDry     YeastForm = 1
	YeastFormSlant   YeastForm = 2
	YeastFormCulture YeastForm = 3
)

type YeastType int

var (
	YeastTypeAle       YeastType = 0
	YeastTypeLager     YeastType = 1
	YeastTypeWine      YeastType = 2
	YeastTypeChampagne YeastType = 3
	YeastTypeWheat     YeastType = 4
)
