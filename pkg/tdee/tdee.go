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

func ChooseGender() {
	prompt := promptui.Select{
		Label: "Select gender",
		Items: []string{"male", "female"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func ChooseMeasures(measure string) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    measure,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func ChooseAge() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "age",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
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
 */
func CalculateMenBmr() float64 {
	var age = 34.0
	var weight, height = 85.0, 185.0

	return 66 + (13.7 * weight) + (5 * height) - (6.8 * age)

}

func CalculateWomanBmr() float64 {
	var age = 41.0
	var weight, height = 71.0, 171.0

	return 655 + (9.6 * weight) + (1.8 * height) - (4.7 * age)

}
