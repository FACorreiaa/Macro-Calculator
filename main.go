package main

import (
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"FACorreiaa/Macro-Calculator/pkg/plan"
	"FACorreiaa/Macro-Calculator/pkg/tdee"

	"fmt"
)

func main() {
	//info and formula from:
	//https://prophysiquemacros.com/
	//Mifflin-St Jeor Equation
	fmt.Println("************************")
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")

	fmt.Println("***** Using the formula: *****")
	fmt.Println("*** Mifflin-St Jeor Equation ***")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	fmt.Println("************************")
	tdee := tdee.CalculateTdee()
	fmt.Printf("Your TDEE is %.2f\n", tdee)
	goal := goals.GetGoal(tdee)
	fmt.Printf("Your goal is %.2f\n", goal)
	protein, fats, carbs := plan.CalculateMacroNutrients(goal)
	fmt.Printf("Protein %.2f\n", protein)
	fmt.Printf("Fats %.2f\n", fats)
	fmt.Printf("Carbs %.2f\n", carbs)

}
