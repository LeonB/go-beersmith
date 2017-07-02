package bsmx

import "github.com/leonb/go-beersmith/units"

type Recipe struct {
	AgeData                    []*AgeData              `xml:"AgeData,omitempty"`
	F_R_ACID                   units.Bool              `xml:"F_R_ACID,omitempty"`
	F_R_ACID_CONC              float64                 `xml:"F_R_ACID_CONC,omitempty"`
	AgingProfile               *AgingProfile           `xml:"F_R_AGE,omitempty"`
	F_R_ASST_BREWER            string                  `xml:"F_R_ASST_BREWER,omitempty"`
	F_R_BASE_GRAIN             Grain                   `xml:"F_R_BASE_GRAIN,omitempty"`
	F_R_BOIL_TIMER             units.Bool              `xml:"F_R_BOIL_TIMER,omitempty"`
	F_R_BOIL_VOL_MEASURED      units.Volume            `xml:"F_R_BOIL_VOL_MEASURED,omitempty"`
	F_R_BOIL_VOL_MEASURED_SET  units.Bool              `xml:"F_R_BOIL_VOL_MEASURED_SET,omitempty"`
	F_R_BREWER                 string                  `xml:"F_R_BREWER,omitempty"`
	F_R_BTIMER_DOWN            units.Duration          `xml:"F_R_BTIMER_DOWN,omitempty"`
	CarbonationMethod          CarbonationMethod       `xml:"F_R_CARB,omitempty"`
	F_R_CARB_VOLS              units.CarbonationVolume `xml:"F_R_CARB_VOLS,omitempty"`
	F_R_COLOR_ADJ_STRING       units.Unknown           `xml:"F_R_COLOR_ADJ_STRING,omitempty"`
	F_R_DATE                   units.Date              `xml:"F_R_DATE,omitempty"`
	F_R_DESCRIPTION            string                  `xml:"F_R_DESCRIPTION,omitempty"`
	F_R_DESIRED_COLOR          units.Color             `xml:"F_R_DESIRED_COLOR,omitempty"`
	F_R_DESIRED_IBU            units.IBU               `xml:"F_R_DESIRED_IBU,omitempty"`
	F_R_DESIRED_OG             units.Gravity           `xml:"F_R_DESIRED_OG,omitempty"`
	F_R_EQUIPMENT              Equipment               `xml:"F_R_EQUIPMENT,omitempty"`
	F_R_FERMENTER              string                  `xml:"F_R_FERMENTER,omitempty"`
	F_R_FG_MEASURED            units.Gravity           `xml:"F_R_FG_MEASURED,omitempty"`
	F_R_FG_MEASURED_SET        units.Bool              `xml:"F_R_FG_MEASURED_SET,omitempty"`
	F_R_FINAL_VOL_MEASURED     units.Volume            `xml:"F_R_FINAL_VOL_MEASURED,omitempty"`
	F_R_FINAL_VOL_MEASURED_SET units.Bool              `xml:"F_R_FINAL_VOL_MEASURED_SET,omitempty"`
	F_R_FOLDER_NAME            string                  `xml:"F_R_FOLDER_NAME,omitempty"`
	F_R_GRAIN_STEEP_TEMP       units.Temperature       `xml:"F_R_GRAIN_STEEP_TEMP,omitempty"`
	F_R_GRAIN_STEEP_TIME       units.Duration          `xml:"F_R_GRAIN_STEEP_TIME,omitempty"`
	F_R_IMAGE_X                int                     `xml:"F_R_IMAGE_X,omitempty"`
	F_R_IMAGE_Y                int                     `xml:"F_R_IMAGE_Y,omitempty"`
	F_R_INCLUDE_STARTER        units.Bool              `xml:"F_R_INCLUDE_STARTER,omitempty"`
	F_R_INV_DATE               units.Date              `xml:"F_R_INV_DATE,omitempty"`
	F_R_LOCKED                 units.Bool              `xml:"F_R_LOCKED,omitempty"`
	F_R_MASH                   MashProfile             `xml:"F_R_MASH,omitempty"`
	F_R_MASH_PH                units.PH                `xml:"F_R_MASH_PH,omitempty"`
	F_R_MASH_PH_SET            units.Bool              `xml:"F_R_MASH_PH_SET,omitempty"`
	F_R_MASH_TIMER             units.Bool              `xml:"F_R_MASH_TIMER,omitempty"`
	F_R_MTIMER_DOWN            units.Duration          `xml:"F_R_MTIMER_DOWN,omitempty"`
	F_R_NAME                   string                  `xml:"F_R_NAME,omitempty"`
	F_R_NOTES                  string                  `xml:"F_R_NOTES,omitempty"`
	F_R_OG_BOIL_MEASURED       units.Gravity           `xml:"F_R_OG_BOIL_MEASURED,omitempty"`
	F_R_OG_BOIL_MEASURED_SET   units.Bool              `xml:"F_R_OG_BOIL_MEASURED_SET,omitempty"`
	F_R_OG_MASH_MEASURED       units.Gravity           `xml:"F_R_OG_MASH_MEASURED,omitempty"`
	F_R_OG_MASH_MEASURED_SET   units.Bool              `xml:"F_R_OG_MASH_MEASURED_SET,omitempty"`
	F_R_OG_MEASURED            units.Gravity           `xml:"F_R_OG_MEASURED,omitempty"`
	F_R_OG_MEASURED_SET        units.Bool              `xml:"F_R_OG_MEASURED_SET,omitempty"`
	F_R_OG_PRIMARY             units.Gravity           `xml:"F_R_OG_PRIMARY,omitempty"`
	F_R_OG_PRIMARY_SET         units.Bool              `xml:"F_R_OG_PRIMARY_SET,omitempty"`
	F_R_OG_SECONDARY           units.Gravity           `xml:"F_R_OG_SECONDARY,omitempty"`
	F_R_OG_SECONDARY_SET       units.Bool              `xml:"F_R_OG_SECONDARY_SET,omitempty"`
	F_R_OLD_BOIL_VOL           units.Volume            `xml:"F_R_OLD_BOIL_VOL,omitempty"`
	F_R_OLD_EFFICIENCY         units.Percentage        `xml:"F_R_OLD_EFFICIENCY,omitempty"`
	F_R_OLD_TYPE               RecipeOldType           `xml:"F_R_OLD_TYPE,omitempty"`
	F_R_OLD_VOL                units.Volume            `xml:"F_R_OLD_VOL,omitempty"`
	F_R_RATING                 float64                 `xml:"F_R_RATING,omitempty"`
	F_R_REBALANCE_SCALE        units.Bool              `xml:"F_R_REBALANCE_SCALE,omitempty"`
	F_R_RUNNING_GRAVITY        units.Gravity           `xml:"F_R_RUNNING_GRAVITY,omitempty"`
	F_R_RUNNING_GRAVITY_SET    units.Bool              `xml:"F_R_RUNNING_GRAVITY_SET,omitempty"`
	F_R_RUNOFF_PH              units.PH                `xml:"F_R_RUNOFF_PH,omitempty"`
	F_R_RUNOFF_PH_SET          units.Bool              `xml:"F_R_RUNOFF_PH_SET,omitempty"`
	F_R_STARTER_OG             units.Gravity           `xml:"F_R_STARTER_OG,omitempty"`
	F_R_STARTER_SIZE           units.Volume            `xml:"F_R_STARTER_SIZE,omitempty"`
	F_R_STIR_PLATE             units.Bool              `xml:"F_R_STIR_PLATE,omitempty"`
	F_R_STYLE                  Style                   `xml:"F_R_STYLE,omitempty"`
	F_R_TARGET_PH              float64                 `xml:"F_R_TARGET_PH,omitempty"`
	F_R_TYPE                   RecipeType              `xml:"F_R_TYPE,omitempty"`
	F_R_VERSION                float64                 `xml:"F_R_VERSION,omitempty"`
	F_R_VOLUME_MEASURED        units.Volume            `xml:"F_R_VOLUME_MEASURED,omitempty"`
	F_R_VOLUME_MEASURED_SET    units.Bool              `xml:"F_R_VOLUME_MEASURED_SET,omitempty"`
	Image                      *Image                  `xml:"Image,omitempty"`
	Ingredients                *Ingredients            `xml:"Ingredients,omitempty"`
	_MOD_                      units.Date              `xml:"_MOD_,omitempty"`
}

type RecipeOldType int

type RecipeType int
