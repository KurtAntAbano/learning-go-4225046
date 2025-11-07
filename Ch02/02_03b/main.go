package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ =	reader.ReadString('\n')
	//no := as we've already declared these to variables
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) //gets rid of spaces
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Value of number:", f)
	}
}
