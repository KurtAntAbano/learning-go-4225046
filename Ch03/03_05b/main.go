package main

import (
	"fmt"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington" // like a dictionary
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println("Maps")

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)
}
