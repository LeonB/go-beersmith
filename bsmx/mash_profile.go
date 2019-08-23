package bsmx

import "github.com/leonb/go-beersmith/units"

type MashProfiles []MashProfile

type MashProfile struct {
	MOD          units.Date                    `xml:"_MOD_"`
	Name         string                        `xml:"F_MH_NAME"`
	GrainWeight  units.Weight                  `xml:"F_MH_GRAIN_WEIGHT"`
	GrainTemp    units.Temperature             `xml:"F_MH_GRAIN_TEMP"`
	BoilTemp     units.Temperature             `xml:"F_MH_BOIL_TEMP"`
	TunTemp      units.Temperature             `xml:"F_MH_TUN_TEMP"`
	PH           units.PH                      `xml:"F_MH_PH"`
	SpargeTemp   units.Temperature             `xml:"F_MH_SPARGE_TEMP"`
	Batch        units.Bool                    `xml:"F_MH_BATCH"`
	BatchPct     units.Percentage              `xml:"F_MH_BATCH_PCT"`
	BatchEven    units.Bool                    `xml:"F_MH_BATCH_EVEN"`
	BatchDrain   units.Bool                    `xml:"F_MH_BATCH_DRAIN"`
	F_MASH_39    units.Bool                    `xml:"F_MASH_39"`
	TunDeadspace units.Volume                  `xml:"F_MH_TUN_DEADSPACE"`
	BiabVol      units.Volume                  `xml:"F_MH_BIAB_VOL"`
	Biab         units.Bool                    `xml:"F_MH_BIAB"`
	Notes        string                        `xml:"F_MH_NOTES"`
	MashSteps    MashSteps                     `xml:"steps"`
	EquipAdjust  units.Bool                    `xml:"F_MH_EQUIP_ADJUST"`
	TunMass      units.Mass                    `xml:"F_MH_TUN_MASS"`
	TunVol       units.Volume                  `xml:"F_MH_TUN_VOL"`
	TunHC        units.HeatTransferCoefficient `xml:"F_MH_TUN_HC"`
}
