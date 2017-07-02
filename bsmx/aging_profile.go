package bsmx

import "github.com/leonb/go-beersmith/units"

type AgingProfile struct {
	F_A_AGE           units.Days        `xml:"F_A_AGE"`
	F_A_AGE_TEMP      units.Temperature `xml:"F_A_AGE_TEMP"`
	F_A_END_AGE_TEMP  units.Temperature `xml:"F_A_END_AGE_TEMP"`
	F_A_END_TEMPS_SET units.Bool        `xml:"F_A_END_TEMPS_SET"`
	F_A_NAME          string            `xml:"F_A_NAME"`
	F_A_NOTES         string            `xml:"F_A_NOTES"`
	F_A_PRIM_DAYS     units.Days        `xml:"F_A_PRIM_DAYS"`
	F_A_PRIM_END_TEMP units.Temperature `xml:"F_A_PRIM_END_TEMP"`
	F_A_PRIM_TEMP     units.Temperature `xml:"F_A_PRIM_TEMP"`
	F_A_SEC_DAYS      units.Days        `xml:"F_A_SEC_DAYS"`
	F_A_SEC_END_TEMP  units.Temperature `xml:"F_A_SEC_END_TEMP"`
	F_A_SEC_TEMP      units.Temperature `xml:"F_A_SEC_TEMP"`
	F_A_TERT_DAYS     units.Days        `xml:"F_A_TERT_DAYS"`
	F_A_TERT_END_TEMP units.Temperature `xml:"F_A_TERT_END_TEMP"`
	F_A_TERT_TEMP     units.Temperature `xml:"F_A_TERT_TEMP"`
	F_A_TYPE          AgingProfileType  `xml:"F_A_TYPE"`
	_MOD_             units.Date        `xml:"_MOD_"`
}

type AgingProfileType int
