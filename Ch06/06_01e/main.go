package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./fromString.txt"   // define the file name (relative path)
	file, err := os.Create(fileName) // create a new file (or truncate if it exists)
	defer file.Close()               // defer means wait till everything is executed, then close the file
	checkError(err)                  // check if file creation failed

	// write a string to the file and get the number of bytes written
	length, err := io.WriteString(file, "Hello from Go!")
	fmt.Printf("Wrote a file with %v characters\n", length) // print how many characters were written
	checkError(err)                                         // check if writing failed

	readFile(fileName) // call function to read and print file contents
}

// checkError() panics if an error is encountered
func checkError(err error) {
	if err != nil {
		panic(err) // crash the program and show the error
	}
}

// readFile() reads the file content and prints it
func readFile(fileName string) {
	data, err := os.ReadFile(fileName)                // read entire file content into memory
	checkError(err)                                   // if we get past the error check then move on
	fmt.Println("Text read from file:", string(data)) // convert bytes to string and print
}
