package bsmx

import "github.com/leonb/go-beersmith/units"

type Hop struct {
	F_H_ALPHA        float64     `xml:"F_H_ALPHA"`
	F_H_AMOUNT       float64     `xml:"F_H_AMOUNT"`
	F_H_BETA         float64     `xml:"F_H_BETA"`
	F_H_BOIL_TIME    float64     `xml:"F_H_BOIL_TIME"`
	F_H_DRY_HOP_TIME float64     `xml:"F_H_DRY_HOP_TIME"`
	F_H_FORM         HopForm     `xml:"F_H_FORM"`
	F_H_HSI          float64     `xml:"F_H_HSI"`
	F_H_IBU_CONTRIB  float64     `xml:"F_H_IBU_CONTRIB"`
	F_H_INVENTORY    float64     `xml:"F_H_INVENTORY"`
	F_H_IN_RECIPE    units.Bool  `xml:"F_H_IN_RECIPE"`
	F_H_NAME         string      `xml:"F_H_NAME"`
	F_H_NOTES        string      `xml:"F_H_NOTES"`
	F_H_ORIGIN       string      `xml:"F_H_ORIGIN"`
	F_H_PERCENT      float64     `xml:"F_H_PERCENT"`
	F_H_PRICE        units.Price `xml:"F_H_PRICE"`
	F_H_TYPE         HopType     `xml:"F_H_TYPE"`
	F_H_USE          HopUse      `xml:"F_H_USE"`
	F_ORDER          int         `xml:"F_ORDER"`
	_MOD_            units.Date  `xml:"_MOD_"`
}

type HopForm int

type HopType int

type HopUse int
