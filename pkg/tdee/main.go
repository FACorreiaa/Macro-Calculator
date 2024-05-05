package tdee

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"
	"log"
)

type UserData struct {
	Age    float64
	Weight float64
	Height float64
	Metric string
	Gender string
	Goal   string
}

type Units struct {
	Height string
	Weight string
}

const (
	MetricHeight   = "cm"
	MetricWeight   = "kg"
	ImperialHeight = "ft"
	ImperialWeight = "lb"
)

const (
	maleAgeFactor   = 5
	femaleAgeFactor = -161
)

var (
	sedentaryActivityValue   = 1.2
	lightActivityValue       = 1.375
	moderateActivityValue    = 1.55
	veryActiveActivityValue  = 1.725
	extraActiveActivityValue = 1.9
)

func convertWeight(value float64, data UserData) float64 {
	if data.Metric == constants.Metric {
		return value
	} else {
		return value * 0.453592 // 1 lb = 0.453592 kg
	}
}

func convertHeight(value float64, data UserData) float64 {
	if data.Metric == constants.Metric {
		return value
	} else {
		return value * 2.54 // 1 in = 2.54 cm
	}
}

func getActivityValues(label string) float64 {
	mapActivity := make(map[string]float64)
	mapActivity[constants.SedentaryActivity] = sedentaryActivityValue
	mapActivity[constants.LightActivity] = lightActivityValue
	mapActivity[constants.ModerateActivity] = moderateActivityValue
	mapActivity[constants.HeavyActivity] = veryActiveActivityValue
	mapActivity[constants.ExtraHeavyActivity] = extraActiveActivityValue
	return float64(mapActivity[label])
}

func getActivityExpplanation(activity float64) string {
	mapActivityLabel := make(map[float64]string)
	mapActivityLabel[sedentaryActivityValue] = constants.SedentaryActivityDescription
	mapActivityLabel[lightActivityValue] = constants.LightActivityDescription
	mapActivityLabel[moderateActivityValue] = constants.ModerateActivityDescription
	mapActivityLabel[veryActiveActivityValue] = constants.HeavyActivityDescription
	mapActivityLabel[extraActiveActivityValue] = constants.ExtraHeavyActivityDescription
	return (mapActivityLabel[activity])

}

func calculateBMR(data UserData, gender string) float64 {
	var ageFactor float64
	if gender == "male" {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	if data.Metric == "metric" {
		return (10*data.Weight + 6.25*data.Height - 5.0*(data.Age)) + ageFactor
	} else {
		return (4.536*data.Weight + 15.88*data.Height - 5.0*(data.Age)) + ageFactor
	}
}

func calculateTDEE(bmr float64, activityValue float64) float64 {
	return bmr * activityValue
}

func GetCorrectUnitSystem(metric string) Units {
	unit := Units{}
	switch metric {
	case "metric":
		unit.Height = MetricHeight
		unit.Weight = MetricWeight
	case "imperial":
		unit.Height = ImperialHeight
		unit.Weight = ImperialWeight
	default:
		return unit
	}
	return unit
}

func CalculateTdee() (float64, UserData, string, Units) {
	UserData := UserData{}
	Unit := Units{}
	genderOptions := []string{constants.Male, constants.Female}
	metricOptions := []string{
		constants.Metric, constants.Imperial,
	}
	metric := menu.GetSelectMenu(constants.QuestionSelectMeasure, metricOptions)
	gender := menu.GetSelectMenu(constants.QuestionSelectGender, genderOptions)

	activityOptions := []string{
		constants.SedentaryActivity,
		constants.LightActivity,
		constants.ModerateActivity,
		constants.HeavyActivity,
		constants.ExtraHeavyActivity,
	}
	UserData.Metric = metric
	UserData.Gender = gender
	age, err := menu.GetInputPrompt(constants.QuestionInsertAge)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Age = age

	weight, err := menu.GetInputPrompt(constants.QuestionInsertWeight)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Weight = convertWeight(weight, UserData)

	height, err := menu.GetInputPrompt(constants.QuestionInsertHeight)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Height = convertHeight(height, UserData)
	activityOption := menu.GetSelectMenu(constants.QuestionSelectActivity, activityOptions)
	bmr := calculateBMR(UserData, UserData.Gender)
	activityValue := getActivityValues(activityOption)
	activityDescription := getActivityExpplanation(activityValue)

	tdee := calculateTDEE(bmr, activityValue)
	unit := GetCorrectUnitSystem(UserData.Metric)
	Unit.Height = unit.Height
	Unit.Weight = unit.Weight
	return tdee, UserData, activityDescription, Unit
}
