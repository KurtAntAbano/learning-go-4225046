package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) //+v means both the names and the values of the fields 
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 29
	fmt.Println(poodle)
}

type Dog struct {
	Breed string
	Weight int
}