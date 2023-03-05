package main

import (
	"FACorreiaa/Macro-Calculator/pkg/goals"
	"FACorreiaa/Macro-Calculator/pkg/plan"
	"FACorreiaa/Macro-Calculator/pkg/tdee"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	orange := color.New(color.FgHiYellow).SprintFunc()

	magenta := color.New(color.FgMagenta).SprintFunc()
	tdee, UserData, activityDescription, Unit := tdee.CalculateTdee()
	goal := goals.GetGoal(tdee)
	macros := plan.CalculateMacroNutrients(goal)
	fmt.Println(magenta("****************************************************************************"))
	fmt.Println(orange("****************************** Inspired by:  *******************************"))
	fmt.Println(magenta("************************* https://prophysiquemacros.com/ ******************"))
	fmt.Println(orange("************************** Using the formula: ******************************"))
	fmt.Println(magenta("************************** Mifflin-St Jeor Equation ************************"))
	fmt.Println(orange("***************************** Give it a try ********************************"))
	fmt.Println(magenta("****************************** Calculate your macros ***********************"))
	fmt.Println(orange("****************************************************************************"))

	log.Printf("\n******** Age: %.0f Height:%.0f %s Weight:%.1f %s Gender: %s******** \n",
		UserData.Age,
		UserData.Height, red(Unit.Height),
		UserData.Weight, red(Unit.Weight),
		UserData.Gender)
	fmt.Printf("***************** Your TDEE is %.2f *****************\n", tdee)
	fmt.Printf("***************** Your goal is %.2f *****************\n", goal)
	fmt.Printf("***************** Your activity Description %s *****************\n", red(activityDescription))
	fmt.Print(magenta("***************** Based on your TDEE and your goal, your macro distributions are: *****************\n"))
	fmt.Printf("\n***************** Protein %.2f: Fats:  %.2f Carbs %.2f *****************\n", macros.Protein, macros.Fats, macros.Carbs)
	fmt.Println(orange("****************************************************************************"))

}
