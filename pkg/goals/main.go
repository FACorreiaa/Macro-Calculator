package goals

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/tdee"

	"fmt"

	"github.com/manifoldco/promptui"
)

//maintenance
//cut
//bulk

func calculateGoals() (float64, float64, float64) {
	var tdeeValue = tdee.CalculateTdee()
	var fatLoss = tdeeValue - 500
	var bulk = tdeeValue + 500
	println("Your calories on cut are: ", fatLoss)
	println("Your calories on bulk are: ", bulk)
	println("Your calories on maintenance are: ", tdeeValue)

	return fatLoss, tdeeValue, bulk
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

func GetGoal() float64 {
	var option = chooseGoal()
	var fatloss, maintenance, bulk = calculateGoals()
	mapGoals := make(map[string]float64)
	mapGoals[constants.Maintenance] = maintenance
	mapGoals[constants.Fat_Loss] = fatloss
	mapGoals[constants.Bulking] = bulk
	return (mapGoals[option])
}
