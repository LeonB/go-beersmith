package bsmx

import "github.com/leonb/go-beersmith/units"

type Misc struct {
	F_M_AMOUNT           float64      `xml:"F_M_AMOUNT"`
	F_M_IMPORT_AS_WEIGHT units.Bool   `xml:"F_M_IMPORT_AS_WEIGHT"`
	F_M_IMPORT_UNITS     units.Bool   `xml:"F_M_IMPORT_UNITS"`
	F_M_INVENTORY        float64      `xml:"F_M_INVENTORY"`
	F_M_NAME             string       `xml:"F_M_NAME"`
	F_M_NOTES            string       `xml:"F_M_NOTES"`
	F_M_PRICE            units.Price  `xml:"F_M_PRICE"`
	F_M_TEMP_VOL         units.Volume `xml:"F_M_TEMP_VOL"`
	F_M_TIME             float64      `xml:"F_M_TIME"`
	F_M_TIME_UNITS       TimeUnit     `xml:"F_M_TIME_UNITS"`
	F_M_TYPE             MiscType     `xml:"F_M_TYPE"`
	F_M_UNITS            MiscUnit     `xml:"F_M_UNITS"`
	F_M_USE              MiscUse      `xml:"F_M_USE"`
	F_M_USE_FOR          MiscUseFor   `xml:"F_M_USE_FOR"`
	F_M_VOLUME           units.Volume `xml:"F_M_VOLUME"`
	F_ORDER              int          `xml:"F_ORDER"`
	_MOD_                units.Date   `xml:"_MOD_"`
}

type MiscType int

type MiscUnit int

type MiscUse int

type MiscUseFor string
