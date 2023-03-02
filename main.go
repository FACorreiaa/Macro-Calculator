package main

import (
	"fmt"

	"FACorreiaa/Macro-Calculator/pkg/tdee"
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
	result := tdee.CalculateTdee()
	fmt.Println("You calorie intake for yout objective is: ", result)
}
