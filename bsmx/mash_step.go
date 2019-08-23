package bsmx

import "github.com/leonb/go-beersmith/units"

// type MashSteps []MashStep

type MashSteps struct {
	MOD       units.Date `xml:"_MOD_"`
	Name      String     `xml:"Name"`
	Type      int        `xml:"Type"`
	Dirty     IntBool    `xml:"Dirty"`
	Owndata   int        `xml:"Owndata"`
	TID       int        `xml:"TID"`
	Size      int        `xml:"Size"`
	XName     String     `xml:"_XName"`
	Allocinc  int        `xml:"Allocinc"`
	MashStep  []MashStep `xml:"Data>MashStep"`
	TExpanded units.Bool `xml:"_TExpanded"`
	TExtra    IntBool    `xml:"TExtra"`
}

type MashStep struct {
	MOD            units.Date                    `xml:"_MOD_"`
	Name           string                        `xml:"F_MS_NAME"`
	Type           MashStepType                  `xml:"F_MS_TYPE"`
	InfusionVolume units.Volume                  `xml:"F_MS_INFUSION"`
	StepTemp       units.Temperature             `xml:"F_MS_STEP_TEMP"`
	StepTime       units.Duration                `xml:"F_MS_STEP_TIME"`
	RiseTime       units.Duration                `xml:"F_MS_RISE_TIME"`
	TunAddition    units.Volume                  `xml:"F_MS_TUN_ADDITION"`
	TunHC          units.HeatTransferCoefficient `xml:"F_MS_TUN_HC"`
	TunVol         units.Volume                  `xml:"F_MS_TUN_VOL"`
	TunTemp        units.Temperature             `xml:"F_MS_TUN_TEMP"`
	TunMass        units.Mass                    `xml:"F_MS_TUN_MASS"`
	StartTemp      units.Temperature             `xml:"F_MS_START_TEMP"`
	GrainTemp      units.Temperature             `xml:"F_MS_GRAIN_TEMP"`
	StartVol       units.Volume                  `xml:"F_MS_START_VOL"`
	GrainWeight    units.Weight                  `xml:"F_MS_GRAIN_WEIGHT"`
	InfusionTemp   units.Temperature             `xml:"F_MS_INFUSION_TEMP"`
	// Calculated volume of mash to decoct. Only applicable for a decoction step. Includes the units as in "7.5 l" or "2.3 gal"
	DecoctionAmt units.Volume `xml:"F_MS_DECOCTION_AMT"`
}

type MashStepType int

var (
	MashStepTypeInfusion    MashStepType = 0
	MashStepTypeDecoction   MashStepType = 1
	MashStepTypeTemperature MashStepType = 3
)
