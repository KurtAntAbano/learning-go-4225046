package main

import (
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something")
	value1 := 5
	value2 := 10
	value3 := 56
	sum, count, average := addAllValues(value1, value2, value3, 125)
	fmt.Printf("The sum is %v.\n", sum)
	fmt.Printf("The count is %v.\n", count)
	fmt.Printf("The average is %v.\n", average)

}

func addValues(value1, value2 int) int { //last int is the return data type
	return value1 + value2
}

func addAllValues(values ...int) (int, int, float64) {  // slice of integers (SO ITS NOT CONSTRAINED TO SPECIFIC LENGTH, can be 2 numbers, can be 10)
	sum := 0
	for _, v := range values { // values is a slice
		sum += v
	}
	count := len(values)
	average := float64(sum) / float64(count)
	return sum, count, average // IMPORTANT: look ho it mirrors the (int, int, float64) at the top
}
