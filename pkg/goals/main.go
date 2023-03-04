package goals

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"

	"fmt"
)

var (
	caloricDeficit  = 450.0
	caloricExcedent = 350.0
)

func calculateGoals(tdee float64) (float64, float64, float64) {

	var fatLoss = tdee - caloricDeficit
	var bulk = tdee + caloricExcedent
	fmt.Printf("Your calorie intake on cut are: %.2f\n", fatLoss)
	fmt.Printf("Your calorie intake on bulk are: %.2f\n", bulk)
	fmt.Printf("Your calorie intake on maintenance are: %.2f\n", tdee)

	return fatLoss, tdee, bulk
}

func GetGoal(tdee float64) float64 {
	var goal = []string{
		constants.Maintenance,
		constants.Bulking,
		constants.Fat_Loss,
	}
	var option = menu.GetSelectMenu(constants.Question_Choose_Goal, goal)
	var fatloss, maintenance, bulk = calculateGoals(tdee)
	mapGoals := make(map[string]float64)
	mapGoals[constants.Maintenance] = maintenance
	mapGoals[constants.Fat_Loss] = fatloss
	mapGoals[constants.Bulking] = bulk
	return (mapGoals[option])
}
