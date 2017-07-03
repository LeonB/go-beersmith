package bsmx

import "github.com/leonb/go-beersmith/units"

type FermentationReading struct {
	Date    units.Time        `xml:"Date"`
	Gravity units.Gravity     `xml:"F_AD_GRAVITY"`
	Temp    units.Temperature `xml:"F_AD_TEMP"`
	Type    AgeDateType       `xml:"F_AD_TYPE"`
}

type AgeDateType int

var (
	AgeDateTypeBoth        AgeDateType = 0
	AgeDateTypeTemperature AgeDateType = 1
	AgeDateTypeGravity     AgeDateType = 2
)
