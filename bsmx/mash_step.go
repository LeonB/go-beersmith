package bsmx

import "github.com/leonb/go-beersmith/units"

type MashStep struct {
	F_MS_DECOCTION_AMT float64      `xml:"F_MS_DECOCTION_AMT,omitempty"`
	F_MS_GRAIN_TEMP    float64      `xml:"F_MS_GRAIN_TEMP,omitempty"`
	F_MS_GRAIN_WEIGHT  float64      `xml:"F_MS_GRAIN_WEIGHT,omitempty"`
	F_MS_INFUSION      float64      `xml:"F_MS_INFUSION,omitempty"`
	F_MS_INFUSION_TEMP float64      `xml:"F_MS_INFUSION_TEMP,omitempty"`
	F_MS_NAME          string       `xml:"F_MS_NAME,omitempty"`
	F_MS_RISE_TIME     float64      `xml:"F_MS_RISE_TIME,omitempty"`
	F_MS_START_TEMP    float64      `xml:"F_MS_START_TEMP,omitempty"`
	F_MS_START_VOL     float64      `xml:"F_MS_START_VOL,omitempty"`
	F_MS_STEP_TEMP     float64      `xml:"F_MS_STEP_TEMP,omitempty"`
	F_MS_STEP_TIME     float64      `xml:"F_MS_STEP_TIME,omitempty"`
	F_MS_TUN_ADDITION  float64      `xml:"F_MS_TUN_ADDITION,omitempty"`
	F_MS_TUN_HC        float64      `xml:"F_MS_TUN_HC,omitempty"`
	F_MS_TUN_MASS      float64      `xml:"F_MS_TUN_MASS,omitempty"`
	F_MS_TUN_TEMP      float64      `xml:"F_MS_TUN_TEMP,omitempty"`
	F_MS_TUN_VOL       float64      `xml:"F_MS_TUN_VOL,omitempty"`
	F_MS_TYPE          MashStepType `xml:"F_MS_TYPE,omitempty"`
	_MOD_              units.Date   `xml:"_MOD_,omitempty"`
}

type MashStepType int
