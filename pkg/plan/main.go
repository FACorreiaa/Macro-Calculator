package plan

import (
	"FACorreiaa/Macro-Calculator/constants"
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

func calculateMacroDistribution(calorieFactor float64, calorieGoal float64, caloriesPerGram int) float64 {
	return (calorieFactor * calorieGoal) / float64(caloriesPerGram)
}

func CalculateMacroNutrients(calorieGoal float64) (float64, float64, float64) {
	//calculate based on option

	var calorieDistributionOption = chooseCalorieDistribution()
	if calorieDistributionOption == constants.High_Carb {
		protein := calculateMacroDistribution(0.3, calorieGoal, 4)
		fats := calculateMacroDistribution(0.2, calorieGoal, 9)
		carbs := calculateMacroDistribution(0.5, calorieGoal, 4)
		return protein, fats, carbs
	}

	if calorieDistributionOption == constants.Moderate_Carb {
		protein := calculateMacroDistribution(0.3, calorieGoal, 4)
		fats := calculateMacroDistribution(0.35, calorieGoal, 9)
		carbs := calculateMacroDistribution(0.35, calorieGoal, 4)
		return protein, fats, carbs
	}

	if calorieDistributionOption == constants.Low_Carb {
		protein := calculateMacroDistribution(0.4, calorieGoal, 4)
		fats := calculateMacroDistribution(0.4, calorieGoal, 9)
		carbs := calculateMacroDistribution(0.2, calorieGoal, 4)
		return protein, fats, carbs
	}

	return 0, 0, 0
}
