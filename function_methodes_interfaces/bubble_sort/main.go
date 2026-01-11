package main

import (
	"fmt"
)


func main() {
	integers := make([]int, 0)
	var usrInput int
	
	fmt.Println("Program starting")
	for i := range 10 {
		fmt.Printf("Enter Integer (%v/10): ", i + 1)
		fmt.Scan(&usrInput)
		integers = append(integers, usrInput)
	}

	fmt.Println()
	fmt.Println("Unsorted Slice of integers: ")
	for _, v := range integers {
		fmt.Print(v, " ")
	}

	BubbleSort(integers)
	
	fmt.Println()
	fmt.Println()
	fmt.Println("Sorted Slice of integers: ")
	for _, v := range integers {
		fmt.Print(v, " ")
	}
}

func BubbleSort(numbers []int) {
	length := len(numbers)
	for i := 1; i < length; i++ {
		for j := 0; j < (length - 1); j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func Swap(numbers []int, i int){
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}