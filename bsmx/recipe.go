package bsmx

import "github.com/leonb/go-beersmith/units"

type Recipe struct {
	FermentationData FermentationData `xml:"AgeData,omitempty"`
	Acid             units.Bool       `xml:"F_R_ACID,omitempty"`
	// Percent acid concentration for mash PH adjustment
	AcidConc            units.Percentage        `xml:"F_R_ACID_CONC,omitempty"`
	FermentationProfile *FermentationProfile    `xml:"F_R_AGE,omitempty"`
	AsstBrewer          string                  `xml:"F_R_ASST_BREWER,omitempty"`
	BaseGrain           Grain                   `xml:"F_R_BASE_GRAIN,omitempty"`
	BoilTimer           units.Bool              `xml:"F_R_BOIL_TIMER,omitempty"`
	BoilVolMeasured     units.Volume            `xml:"F_R_BOIL_VOL_MEASURED,omitempty"`
	BoilVolMeasuredSet  units.Bool              `xml:"F_R_BOIL_VOL_MEASURED_SET,omitempty"`
	Brewer              string                  `xml:"F_R_BREWER,omitempty"`
	BtimerDown          units.Duration          `xml:"F_R_BTIMER_DOWN,omitempty"`
	CarbonationMethod   CarbonationMethod       `xml:"F_R_CARB,omitempty"`
	CarbVols            units.CarbonationVolume `xml:"F_R_CARB_VOLS,omitempty"`
	ColorAdjString      units.Unknown           `xml:"F_R_COLOR_ADJ_STRING,omitempty"`
	Date                units.Date              `xml:"F_R_DATE,omitempty"`
	Description         string                  `xml:"F_R_DESCRIPTION,omitempty"`
	DesiredColor        units.Color             `xml:"F_R_DESIRED_COLOR,omitempty"`
	DesiredIBU          units.IBU               `xml:"F_R_DESIRED_IBU,omitempty"`
	DesiredOG           units.Gravity           `xml:"F_R_DESIRED_OG,omitempty"`
	EquipmentProfile    EquipmentProfile        `xml:"F_R_EQUIPMENT,omitempty"`
	Fermenter           string                  `xml:"F_R_FERMENTER,omitempty"`
	FGMeasured          units.Gravity           `xml:"F_R_FG_MEASURED,omitempty"`
	FGMeasuredSet       units.Bool              `xml:"F_R_FG_MEASURED_SET,omitempty"`
	FinalVolMeasured    units.Volume            `xml:"F_R_FINAL_VOL_MEASURED,omitempty"`
	FinalVolMeasuredSet units.Bool              `xml:"F_R_FINAL_VOL_MEASURED_SET,omitempty"`
	FolderName          string                  `xml:"F_R_FOLDER_NAME,omitempty"`
	GrainSteepTemp      units.Temperature       `xml:"F_R_GRAIN_STEEP_TEMP,omitempty"`
	GrainSteepTime      units.Duration          `xml:"F_R_GRAIN_STEEP_TIME,omitempty"`
	ImageX              int                     `xml:"F_R_IMAGE_X,omitempty"`
	ImageY              int                     `xml:"F_R_IMAGE_Y,omitempty"`
	IncludeStarter      units.Bool              `xml:"F_R_INCLUDE_STARTER,omitempty"`
	InvDate             units.Date              `xml:"F_R_INV_DATE,omitempty"`
	Locked              units.Bool              `xml:"F_R_LOCKED,omitempty"`
	MashProfile         MashProfile             `xml:"F_R_MASH,omitempty"`
	MashPH              units.PH                `xml:"F_R_MASH_PH,omitempty"`
	MashPHSet           units.Bool              `xml:"F_R_MASH_PH_SET,omitempty"`
	MashTimer           units.Bool              `xml:"F_R_MASH_TIMER,omitempty"`
	MashTimerDown       units.Duration          `xml:"F_R_MTIMER_DOWN,omitempty"`
	Name                string                  `xml:"F_R_NAME,omitempty"`
	Notes               string                  `xml:"F_R_NOTES,omitempty"`
	OGBoilMeasured      units.Gravity           `xml:"F_R_OG_BOIL_MEASURED,omitempty"`
	OGBoilMeasuredSet   units.Bool              `xml:"F_R_OG_BOIL_MEASURED_SET,omitempty"`
	OGMashMeasured      units.Gravity           `xml:"F_R_OG_MASH_MEASURED,omitempty"`
	OGMashMeasuredSet   units.Bool              `xml:"F_R_OG_MASH_MEASURED_SET,omitempty"`
	OGMeasured          units.Gravity           `xml:"F_R_OG_MEASURED,omitempty"`
	OGMeasuredSet       units.Bool              `xml:"F_R_OG_MEASURED_SET,omitempty"`
	OGPrimary           units.Gravity           `xml:"F_R_OG_PRIMARY,omitempty"`
	OGPrimarySet        units.Bool              `xml:"F_R_OG_PRIMARY_SET,omitempty"`
	OGSecondary         units.Gravity           `xml:"F_R_OG_SECONDARY,omitempty"`
	OGSecondarySet      units.Bool              `xml:"F_R_OG_SECONDARY_SET,omitempty"`
	OldBoilVol          units.Volume            `xml:"F_R_OLD_BOIL_VOL,omitempty"`
	OldEfficiency       units.Percentage        `xml:"F_R_OLD_EFFICIENCY,omitempty"`
	OldType             RecipeOldType           `xml:"F_R_OLD_TYPE,omitempty"`
	OldVol              units.Volume            `xml:"F_R_OLD_VOL,omitempty"`
	Rating              float64                 `xml:"F_R_RATING,omitempty"`
	RebalanceScale      units.Bool              `xml:"F_R_REBALANCE_SCALE,omitempty"`
	RunningGravity      units.Gravity           `xml:"F_R_RUNNING_GRAVITY,omitempty"`
	RunningGravitySet   units.Bool              `xml:"F_R_RUNNING_GRAVITY_SET,omitempty"`
	RunoffPH            units.PH                `xml:"F_R_RUNOFF_PH,omitempty"`
	RunoffPHSet         units.Bool              `xml:"F_R_RUNOFF_PH_SET,omitempty"`
	StarterOG           units.Gravity           `xml:"F_R_STARTER_OG,omitempty"`
	StarterSize         units.Volume            `xml:"F_R_STARTER_SIZE,omitempty"`
	StirPlate           units.Bool              `xml:"F_R_STIR_PLATE,omitempty"`
	Style               Style                   `xml:"F_R_STYLE,omitempty"`
	TargetPH            units.PH                `xml:"F_R_TARGET_PH,omitempty"`
	Type                RecipeType              `xml:"F_R_TYPE,omitempty"`
	Version             float64                 `xml:"F_R_VERSION,omitempty"`
	VolumeMeasured      units.Volume            `xml:"F_R_VOLUME_MEASURED,omitempty"`
	VolumeMeasuredSet   units.Bool              `xml:"F_R_VOLUME_MEASURED_SET,omitempty"`
	Image               *Image                  `xml:"Image,omitempty"`
	Ingredients         *Ingredients            `xml:"Ingredients,omitempty"`
	_MOD_               units.Date              `xml:"_MOD_,omitempty"`
}

type RecipeOldType int

type RecipeType int
