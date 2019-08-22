package bsmx

import (
	"encoding/xml"

	"github.com/leonb/go-beersmith/units"
)

type EquipmentProfiles []EquipmentProfile

type EquipmentProfile struct {
	MOD             units.Date           `xml:"_MOD_"`
	Name            string               `xml:"F_E_NAME"`
	Type            EquipmentProfileType `xml:"F_E_TYPE"`
	ShowBoil        IntBool              `xml:"F_E_SHOW_BOIL"`
	MashVol         units.Volume         `xml:"F_E_MASH_VOL"`
	TunMass         units.Mass           `xml:"F_E_TUN_MASS"`
	BoilRateFlag    units.Bool           `xml:"F_E_BOIL_RATE_FLAG"`
	TunSpecificHeat units.SpecificHeat   `xml:"F_E_TUN_SPECIFIC_HEAT"`
	TunDeadspace    units.Volume         `xml:"F_E_TUN_DEADSPACE"`
	TunAddition     units.Volume         `xml:"F_E_TUN_ADDITION"`
	TunAdjDeadspace units.Volume2        `xml:"F_E_TUN_ADJ_DEADSPACE"`
	CalcBoil        units.Bool           `xml:"F_E_CALC_BOIL"`
	BoilVol         units.Volume         `xml:"F_E_BOIL_VOL"`
	BoilTime        units.Duration       `xml:"F_E_BOIL_TIME"`
	OldEvapRate     units.Percentage     `xml:"F_E_OLD_EVAP_RATE"`
	F_EQUIP_39      units.Bool           `xml:"F_EQUIP_39"`
	BoilOff         units.Volume         `xml:"F_E_BOIL_OFF"`
	TrubLoss        units.Volume         `xml:"F_E_TRUB_LOSS"`
	CoolPct         units.Percentage     `xml:"F_E_COOL_PCT"`
	TopUpKettle     units.Volume         `xml:"F_E_TOP_UP_KETTLE"`
	BatchVol        units.Volume         `xml:"F_E_BATCH_VOL"`
	FermenterLoss   units.Volume         `xml:"F_E_FERMENTER_LOSS"`
	TopUp           units.Volume         `xml:"F_E_TOP_UP"`
	Efficiency      units.Percentage     `xml:"F_E_EFFICIENCY"`
	HopUtil         units.Percentage     `xml:"F_E_HOP_UTIL"`
	Notes           String               `xml:"F_E_NOTES"`
	Altitude        units.Altitude       `xml:"F_E_ALTITUDE"`
	WhirlpoolCarry  units.Bool           `xml:"F_E_WHIRLPOOL_CARRY"`
	WhirlpoolTime   units.Duration       `xml:"F_E_WHIRLPOOL_TIME"`
	_MOD_           units.Date           `xml:"_MOD_"`
}

func (ep EquipmentProfile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias EquipmentProfile
	a := alias(ep)
	a.WhirlpoolTime.SetPrintf("%g")
	return e.EncodeElement(a, start)
}

type EquipmentProfileType int

const (
	EquipmentProfileTypeGeneral     EquipmentProfileType = 0
	EquipmentProfileTypeExtract     EquipmentProfileType = 1
	EquipmentProfileTypePartialMash EquipmentProfileType = 2
	EquipmentProfileTypeAllGrain    EquipmentProfileType = 3
	EquipmentProfileTypeCider       EquipmentProfileType = 4
	EquipmentProfileTypeMead        EquipmentProfileType = 5
	EquipmentProfileTypeWine        EquipmentProfileType = 6
)

func (ept EquipmentProfileType) String() string {
	switch ept {
	case 0:
		return "General"
	case 1:
		return "Extract"
	case 2:
		return "Partial Mash"
	case 3:
		return "All Grain"
	case 4:
		return "Cider"
	case 5:
		return "Mead"
	case 6:
		return "Wine"
	}

	return ""
}
