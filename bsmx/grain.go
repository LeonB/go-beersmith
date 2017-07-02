package bsmx

import "github.com/leonb/go-beersmith/units"

type Grain struct {
	F_G_ADD_AFTER_BOIL   units.Bool  `xml:"F_G_ADD_AFTER_BOIL"`
	F_G_AMOUNT           float64     `xml:"F_G_AMOUNT"`
	F_G_BOIL_TIME        float64     `xml:"F_G_BOIL_TIME"`
	F_G_COARSE_FINE_DIFF float64     `xml:"F_G_COARSE_FINE_DIFF"`
	F_G_COLOR            float64     `xml:"F_G_COLOR"`
	F_G_CONVERT_GRAIN    string      `xml:"F_G_CONVERT_GRAIN"`
	F_G_DIASTATIC_POWER  float64     `xml:"F_G_DIASTATIC_POWER"`
	F_G_IBU_GAL_PER_LB   float64     `xml:"F_G_IBU_GAL_PER_LB"`
	F_G_INVENTORY        float64     `xml:"F_G_INVENTORY"`
	F_G_IN_RECIPE        units.Bool  `xml:"F_G_IN_RECIPE"`
	F_G_LATE_EXTRACT     float64     `xml:"F_G_LATE_EXTRACT"`
	F_G_MAX_IN_BATCH     float64     `xml:"F_G_MAX_IN_BATCH"`
	F_G_MOISTURE         float64     `xml:"F_G_MOISTURE"`
	F_G_NAME             string      `xml:"F_G_NAME"`
	F_G_NOTES            string      `xml:"F_G_NOTES"`
	F_G_NOT_FERMENTABLE  units.Bool  `xml:"F_G_NOT_FERMENTABLE"`
	F_G_ORIGIN           string      `xml:"F_G_ORIGIN"`
	F_G_PERCENT          float64     `xml:"F_G_PERCENT"`
	F_G_PRICE            units.Price `xml:"F_G_PRICE"`
	F_G_PROTEIN          float64     `xml:"F_G_PROTEIN"`
	F_G_RECOMMEND_MASH   units.Bool  `xml:"F_G_RECOMMEND_MASH"`
	F_G_SUPPLIER         string      `xml:"F_G_SUPPLIER"`
	F_G_TYPE             GrainType   `xml:"F_G_TYPE"`
	F_G_YIELD            float64     `xml:"F_G_YIELD"`
	F_ORDER              int         `xml:"F_ORDER"`
	_MOD_                units.Date  `xml:"_MOD_"`
}

type GrainType int
