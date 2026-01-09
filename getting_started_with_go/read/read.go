package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main()  {
	type name struct {
		fname string
		lname string
	}
	fmt.Println("Program Start!...")

	// Read/Create file the file
	var userInput string
	fmt.Println("Enter text file Name (with .txt at the end)")
	fmt.Scan(&userInput)

	data, _ := ioutil.ReadFile(userInput)
	fileContent := string(data)
	arrayOfLines := strings.Split(fileContent, "\n")

	sliceOfNames := make([]name, 0)

	for _, v := range arrayOfLines {
		newNameElts := strings.Split(v, " ")
		newName := name{fname: newNameElts[0], lname: newNameElts[1]}
		sliceOfNames = append(sliceOfNames, newName)
	}

	fmt.Println()
	fmt.Println("Displaying the 'name's in the slice")
	fmt.Println()
	for _, n := range sliceOfNames {
		fmt.Println("First Name: ", n.fname)
		fmt.Println("Last Name: ", n.lname)
		fmt.Println()
	}
}
