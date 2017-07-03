package bsmx

import "github.com/leonb/go-beersmith/units"

type Equipment struct {
	F_EQUIP_39      units.Bool         `xml:"F_EQUIP_39"`
	BatchVol        units.Volume       `xml:"F_E_BATCH_VOL"`
	BoilOff         units.Volume       `xml:"F_E_BOIL_OFF"`
	BoilRateFlag    units.Bool         `xml:"F_E_BOIL_RATE_FLAG"`
	BoilTime        units.Duration     `xml:"F_E_BOIL_TIME"`
	BoilVol         units.Volume       `xml:"F_E_BOIL_VOL"`
	CalcBoil        units.Bool         `xml:"F_E_CALC_BOIL"`
	CoolPct         units.Percentage   `xml:"F_E_COOL_PCT"`
	Efficiency      units.Percentage   `xml:"F_E_EFFICIENCY"`
	FermenterLoss   units.Volume       `xml:"F_E_FERMENTER_LOSS"`
	HopUtil         units.Percentage   `xml:"F_E_HOP_UTIL"`
	MashVol         units.Volume       `xml:"F_E_MASH_VOL"`
	Name            string             `xml:"F_E_NAME"`
	Notes           string             `xml:"F_E_NOTES"`
	OldEvapRate     units.Percentage   `xml:"F_E_OLD_EVAP_RATE"`
	TopUp           units.Volume       `xml:"F_E_TOP_UP"`
	TopUpKettle     units.Volume       `xml:"F_E_TOP_UP_KETTLE"`
	TrubLoss        units.Volume       `xml:"F_E_TRUB_LOSS"`
	TunAddition     units.Volume       `xml:"F_E_TUN_ADDITION"`
	TunAdjDeadspace units.Volume       `xml:"F_E_TUN_ADJ_DEADSPACE"`
	TunDeadspace    units.Volume       `xml:"F_E_TUN_DEADSPACE"`
	TunMass         units.Mass         `xml:"F_E_TUN_MASS"`
	TunSpecificHeat units.SpecificHeat `xml:"F_E_TUN_SPECIFIC_HEAT"`
	WhirlpoolCarry  units.Bool         `xml:"F_E_WHIRLPOOL_CARRY"`
	WhirlpoolTime   units.Duration     `xml:"F_E_WHIRLPOOL_TIME"`
	_MOD_           units.Date         `xml:"_MOD_"`
}
