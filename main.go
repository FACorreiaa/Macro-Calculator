package main

import (
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"FACorreiaa/Macro-Calculator/pkg/plan"
	"FACorreiaa/Macro-Calculator/pkg/tdee"
	"fmt"
	"log"
)

func main() {

	// var activityOptions = []string{
	// 	constants.Sedentary_Activity,
	// 	constants.Light_Activity,
	// 	constants.Moderate_Activity,
	// 	constants.Heavy_Activity,
	// 	constants.Extra_Heavy_Activity,
	// }
	tdee, UserData, activityDescription, Unit := tdee.CalculateTdee()
	goal := goals.GetGoal(tdee)
	macros := plan.CalculateMacroNutrients(goal)
	// var activityOption = menu.GetSelectMenu(constants.Question_Select_Activity, activityOptions)

	// activityDescription := tdee.GetActivityExpplanation(activityOption)
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
