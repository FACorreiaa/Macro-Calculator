package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
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

func chooseSystem() {
	prompt := promptui.Select{
		Label: "Select Metric",
		Items: []string{"imperial", "metric"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func chooseGender() {
	prompt := promptui.Select{
		Label: "Select gender",
		Items: []string{"male", "female"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
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
	fmt.Println("************************")
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	fmt.Println("************************")
	chooseSystem()
	chooseGender()
}
