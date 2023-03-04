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
}

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
	mapActivity[constants.Sedentary_Activity] = sedentaryActivityValue
	mapActivity[constants.Light_Activity] = lightActivityValue
	mapActivity[constants.Moderate_Activity] = moderateActivityValue
	mapActivity[constants.Heavy_Activity] = veryActiveActivityValue
	mapActivity[constants.Extra_Heavy_Activity] = extraActiveActivityValue
	return float64(mapActivity[label])
}

func calculateBMR(userData UserData, gender string) float64 {
	var ageFactor float64
	if gender == "male" {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	if userData.Metric == "metric" {
		return (10*userData.Weight + 6.25*userData.Height - 5.0*(userData.Age)) + ageFactor
	} else {
		return (4.536*userData.Weight + 15.88*userData.Height - 5.0*(userData.Age)) + ageFactor
	}

}

func calculateTDEE(bmr float64, activityValue float64) float64 {
	return bmr * activityValue
}

func CalculateTdee() float64 {
	userData := UserData{}
	var genderOptions = []string{constants.Male, constants.Female}
	var metricOptions = []string{
		constants.Metric, constants.Imperial,
	}
	var activityOptions = []string{
		constants.Sedentary_Activity,
		constants.Light_Activity,
		constants.Moderate_Activity,
		constants.Heavy_Activity,
		constants.Extra_Heavy_Activity,
	}
	var metric = menu.GetSelectMenu(constants.Question_Select_Measure, metricOptions)
	userData.Metric = metric
	var gender = menu.GetSelectMenu(constants.Question_Select_Gender, genderOptions)
	age, err := menu.GetInputPrompt(constants.Question_Insert_Age)
	if err != nil {
		log.Fatal(err)
	}
	userData.Age = age

	weight, err := menu.GetInputPrompt(constants.Question_Insert_Weight)
	if err != nil {
		log.Fatal(err)
	}
	userData.Weight = convertWeight(weight, userData)

	height, err := menu.GetInputPrompt(constants.Question_Insert_Height)
	if err != nil {
		log.Fatal(err)
	}
	userData.Height = convertHeight(height, userData)
	var activityOption = menu.GetSelectMenu(constants.Question_Select_Activity, activityOptions)
	var bmr = calculateBMR(userData, gender)
	var activityValue = getActivityValues(activityOption)
	var tdee = calculateTDEE(bmr, activityValue)
	return tdee
}
