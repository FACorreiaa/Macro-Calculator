package main

import (
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"FACorreiaa/Macro-Calculator/pkg/plan"
	"FACorreiaa/Macro-Calculator/pkg/tdee"
	"fmt"
	"log"
)

func main() {
	tdee, UserData, activityDescription, Unit := tdee.CalculateTdee()
	goal := goals.GetGoal(tdee)
	macros := plan.CalculateMacroNutrients(goal)
	fmt.Println("************************")
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")

	fmt.Println("***** Using the formula: *****")
	fmt.Println("*** Mifflin-St Jeor Equation ***")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	fmt.Println("************************")
	println("****************************************************************************")
	log.Printf("\n*Age: %.1f Height:%.1f %s Weight:%.1f %s Gender: %s*\n", UserData.Age,
		UserData.Height, Unit.Height,
		UserData.Weight, Unit.Weight,
		UserData.Gender)
	fmt.Printf("*Your TDEE is %.2f*\n", tdee)
	fmt.Printf("*Your goal is %.2f*\n", goal)
	fmt.Printf("*Your activity Description %s \n", activityDescription)

	fmt.Printf("*Based on your TDEE and your goal, your macro distributions are: *\n")
	fmt.Printf("\nProtein %.2f: Fats:  %.2f Carbs %.2f\n", macros.Protein, macros.Fats, macros.Carbs)
	println("****************************************************************************")

}
