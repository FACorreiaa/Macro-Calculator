package menu

import (
	"FACorreiaa/Macro-Calculator/constants"

	"fmt"

	"github.com/manifoldco/promptui"
)

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
