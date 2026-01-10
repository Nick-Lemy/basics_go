package main

import "fmt"

func main(){
	var user float32
	fmt.Printf("Welcome to this super nice program that truncate floats\n\n")
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&user)
	result := int(user)
	fmt.Println(result)
}
