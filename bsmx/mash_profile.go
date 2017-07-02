package bsmx

import "github.com/leonb/go-beersmith/units"

type MashProfile struct {
	F_MASH_39          units.Bool        `xml:"F_MASH_39"`
	F_MH_BATCH         units.Bool        `xml:"F_MH_BATCH"`
	F_MH_BATCH_DRAIN   units.Bool        `xml:"F_MH_BATCH_DRAIN"`
	F_MH_BATCH_EVEN    units.Bool        `xml:"F_MH_BATCH_EVEN"`
	F_MH_BATCH_PCT     units.Percentage  `xml:"F_MH_BATCH_PCT"`
	F_MH_BIAB          units.Bool        `xml:"F_MH_BIAB"`
	F_MH_BIAB_VOL      units.Volume      `xml:"F_MH_BIAB_VOL"`
	F_MH_BOIL_TEMP     units.Temperature `xml:"F_MH_BOIL_TEMP"`
	F_MH_EQUIP_ADJUST  units.Bool        `xml:"F_MH_EQUIP_ADJUST"`
	F_MH_GRAIN_TEMP    units.Temperature `xml:"F_MH_GRAIN_TEMP"`
	F_MH_GRAIN_WEIGHT  units.Weight      `xml:"F_MH_GRAIN_WEIGHT"`
	F_MH_NAME          string            `xml:"F_MH_NAME"`
	F_MH_NOTES         string            `xml:"F_MH_NOTES"`
	F_MH_PH            units.PH          `xml:"F_MH_PH"`
	F_MH_SPARGE_TEMP   units.Temperature `xml:"F_MH_SPARGE_TEMP"`
	F_MH_TUN_DEADSPACE units.Volume      `xml:"F_MH_TUN_DEADSPACE"`
	F_MH_TUN_HC        units.Unknown     `xml:"F_MH_TUN_HC"`
	F_MH_TUN_MASS      units.Mass        `xml:"F_MH_TUN_MASS"`
	F_MH_TUN_TEMP      units.Temperature `xml:"F_MH_TUN_TEMP"`
	F_MH_TUN_VOL       units.Volume      `xml:"F_MH_TUN_VOL"`
	Steps              MashSteps         `xml:"steps"`
	_MOD_              units.Date        `xml:"_MOD_"`
}
