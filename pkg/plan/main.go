package plan

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"
)

type MacroNutrients struct {
	Protein float64
	Fats    float64
	Carbs   float64
}

const (
	HighCarbRatios     = "High Carb"
	ModerateCarbRatios = "Moderate Carb"
	LowCarbRatios      = "Low Carb"
)

var macroRatios = map[string]struct {
	ProteinRatio float64
	FatRatio     float64
	CarbRatio    float64
}{
	HighCarbRatios: {
		ProteinRatio: 0.3,
		FatRatio:     0.2,
		CarbRatio:    0.5,
	},
	ModerateCarbRatios: {
		ProteinRatio: 0.3,
		FatRatio:     0.35,
		CarbRatio:    0.35,
	},
	LowCarbRatios: {
		ProteinRatio: 0.4,
		FatRatio:     0.4,
		CarbRatio:    0.2,
	},
}

var (
	fatGramValue     = 9
	proteinGramValue = 4
	carbGramValue    = 4
)

func calculateMacroDistribution(calorieFactor float64, calorieGoal float64, caloriesPerGram int) float64 {
	return (calorieFactor * calorieGoal) / float64(caloriesPerGram)
}

func CalculateMacroNutrients(calorieGoal float64) (MacroNutrients, string) {
	calorieDistributionOptions := []string{
		constants.LowCarb,
		constants.ModerateCarb,
		constants.HighCarb,
	}
	calorieDistributionOption := menu.GetSelectMenu(constants.QuestionSelectMacroDistribution, calorieDistributionOptions)
	if ratios, ok := macroRatios[calorieDistributionOption]; ok {
		protein := calculateMacroDistribution(ratios.ProteinRatio, calorieGoal, proteinGramValue)
		fats := calculateMacroDistribution(ratios.FatRatio, calorieGoal, fatGramValue)
		carbs := calculateMacroDistribution(ratios.CarbRatio, calorieGoal, carbGramValue)

		return MacroNutrients{
			Protein: protein,
			Fats:    fats,
			Carbs:   carbs,
		}, calorieDistributionOption
	}

	return MacroNutrients{}, calorieDistributionOption
}
