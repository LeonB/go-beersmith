package bsmx

import "github.com/leonb/go-beersmith/units"

type MashProfiles []MashProfile

type MashProfile struct {
	F_MASH_39    units.Bool                    `xml:"F_MASH_39"`
	Batch        units.Bool                    `xml:"F_MH_BATCH"`
	BatchDrain   units.Bool                    `xml:"F_MH_BATCH_DRAIN"`
	BatchEven    units.Bool                    `xml:"F_MH_BATCH_EVEN"`
	BatchPct     units.Percentage              `xml:"F_MH_BATCH_PCT"`
	Biab         units.Bool                    `xml:"F_MH_BIAB"`
	BiabVol      units.Volume                  `xml:"F_MH_BIAB_VOL"`
	BoilTemp     units.Temperature             `xml:"F_MH_BOIL_TEMP"`
	EquipAdjust  units.Bool                    `xml:"F_MH_EQUIP_ADJUST"`
	GrainTemp    units.Temperature             `xml:"F_MH_GRAIN_TEMP"`
	GrainWeight  units.Weight                  `xml:"F_MH_GRAIN_WEIGHT"`
	Name         string                        `xml:"F_MH_NAME"`
	Notes        string                        `xml:"F_MH_NOTES"`
	PH           units.PH                      `xml:"F_MH_PH"`
	SpargeTemp   units.Temperature             `xml:"F_MH_SPARGE_TEMP"`
	TunDeadspace units.Volume                  `xml:"F_MH_TUN_DEADSPACE"`
	TunHC        units.HeatTransferCoefficient `xml:"F_MH_TUN_HC"`
	TunMass      units.Mass                    `xml:"F_MH_TUN_MASS"`
	TunTemp      units.Temperature             `xml:"F_MH_TUN_TEMP"`
	TunVol       units.Volume                  `xml:"F_MH_TUN_VOL"`
	MashSteps    MashSteps                     `xml:"steps"`
	_MOD_        units.Date                    `xml:"_MOD_"`
}
