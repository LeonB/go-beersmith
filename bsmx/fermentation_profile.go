package bsmx

import "github.com/leonb/go-beersmith/units"

type FermentationProfiles []FermentationProfile

type FermentationProfile struct {
	Name                    string                  `xml:"F_A_NAME"`
	PrimaryTemperature      units.Temperature       `xml:"F_A_PRIM_TEMP"`
	PrimaryEndTemperature   units.Temperature       `xml:"F_A_PRIM_END_TEMP"`
	SecundaryTemperature    units.Temperature       `xml:"F_A_SEC_TEMP"`
	SecundaryEndTemperature units.Temperature       `xml:"F_A_SEC_END_TEMP"`
	TertiaryTemperature     units.Temperature       `xml:"F_A_TERT_TEMP"`
	TertiaryEndTemperature  units.Temperature       `xml:"F_A_TERT_END_TEMP"`
	AgeTemperature          units.Temperature       `xml:"F_A_AGE_TEMP"`
	EndAgeTemperature       units.Temperature       `xml:"F_A_END_AGE_TEMP"`
	BulkTemperature         units.Temperature       `xml:"F_A_BULK_TEMP"`
	BulkEndTemperature      units.Temperature       `xml:"F_A_BULK_END_TEMP"`
	EndTemperaturesSet      units.Bool              `xml:"F_A_END_TEMPS_SET"`
	PrimaryDays             units.Duration          `xml:"F_A_PRIM_DAYS"`
	SecundaryDays           units.Duration          `xml:"F_A_SEC_DAYS"`
	TertiaryDays            units.Duration          `xml:"F_A_TERT_DAYS"`
	BulkDays                units.Duration          `xml:"F_A_BULK_DAYS"`
	Age                     units.Duration          `xml:"F_A_AGE"`
	Type                    FermentationProfileType `xml:"F_A_TYPE"`
	Notes                   string                  `xml:"F_A_NOTES"`
}

type FermentationProfileType int

const (
	FermentationProfileTypeSingleStage FermentationProfileType = 0
	FermentationProfileTypeTwoStage    FermentationProfileType = 1
	FermentationProfileTypeThreeStage  FermentationProfileType = 2
	FermentationProfileTypeFourStage   FermentationProfileType = 3
)

func (fpt FermentationProfileType) String() string {
	switch fpt {
	case 0:
		return "Single Stage"
	case 1:
		return "Two Stage"
	case 2:
		return "Three Stage"
	case 3:
		return "Four Stage"
	}

	return ""
}
