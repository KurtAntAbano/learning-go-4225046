package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100 //elimanating some precision
	fmt.Printf("The sum is now %v\n\n", sum)

	fmt.Println("The value of Pi is", math.Pi)

	circleRadius := 15.5
	cirumcference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", cirumcference) //%.2f\n is for two decimal palces

}
