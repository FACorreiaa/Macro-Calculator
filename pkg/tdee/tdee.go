package tdee

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func ChooseSystem() {
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
func ChooseGender() string {
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
func ChooseMeasures(label string) float64 {
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

func ChooseFormula() {
	prompt := promptui.Select{
		Label:     "Select formula",
		Templates: nil,
		Items:     []string{"athlete", "lean"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func ChooseActivity() {
	prompt := promptui.Select{
		Label:     "Select your daily activity",
		Templates: nil,
		Items:     []string{"Sedentary", "Light", "Active", "Very Active"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func ChooseIntensity() {
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

//https://tdeecalculator.net/
/*
* Women BMR = 655 + (9.6 X weight in kg) + (1.8 x height in cm) – (4.7 x age in yrs)
* Men BMR = 66 + (13.7 X weight in kg) + (5 x height in cm) – (6.8 x age in yrs)

Mifflin-St Jeor Equation FORMULA

*/
func CalculateBmr() float64 {

	var gender = ChooseGender()
	var age = ChooseMeasures("Insert age")
	var weight = ChooseMeasures("Insert weight")
	var height = ChooseMeasures("Insert height")

	var s float64
	fmt.Println("gender", gender)
	if gender == "male" {
		s = 5
	} else {
		s = -151
	}

	return (10*weight + 6.25*height - 5.0*(age)) + s

}
