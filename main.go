package main

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"FACorreiaa/Macro-Calculator/pkg/menu"
	"FACorreiaa/Macro-Calculator/pkg/plan"
	"FACorreiaa/Macro-Calculator/pkg/tdee"

	"fmt"
)

func simpleMenu() {
	tdee := tdee.CalculateTdee()
	fmt.Printf("Your TDEE is %.2f\n", tdee)
	goal := goals.GetGoal(tdee)
	fmt.Printf("Your goal is %.2f\n", goal)
}

func advancedMenu() {
	tdee := tdee.CalculateTdee()
	fmt.Printf("Your TDEE is %.2f\n", tdee)
	goal := goals.GetGoal(tdee)
	fmt.Printf("Your goal is %.2f\n", goal)
	macros := plan.CalculateMacroNutrients(goal)
	fmt.Printf("Protein %.2f\n", macros.Protein)
	fmt.Printf("Fats %.2f\n", macros.Fats)
	fmt.Printf("Carbs %.2f\n", macros.Carbs)
}
func main() {
	fmt.Println("************************")
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")

	fmt.Println("***** Using the formula: *****")
	fmt.Println("*** Mifflin-St Jeor Equation ***")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	fmt.Println("************************")
	option := menu.PickOption()
	if option == constants.Option_Simple {
		simpleMenu()
	} else {
		advancedMenu()
	}

}
