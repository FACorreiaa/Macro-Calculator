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
	tdeeVale, UserData, activityDescription, Unit := tdee.CalculateTdee()
	goal, option := goals.GetGoal(tdeeVale)
	macros, nutPlan := plan.CalculateMacroNutrients(goal)

	fmt.Println(magenta("****************************************************************************"))
	fmt.Println(orange("Inspired by:  https://prophysiquemacros.com/"))
	fmt.Println(magenta("Using the formula: Mifflin-St Jeor Equation"))
	log.Printf("\nAge: %.0f Height:%.0f %s Weight:%.1f %s Gender: %s \n",
		UserData.Age,
		UserData.Height, red(Unit.Height),
		UserData.Weight, red(Unit.Weight),
		UserData.Gender)
	fmt.Printf("TDEE: %.2f\n", tdeeVale)
	fmt.Printf("Goal: %.2f\n", goal)
	fmt.Printf("Activity: %s\n", red(activityDescription))
	fmt.Printf("Option: %s\n", red(option))
	fmt.Printf("Nutrition plan: %s\n", red(nutPlan))
	fmt.Print(orange("Based on your TDEE and your goal, your macro distributions are: \n"))
	fmt.Printf("\n Protein %.2f: Fats:  %.2f Carbs %.2f \n", macros.Protein, macros.Fats, macros.Carbs)
	fmt.Println(magenta("****************************************************************************"))

}
