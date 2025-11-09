package main

import (
    "fmt"
    "io"
    "net/http"
)

// URL of the API endpoint we're requesting data from
const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
    fmt.Println("Network requests") // log that we're starting a network request

    // Create an HTTP client (used to send the request)
    client := http.Client{}

    // Build a new GET request to the specified URL
    req, err := http.NewRequest("GET", url, nil) // nil means no request body
    checkError(err) // check if request creation failed

    // Set a custom User-Agent header (empty here, but can be used to identify the client)
    req.Header.Set("User-Agent", "")

    // Send the request and receive the response
    resp, err := client.Do(req)
    checkError(err) // check if the request failed

    // Ensure the response body is closed after we're done reading it
    defer resp.Body.Close()

    // Print the type of the response object (for debugging or inspection)
    fmt.Printf("Response type:%T\n", resp)

    // Read the entire response body into a byte slice
    bytes, err := io.ReadAll(resp.Body)
    checkError(err) // check if reading the response failed

    // Convert the byte slice to a string (the actual content)
    content := string(bytes)

    // Print the content received from the API
    fmt.Print(content)
}

// checkError() panics if an error is encountered — used for quick error handling
func checkError(err error) {
    if err != nil {
        panic(err) // crash the program and show the error message
    }
}
