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

Mifflin-St Jeor Equation FORMULA

*/
func CalculateBmr() float64 {
	var age = 34.0
	var weight, height = 85.0, 185.0
	var gender = ChooseGender()
	var s float64
	fmt.Println("gender", gender)
	if gender == "male" {
		s = 5
	} else {
		s = -151
	}

	return (10*weight + 6.25*height - 5.0*age) + s

}
