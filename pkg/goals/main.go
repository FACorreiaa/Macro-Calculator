package goals

import (
	"FACorreiaa/Macro-Calculator/pkg/tdee"
)

//maintenance
//cut
//bulk

func CalculateGoals() (float64, float64, float64) {
	var tdeeValue = tdee.CalculateTdee()
	var fatLoss = tdeeValue - 400
	var bulk = tdeeValue + 300
	println("Your calories on cut are: ", fatLoss)
	println("Your calories on bulk are: ", bulk)
	println("Your calories on maintenance are: ", tdeeValue)
	return fatLoss, tdeeValue, bulk
}
