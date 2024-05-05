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

func GetGoal(tdee float64) (float64, string) {
	var goal = []string{
		constants.Maintenance,
		constants.Bulking,
		constants.FatLoss,
	}
	var option = menu.GetSelectMenu(constants.QuestionChooseGoal, goal)
	var fatLoss, maintenance, bulk = calculateGoals(tdee)
	mapGoals := make(map[string]float64)
	mapGoals[constants.Maintenance] = maintenance
	mapGoals[constants.FatLoss] = fatLoss
	mapGoals[constants.Bulking] = bulk
	return mapGoals[option], option
}
