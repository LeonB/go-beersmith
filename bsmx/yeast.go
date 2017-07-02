package bsmx

import "github.com/leonb/go-beersmith/units"

type Yeast struct {
	F_ORDER              int               `xml:"F_ORDER"`
	F_Y_ADD_TO_SECONDARY units.Bool        `xml:"F_Y_ADD_TO_SECONDARY"`
	F_Y_AGE_RATE         float64           `xml:"F_Y_AGE_RATE"`
	F_Y_AMOUNT           float64           `xml:"F_Y_AMOUNT"`
	F_Y_BEST_FOR         string            `xml:"F_Y_BEST_FOR"`
	F_Y_BREW_DATE        units.Date        `xml:"F_Y_BREW_DATE"`
	F_Y_CELLS            float64           `xml:"F_Y_CELLS"`
	F_Y_CULTURE_DATE     units.Date        `xml:"F_Y_CULTURE_DATE"`
	F_Y_FLOCCULATION     YeastFlocculation `xml:"F_Y_FLOCCULATION"`
	F_Y_FORM             YeastForm         `xml:"F_Y_FORM"`
	F_Y_INVENTORY        float64           `xml:"F_Y_INVENTORY"`
	F_Y_IN_RECIPE        units.Bool        `xml:"F_Y_IN_RECIPE"`
	F_Y_LAB              string            `xml:"F_Y_LAB"`
	F_Y_MAX_ATTENUATION  float64           `xml:"F_Y_MAX_ATTENUATION"`
	F_Y_MAX_REUSE        int               `xml:"F_Y_MAX_REUSE"`
	F_Y_MAX_TEMP         float64           `xml:"F_Y_MAX_TEMP"`
	F_Y_MIN_ATTENUATION  float64           `xml:"F_Y_MIN_ATTENUATION"`
	F_Y_MIN_TEMP         float64           `xml:"F_Y_MIN_TEMP"`
	F_Y_NAME             string            `xml:"F_Y_NAME"`
	F_Y_NOTES            string            `xml:"F_Y_NOTES"`
	F_Y_PKG_DATE         units.Date        `xml:"F_Y_PKG_DATE"`
	F_Y_PRICE            units.Price       `xml:"F_Y_PRICE"`
	F_Y_PRODUCT_ID       string            `xml:"F_Y_PRODUCT_ID"`
	F_Y_STARTER_SIZE     float64           `xml:"F_Y_STARTER_SIZE"`
	F_Y_TIMES_CULTURED   int               `xml:"F_Y_TIMES_CULTURED"`
	F_Y_TYPE             YeastType         `xml:"F_Y_TYPE"`
	F_Y_USE_STARTER      units.Bool        `xml:"F_Y_USE_STARTER"`
	_MOD_                units.Date        `xml:"_MOD_"`
}

type YeastFlocculation int

type YeastForm int

type YeastType int
