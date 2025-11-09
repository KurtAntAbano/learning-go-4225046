package main

import (
    "encoding/json"  // for decoding JSON data
    "fmt"            // for printing output
    "io"             // for reading response body
    "net/http"       // for making HTTP requests
    "strconv"        // for converting strings to float
    "strings"        // for working with string readers
)

// API endpoint that returns a JSON array of tour data
const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
    // Fetch the raw JSON content from the API
    content := readHttpContent()

    // Convert the JSON string into a slice of Tour structs
    tours := toursFromJSON(content)

    // Loop through each tour and print its name and price
    for _, tour := range tours {
        // Convert the price string to a float64 (base 16 used here, though base 10 is more typical)
        price, _ := strconv.ParseFloat(tour.Price, 16)
        fmt.Printf("%v: ($%v)\n", tour.Name, price)
    }
}

// toursFromJSON parses a JSON array of tours into a slice of Tour structs
func toursFromJSON(content string) []Tour {
    tours := make([]Tour, 0) // initialize an empty slice of Tour

    // Create a JSON decoder from the string content
    decoder := json.NewDecoder(strings.NewReader(content))

    // Read the opening token (should be the start of the array: '[')
    _, err := decoder.Token()
    checkError(err)

    var tour Tour
    // Loop through each element in the JSON array
    for decoder.More() {
        err := decoder.Decode(&tour) // decode one tour at a time
        checkError(err)
        tours = append(tours, tour) // add the decoded tour to the slice
    }
    return tours
}

// readHttpContent makes a GET request to the API and returns the response body as a string
func readHttpContent() string {
    fmt.Println("Network requests")

    // Create an HTTP client and request
    client := http.Client{}
    req, err := http.NewRequest("GET", url, nil) // nil means no request body
    checkError(err)

    // Set a User-Agent header (can be customized if needed)
    req.Header.Set("User-Agent", "")

    // Send the request and get the response
    resp, err := client.Do(req)
    checkError(err)
    defer resp.Body.Close() // ensure the response body is closed after reading

    fmt.Printf("Response type:%T\n", resp) // print the type of the response object

    // Read the entire response body into a byte slice
    bytes, err := io.ReadAll(resp.Body)
    checkError(err)

    // Convert the byte slice to a string and return it
    return string(bytes)
}

// checkError panics if an error is encountered — simple error handling
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

// Tour defines the structure of each tour item in the JSON response
type Tour struct {
    Name  string // name of the tour
    Price string // price as a string (will be parsed later)
}
