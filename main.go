package main

import "fmt"

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
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
}
