package main

import (
	"fmt"
)

func main() {

	var p *int //dereference operator, will point to a variabl of the type
	if p == nil {
		fmt.Println("p is nill")
	} else {
		fmt.Println("value of p:", p)
	}

}
