package tdee

import (
	"FACorreiaa/Macro-Calculator/constants"
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
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
	age, err := strconv.ParseFloat(result, 8)
	return age
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
	mapActivity := make(map[string]int)
	mapActivity[constants.Sedentary_Activity] = 400
	mapActivity[constants.Light_Activity] = 800
	mapActivity[constants.Moderate_Activity] = 1200
	mapActivity[constants.Heavy_Activity] = 1600
	return float64(mapActivity[label])
}

func chooseCalculationStyle() {
	prompt := promptui.Select{
		Label:     "Select the calculation style",
		Templates: nil,
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

func CalculateTdee() float64 {
	var gender = chooseGender()
	var age = chooseMeasures("Insert age")
	var weight = chooseMeasures("Insert weight")
	var height = chooseMeasures("Insert height")

	var activityOption = chooseActivity()
	var activityValue = getActivityValues(activityOption)

	var ageFactor float64
	fmt.Println("gender", gender)
	if gender == "male" {
		ageFactor = 5
	} else {
		ageFactor = -151
	}

	result := (10*weight + 6.25*height - 5.0*(age)) + ageFactor + (activityValue)
	return result
}
