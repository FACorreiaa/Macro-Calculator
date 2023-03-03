package goals

import (
	"FACorreiaa/Macro-Calculator/constants"

	"fmt"

	"github.com/manifoldco/promptui"
)

func calculateGoals(tdee float64) (float64, float64, float64) {

	var fatLoss = tdee - constants.Caloric_Deficit
	var bulk = tdee + constants.Caloric_Excedent
	fmt.Printf("Your calorie intake on cut are: %.2f\n", fatLoss)
	fmt.Printf("Your calorie intake on bulk are: %.2f\n", bulk)
	fmt.Printf("Your calorie intake on maintenance are: %.2f\n", tdee)

	return fatLoss, tdee, bulk
}

func chooseGoal() string {
	prompt := promptui.Select{
		Label: `Select your weight goal:
		- Maintenance;
		- Bulking;
		- FatLoss;
		`,
		Templates: nil,
		Items: []string{
			constants.Maintenance,
			constants.Bulking,
			constants.Fat_Loss,
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

func GetGoal(tdee float64) float64 {
	var option = chooseGoal()
	var fatloss, maintenance, bulk = calculateGoals(tdee)
	mapGoals := make(map[string]float64)
	mapGoals[constants.Maintenance] = maintenance
	mapGoals[constants.Fat_Loss] = fatloss
	mapGoals[constants.Bulking] = bulk
	return (mapGoals[option])
}
