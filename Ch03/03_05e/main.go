package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)
	
	california := states["CA"]
	fmt.Println(california)
	
	delete(states, "OR")
	fmt.Println(states)
	
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states { // k = key, v = value
		fmt.Printf("%v: %v\n", k, v)  // %v placeholder
	}

	keys := make([]string, len(states)) // exracting key by using slice
	i := 0
	for k := range states {  // adds the keys into the slice
		keys[i] = k
		i++
	}
	sort.Strings(keys)  // now that we have our key slice, sort makes it alphabetical
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])  // we use our sorted key slice as the index to output the key values (map)
	}
}
