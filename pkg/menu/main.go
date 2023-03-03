package menu

import (
	"FACorreiaa/Macro-Calculator/constants"
	"errors"
	"strconv"

	"fmt"

	"github.com/manifoldco/promptui"
)

type Input struct {
	label string
}

type Select struct {
	label string
	items []string
}

func CustomInputMenu(i Input) float64 {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    i.label,
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

func CustomSelectMenu(s Select) {
	prompt := promptui.Select{
		Label: s.label,
		Items: s.items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
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
