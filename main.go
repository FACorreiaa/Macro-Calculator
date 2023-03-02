package main

import (
	"FACorreiaa/Macro-Calculator/pkg/goals"
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
	fmt.Printf("Your TDEE is %f\n", tdee)
	goal := goals.GetGoal()
	fmt.Printf("Your goal is %f\n", goal)
}
