package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	fmt.Println("Program Starts")
	var inputName string
	var inputAddress string
	fmt.Println()

	fmt.Print("Enter your Name: ")
	fmt.Scan(&inputName)

	fmt.Println()

	fmt.Print("Enter your Address: ")
	fmt.Scan(&inputAddress)

	userMap := map[string]string{"name": inputName, "address": inputAddress}
	fmt.Println()

	finalJson, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println("Error Parsing the map")
	} else {
		fmt.Println(string(finalJson))
	}
}