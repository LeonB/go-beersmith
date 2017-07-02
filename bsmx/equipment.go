package bsmx

import "github.com/leonb/go-beersmith/units"

type Equipment struct {
	F_EQUIP_39            units.Bool       `xml:"F_EQUIP_39"`
	F_E_BATCH_VOL         units.Volume     `xml:"F_E_BATCH_VOL"`
	F_E_BOIL_OFF          units.Volume     `xml:"F_E_BOIL_OFF"`
	F_E_BOIL_RATE_FLAG    units.Bool       `xml:"F_E_BOIL_RATE_FLAG"`
	F_E_BOIL_TIME         units.Duration   `xml:"F_E_BOIL_TIME"`
	F_E_BOIL_VOL          units.Volume     `xml:"F_E_BOIL_VOL"`
	F_E_CALC_BOIL         units.Bool       `xml:"F_E_CALC_BOIL"`
	F_E_COOL_PCT          units.Percentage `xml:"F_E_COOL_PCT"`
	F_E_EFFICIENCY        units.Percentage `xml:"F_E_EFFICIENCY"`
	F_E_FERMENTER_LOSS    units.Volume     `xml:"F_E_FERMENTER_LOSS"`
	F_E_HOP_UTIL          units.Percentage `xml:"F_E_HOP_UTIL"`
	F_E_MASH_VOL          units.Volume     `xml:"F_E_MASH_VOL"`
	F_E_NAME              string           `xml:"F_E_NAME"`
	F_E_NOTES             string           `xml:"F_E_NOTES"`
	F_E_OLD_EVAP_RATE     units.Percentage `xml:"F_E_OLD_EVAP_RATE"`
	F_E_TOP_UP            units.Volume     `xml:"F_E_TOP_UP"`
	F_E_TOP_UP_KETTLE     units.Volume     `xml:"F_E_TOP_UP_KETTLE"`
	F_E_TRUB_LOSS         units.Volume     `xml:"F_E_TRUB_LOSS"`
	F_E_TUN_ADDITION      units.Volume     `xml:"F_E_TUN_ADDITION"`
	F_E_TUN_ADJ_DEADSPACE units.Volume     `xml:"F_E_TUN_ADJ_DEADSPACE"`
	F_E_TUN_DEADSPACE     units.Volume     `xml:"F_E_TUN_DEADSPACE"`
	F_E_TUN_MASS          units.Mass       `xml:"F_E_TUN_MASS"`
	F_E_TUN_SPECIFIC_HEAT float64          `xml:"F_E_TUN_SPECIFIC_HEAT"`
	F_E_WHIRLPOOL_CARRY   units.Bool       `xml:"F_E_WHIRLPOOL_CARRY"`
	F_E_WHIRLPOOL_TIME    units.Duration   `xml:"F_E_WHIRLPOOL_TIME"`
	_MOD_                 units.Date       `xml:"_MOD_"`
}
