package main

import (
	"fmt"
	"time"
)

func main() {
	go say("HEY goroutine") // go = using go routines 
	// if u run this u'll notice the bottom line executes first
	// this is because we added a delay to the say funciton
	fmt.Println("Goroutines")

	//anomynous function

	go func(message string){
		fmt.Println(message)
	} ("Hello from the anoymnous function")
	


	time.Sleep(2*time.Second)
	fmt.Println("ALL DONE")

}

func say(message string) {
	time.Sleep(1*time.Second)
	fmt.Println(message)
}
