package main

import (
	"fmt"

	"FACorreiaa/Macro-Calculator/pkg/tdee"
)

type Goals struct {
	name      string
	intensity int
}

type Plan struct {
	name string
}

type Meals struct {
	numberOfMeals int
}

func main() {
	//info and formula from:
	//https://prophysiquemacros.com/

	//creatre struct for values
	//kg to lb and lb to kg values
	//get formulas with values
	//make separate functions for formulas
	//make calculations
	//show values

	//Mifflin-St Jeor Equation
	fmt.Println("************************")
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")

	fmt.Println("***** Using the formula: *****")
	fmt.Println("*** Mifflin-St Jeor Equation ***")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	fmt.Println("************************")
	//tdee.ChooseSystem()
	//tdee.ChooseGender()
	result := tdee.CalculateTdee()
	fmt.Println("You calorie intake for yout objective is: ", result)
	// tdee.ChooseMeasures("height")
	// tdee.ChooseMeasures("weight")
	// tdee.ChooseMeasures("age")
	// tdee.ChooseFormula()
	// tdee.ChooseActivity()
	// tdee.ChooseActivity()
	// tdee.ChooseMeasures("How many week days per week do you exercise?")
	// tdee.ChooseMeasures("How many minutes per day do you exercise (cardio & weight lifting combined)?")
}
