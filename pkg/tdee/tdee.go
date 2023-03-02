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

// used
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

// used
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

// used
func chooseActivity() string {
	prompt := promptui.Select{
		Label:     "Select your daily activity",
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

func getActivity(label string) float64 {
	mapActivity := make(map[string]int)
	mapActivity[constants.Sedentary_Activity] = 400
	mapActivity[constants.Light_Activity] = 800
	mapActivity[constants.Moderate_Activity] = 1200
	mapActivity[constants.Heavy_Activity] = 1600
	return float64(mapActivity[label])
}

func chooseIntensity() {
	prompt := promptui.Select{
		Label:     "Select your daily activity",
		Templates: nil,
		Items:     []string{"Light", "Moderate", "Difficult", "Intense"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

/*
* https://tdeecalculator.net/
* Mifflin-St Jeor Equation FORMULA
 */
func CalculateBmr() float64 {
	var gender = chooseGender()
	var age = chooseMeasures("Insert age")
	var weight = chooseMeasures("Insert weight")
	var height = chooseMeasures("Insert height")
	var activityOption = chooseActivity()
	var activityValue = getActivity(activityOption)
	var s float64
	fmt.Println("gender", gender)
	if gender == "male" {
		s = 5
	} else {
		s = -151
	}

	return (10*weight + 6.25*height - 5.0*(age)) + s + activityValue

}
