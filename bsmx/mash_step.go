package bsmx

import "github.com/leonb/go-beersmith/units"

type MashStep struct {
	// Calculated volume of mash to decoct. Only applicable for a decoction step. Includes the units as in "7.5 l" or "2.3 gal"
	DecoctionAmt units.Volume                  `xml:"F_MS_DECOCTION_AMT,omitempty"`
	GrainTemp    units.Temperature             `xml:"F_MS_GRAIN_TEMP,omitempty"`
	GrainWeight  units.Weight                  `xml:"F_MS_GRAIN_WEIGHT,omitempty"`
	Infusion     units.Unknown                 `xml:"F_MS_INFUSION,omitempty"`
	InfusionTemp units.Temperature             `xml:"F_MS_INFUSION_TEMP,omitempty"`
	Name         string                        `xml:"F_MS_NAME,omitempty"`
	RiseTime     units.Duration                `xml:"F_MS_RISE_TIME,omitempty"`
	StartTemp    units.Temperature             `xml:"F_MS_START_TEMP,omitempty"`
	StartVol     units.Volume                  `xml:"F_MS_START_VOL,omitempty"`
	StepTemp     units.Temperature             `xml:"F_MS_STEP_TEMP,omitempty"`
	StepTime     units.Duration                `xml:"F_MS_STEP_TIME,omitempty"`
	TunAddition  units.Volume                  `xml:"F_MS_TUN_ADDITION,omitempty"`
	TunHC        units.HeatTransferCoefficient `xml:"F_MS_TUN_HC,omitempty"`
	TunMass      units.Mass                    `xml:"F_MS_TUN_MASS,omitempty"`
	TunTemp      units.Temperature             `xml:"F_MS_TUN_TEMP,omitempty"`
	TunVol       units.Volume                  `xml:"F_MS_TUN_VOL,omitempty"`
	Type         MashStepType                  `xml:"F_MS_TYPE,omitempty"`
	_MOD_        units.Date                    `xml:"_MOD_,omitempty"`
}

type MashStepType int

var (
	MashStepTypeInfusion    MashStepType = 0
	MashStepTypeDecoction   MashStepType = 1
	MashStepTypeTemperature MashStepType = 3
)
