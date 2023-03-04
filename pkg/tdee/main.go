package tdee

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"
	"log"

	"fmt"

	"github.com/manifoldco/promptui"
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
	if data.Metric == "metric" {
		return value
	} else {
		return value * 0.453592 // 1 lb = 0.453592 kg
	}
}

func convertHeight(value float64, data UserData) float64 {
	if data.Metric == "metric" {
		return value
	} else {
		return value * 2.54 // 1 in = 2.54 cm
	}
}

func chooseGender() string {
	prompt := promptui.Select{
		Label: "Select gender",
		Items: []string{"male", "female"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	fmt.Printf("You choose %q\n", result)

	return result
}

func chooseActivity() string {
	prompt := promptui.Select{
		Label: `Select your daily activity:
		Sedentary: Spend most of the day sitting (e.g. bank taller, desk job)
		Lightly Activity: Spend a good part of the day on your feet (e.g. teacher, salesman)
		Active: Spend a good part of the day doing some physical activity (e.g. waitress, mailman)
		Very Active: Spend most of the day doing heavy physical activity (e.g. bike messenger, carpenter)
		`,
		Templates: nil,
		Items: []string{
			constants.Sedentary_Activity,
			constants.Light_Activity,
			constants.Moderate_Activity,
			constants.Heavy_Activity},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)

	return result
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

	var metricOptions = []string{constants.Metric, constants.Imperial}
	var metric = menu.GetSelectMenu("Select Measure System", metricOptions)
	userData.Metric = metric
	var gender = chooseGender()
	age, err := menu.GetInputPrompt("Insert age")
	if err != nil {
		log.Fatal(err)
	}
	userData.Age = age

	weight, err := menu.GetInputPrompt("Insert weight")
	if err != nil {
		log.Fatal(err)
	}
	userData.Weight = convertWeight(weight, userData)

	height, err := menu.GetInputPrompt("Insert height")
	if err != nil {
		log.Fatal(err)
	}
	userData.Height = convertHeight(height, userData)
	var activityOption = chooseActivity()
	var bmr = calculateBMR(userData, gender)
	var activityValue = getActivityValues(activityOption)
	var tdee = calculateTDEE(bmr, activityValue)
	return tdee
}
