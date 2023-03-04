package tdee

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"

	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

type UserData struct {
	Age    float64
	Weight float64
	Height float64
}

const (
	maleAgeFactor   = 5
	femaleAgeFactor = -151
)

var (
	sedentaryActivityValue   = 1.2
	lightActivityValue       = 1.375
	moderateActivityValue    = 1.55
	veryActiveActivityValue  = 1.725
	extraActiveActivityValue = 1.9
)

func chooseSystem() {
	prompt := promptui.Select{
		Label: "Select Metric",
		Items: []string{"imperial", "metric"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

// MODIFY
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

func chooseMeasures(label string) float64 {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose %q\n", result)
	data, err := strconv.ParseFloat(result, 8)
	return data
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

func chooseCalculationStyle() {
	prompt := promptui.Select{
		Label: "Select the calculation style",
		Items: []string{
			constants.Option_Simple,
			constants.Option_Advanced,
		},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func calculateBMR(userData UserData, gender string) float64 {
	var ageFactor float64
	if gender == "male" {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	return (10*userData.Weight + 6.25*userData.Height - 5.0*(userData.Age)) + ageFactor
}

func calculateTDEE(bmr float64, activityValue float64) float64 {
	return bmr * activityValue
}

func CalculateTdee() float64 {
	var gender = chooseGender()
	userData := UserData{}
	age, err := menu.GetInputPrompt("Insert age")
	if err != nil {
		// handle error
	}
	userData.Age = age

	weight, err := menu.GetInputPrompt("Insert weight")
	if err != nil {
		// handle error
	}
	userData.Weight = weight

	height, err := menu.GetInputPrompt("Insert height")
	if err != nil {
		// handle error
	}
	userData.Height = height
	var activityOption = chooseActivity()
	var bmr = calculateBMR(userData, gender)
	var activityValue = getActivityValues(activityOption)
	var tdee = calculateTDEE(bmr, activityValue)
	return tdee
}
