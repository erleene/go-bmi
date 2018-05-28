package calc

import (
	"fmt"
	"log"
	"strconv"
)

//method to calculate the height and weight and BMI

//it has to be a type func()
//so i need a method to return a func()
//this func should take in parameters and return a func

func Calculate(meters string, kg string) func() {

	return func() {
		var bmiValue int

		heightVal, err := strconv.Atoi(meters)
		if err != nil {
			log.Print(err)
		}
		weightVal, err := strconv.Atoi(kg)
		if err != nil {
			log.Print(err)
		}
		//formula to return the calculation
		//kg/m2  - kilograms and meters
		bmiValue = weightVal / (heightVal * heightVal)
		fmt.Printf("Your bmi is: %v", bmiValue)
	}

}
