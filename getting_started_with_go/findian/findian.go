package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	fmt.Printf("Welcome to the \"ian\" Finder Program\n\n")
	fmt.Print("Enter a string: ")
	fmt.Scan(&userInput)
	
	userInput = strings.ToLower(userInput)

	startsWithI := strings.HasPrefix(userInput, "i")
	endsWithN := strings.HasSuffix(userInput, "n")
	containsA := strings.Contains(userInput, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else{
		fmt.Println("Not Found!")
	}
}