package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu"
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
	fmt.Println("***** Inspired by: *****")
	fmt.Println("* https://prophysiquemacros.com/ *")
	fmt.Println("***** Give it a try *****")
	fmt.Println("***** Calculate your macros *****")
	menu := wmenu.NewMenu("What is your favorite food?")
	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " is your favorite food."); return nil })
	menu.Option("Pizza", nil, true, nil)
	menu.Option("Ice Cream", nil, false, nil)
	menu.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
