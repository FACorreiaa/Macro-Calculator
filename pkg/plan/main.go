package goals

import (
	"FACorreiaa/Macro-Calculator/pkg/tdee"
)

func CalculateMacros() (float64, float64, float64) {
	var tdeeValue = tdee.CalculateTdee()
	println(tdeeValue)
	return tdeeValue, tdeeValue, tdeeValue
}
