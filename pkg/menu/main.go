package menu

import (
	"FACorreiaa/Macro-Calculator/constants"
	"errors"
	"strconv"

	"fmt"

	"github.com/manifoldco/promptui"
)

func GetInputPrompt(label string) (float64, error) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	data, err := strconv.ParseFloat(result, 64)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func GetSelectMenu(label string, options []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}

func PickOption() string {
	prompt := promptui.Select{
		Label: `Select your weight goal:
		- Simple (only see basic TDEE with objective)
		- Advanced (get values of protein, carbs and fats detailed);
		`,
		Templates: nil,
		Items: []string{
			constants.Option_Simple,
			constants.Option_Advanced,
		},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)

	return result
}
