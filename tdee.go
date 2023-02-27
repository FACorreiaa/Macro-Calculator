package main

type Gender struct {
	slug string
}

var (
	Unknown = Gender{""}
	Guest   = Gender{"male"}
	Member  = Gender{"female"}
)

type TdeeData struct {
	age               int
	gender            Gender
	system            string
	height            int
	weight            int
	formula           string
	routine           string
	daysOfExercise    int
	minutesOfExercise int
	intensity         string
}

// https://steelfitusa.com/blogs/health-and-wellness/calculate-tdee#:~:text=TDEE%20%3D%20BMR%20%2B%20TEF%20%2B%20NEAT,can%20calculate%20your%20individual%20TDEE.
func calculateTDEE() {

}
