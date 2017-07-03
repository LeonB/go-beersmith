package bsmx

import "github.com/leonb/go-beersmith/units"

type AgingProfile struct {
	Age         units.Days        `xml:"F_A_AGE"`
	AgeTemp     units.Temperature `xml:"F_A_AGE_TEMP"`
	EndAgeTemp  units.Temperature `xml:"F_A_END_AGE_TEMP"`
	EndTempsSet units.Bool        `xml:"F_A_END_TEMPS_SET"`
	Name        string            `xml:"F_A_NAME"`
	Notes       string            `xml:"F_A_NOTES"`
	PrimDays    units.Days        `xml:"F_A_PRIM_DAYS"`
	PrimEndTemp units.Temperature `xml:"F_A_PRIM_END_TEMP"`
	PrimTemp    units.Temperature `xml:"F_A_PRIM_TEMP"`
	SecDays     units.Days        `xml:"F_A_SEC_DAYS"`
	SecEndTemp  units.Temperature `xml:"F_A_SEC_END_TEMP"`
	SecTemp     units.Temperature `xml:"F_A_SEC_TEMP"`
	TertDays    units.Days        `xml:"F_A_TERT_DAYS"`
	TertEndTemp units.Temperature `xml:"F_A_TERT_END_TEMP"`
	TertTemp    units.Temperature `xml:"F_A_TERT_TEMP"`
	Type        AgingProfileType  `xml:"F_A_TYPE"`
	_MOD_       units.Date        `xml:"_MOD_"`
}

type AgingProfileType int
