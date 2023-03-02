package goals

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"fmt"

	"github.com/manifoldco/promptui"
)

func chooseCalorieDistribution() string {
	prompt := promptui.Select{
		Label: `Select your Calorie distributions:
		- Moderate Carb (30/35/35);
		- Low Carb (40/40/20);
		- Higher Carb (30/20/50);
		`,
		Templates: nil,
		Items: []string{
			constants.Low_Carb,
			constants.Moderate_Carb,
			constants.High_Carb,
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

func calculateMacroNutrients() (float64, float64, float64) {
	//calculate based on option

	var calorieDistributionOption = chooseCalorieDistribution()
	var calorieGoal = goals.GetGoal()
	if calorieDistributionOption == constants.High_Carb {
		protein := 0.3 * calorieGoal
		fats := 0.2 * calorieGoal
		carbs := 0.5 * calorieGoal
		return protein, fats, carbs
	}

	if calorieDistributionOption == constants.Moderate_Carb {
		protein := 0.3 * calorieGoal
		fats := 0.35 * calorieGoal
		carbs := 0.35 * calorieGoal
		return protein, fats, carbs
	}

	if calorieDistributionOption == constants.Low_Carb {
		protein := 0.4 * calorieGoal
		fats := 0.4 * calorieGoal
		carbs := 0.2 * calorieGoal
		return protein, fats, carbs
	}
	return 0, 0, 0
}
