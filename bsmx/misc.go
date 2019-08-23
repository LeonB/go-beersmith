package bsmx

import "github.com/leonb/go-beersmith/units"

type Miscs []Misc

type Misc struct {
	Amount         float64        `xml:"F_M_AMOUNT"`
	ImportAsWeight units.Bool     `xml:"F_M_IMPORT_AS_WEIGHT"`
	ImportUnits    units.Bool     `xml:"F_M_IMPORT_UNITS"`
	Inventory      float64        `xml:"F_M_INVENTORY"`
	Name           string         `xml:"F_M_NAME"`
	Notes          string         `xml:"F_M_NOTES"`
	Price          units.Price    `xml:"F_M_PRICE"`
	TempVol        units.Volume   `xml:"F_M_TEMP_VOL"`
	Time           units.Duration `xml:"F_M_TIME"`
	TimeUnits      TimeUnit       `xml:"F_M_TIME_UNITS"`
	Type           MiscType       `xml:"F_M_TYPE"`
	Units          MiscUnit       `xml:"F_M_UNITS"`
	Use            MiscUse        `xml:"F_M_USE"`
	UseFor         MiscUseFor     `xml:"F_M_USE_FOR"`
	Volume         units.Volume   `xml:"F_M_VOLUME"`
	Order          int            `xml:"F_ORDER"`
	MOD            units.Date     `xml:"_MOD_"`
}

type MiscType int

type MiscUnit int

type MiscUse int

type MiscUseFor string
