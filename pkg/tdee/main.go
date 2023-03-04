package tdee

import (
	"FACorreiaa/Macro-Calculator/constants"
	"FACorreiaa/Macro-Calculator/pkg/menu"
	"log"
)

type UserData struct {
	Age    float64
	Weight float64
	Height float64
	Metric string
	Gender string
	Goal   string
}

type Units struct {
	Height string
	Weight string
}

const (
	MetricHeight   = "cm"
	MetricWeight   = "kg"
	ImperialHeight = "ft"
	ImperialWeight = "lb"
)

const (
	maleAgeFactor   = 5
	femaleAgeFactor = -161
)

var (
	sedentaryActivityValue   = 1.2
	lightActivityValue       = 1.375
	moderateActivityValue    = 1.55
	veryActiveActivityValue  = 1.725
	extraActiveActivityValue = 1.9
)

func convertWeight(value float64, data UserData) float64 {
	if data.Metric == constants.Metric {
		return value
	} else {
		return value * 0.453592 // 1 lb = 0.453592 kg
	}
}

func convertHeight(value float64, data UserData) float64 {
	if data.Metric == constants.Metric {
		return value
	} else {
		return value * 2.54 // 1 in = 2.54 cm
	}
}

func getActivityValues(label string) float64 {
	mapActivity := make(map[string]float64)
	mapActivity[constants.Sedentary_Activity] = sedentaryActivityValue
	mapActivity[constants.Light_Activity] = lightActivityValue
	mapActivity[constants.Moderate_Activity] = moderateActivityValue
	mapActivity[constants.Heavy_Activity] = veryActiveActivityValue
	mapActivity[constants.Extra_Heavy_Activity] = extraActiveActivityValue
	return float64(mapActivity[label])
}

func getActivityExpplanation(activity float64) string {
	mapActivityLabel := make(map[float64]string)
	mapActivityLabel[sedentaryActivityValue] = constants.Sedentary_Activity_Description
	mapActivityLabel[lightActivityValue] = constants.Light_Activity_Description
	mapActivityLabel[moderateActivityValue] = constants.Moderate_Activity_Description
	mapActivityLabel[veryActiveActivityValue] = constants.Heavy_Activity_Description
	mapActivityLabel[extraActiveActivityValue] = constants.Extra_Heavy_Activity_Description
	return (mapActivityLabel[activity])

}

func calculateBMR(UserData UserData, gender string) float64 {
	var ageFactor float64
	if gender == "male" {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	if UserData.Metric == "metric" {
		return (10*UserData.Weight + 6.25*UserData.Height - 5.0*(UserData.Age)) + ageFactor
	} else {
		return (4.536*UserData.Weight + 15.88*UserData.Height - 5.0*(UserData.Age)) + ageFactor
	}

}

func calculateTDEE(bmr float64, activityValue float64) float64 {
	return bmr * activityValue
}

func GetCorrectUnitSystem(metric string) Units {
	unit := Units{}
	switch metric {
	case "metric":
		unit.Height = MetricHeight
		unit.Weight = MetricWeight
	case "imperial":
		unit.Height = ImperialHeight
		unit.Weight = ImperialWeight
	default:
		return unit
	}
	return unit
}

func CalculateTdee() (float64, UserData, string, Units) {
	UserData := UserData{}
	Unit := Units{}
	var genderOptions = []string{constants.Male, constants.Female}
	var metricOptions = []string{
		constants.Metric, constants.Imperial,
	}
	var metric = menu.GetSelectMenu(constants.Question_Select_Measure, metricOptions)
	var gender = menu.GetSelectMenu(constants.Question_Select_Gender, genderOptions)

	var activityOptions = []string{
		constants.Sedentary_Activity,
		constants.Light_Activity,
		constants.Moderate_Activity,
		constants.Heavy_Activity,
		constants.Extra_Heavy_Activity,
	}
	UserData.Metric = metric
	UserData.Gender = gender
	age, err := menu.GetInputPrompt(constants.Question_Insert_Age)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Age = age

	weight, err := menu.GetInputPrompt(constants.Question_Insert_Weight)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Weight = convertWeight(weight, UserData)

	height, err := menu.GetInputPrompt(constants.Question_Insert_Height)
	if err != nil {
		log.Fatal(err)
	}
	UserData.Height = convertHeight(height, UserData)
	var activityOption = menu.GetSelectMenu(constants.Question_Select_Activity, activityOptions)
	var bmr = calculateBMR(UserData, UserData.Gender)
	var activityValue = getActivityValues(activityOption)
	var activityDescription = getActivityExpplanation(activityValue)

	var tdee = calculateTDEE(bmr, activityValue)
	var unit = GetCorrectUnitSystem(UserData.Metric)
	Unit.Height = unit.Height
	Unit.Weight = unit.Weight
	//goal := goals.GetGoal(tdee)
	//macros := plan.CalculateMacroNutrients(goal)
	return tdee, UserData, activityDescription, Unit
}

// func MainMenu() {
// 	var activityOptions = []string{
// 		constants.Sedentary_Activity,
// 		constants.Light_Activity,
// 		constants.Moderate_Activity,
// 		constants.Heavy_Activity,
// 		constants.Extra_Heavy_Activity,
// 	}

// 	var metricOptions = []string{
// 		constants.Metric, constants.Imperial,
// 	}
// 	var metric = menu.GetSelectMenu(constants.Question_Select_Measure, metricOptions)

// 	UserData := UserData{}

// 	unit := getCorrectUnitSystem(metric)
// 	tdee := CalculateTdee()
// 	goal := goals.GetGoal(tdee)
// 	macros := plan.CalculateMacroNutrients(goal)
// 	var activityOption = menu.GetSelectMenu(constants.Question_Select_Activity, activityOptions)

// 	activityDescription := getActivityExpplanation(activityOption)
// 	println("****************************************************************************")
// 	log.Printf("\n*Age: %.1f Height: %.1f %s Weight: %.1f %s Gender: %s*\n", UserData.Age,
// 		UserData.Height, unit.Height,
// 		UserData.Weight, unit.Weight,
// 		UserData.Gender)
// 	fmt.Printf("*Your TDEE is %.2f*\n", tdee)
// 	fmt.Printf("*Your goal is %.2f*\n", goal)
// 	fmt.Printf("*Based on your TDEE and your goal, your macro distributions are: *\n")
// 	fmt.Printf("\nProtein %.2f: Fats:  %.2f Carbs %.2f\n", macros.Protein, macros.Fats, macros.Carbs)
// 	fmt.Printf("*activityDescription %s: \n", activityDescription)
// 	println("****************************************************************************")
// }
