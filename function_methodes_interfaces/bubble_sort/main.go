package main

import (
	"fmt"
)


func main() {
	integers := make([]int, 0)
	var usrInput int
	ran := []int{1, 2, 3, 4, -1, 0}
	BubbleSort(ran)
	fmt.Println("Blaaaaaaa: ", ran)
	
	fmt.Println("Program starting")
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter Integer (%v/10): ", i + 1)
		fmt.Scan(&usrInput)
		integers = append(integers, usrInput)
	}
	fmt.Println("final Slice of integers: ", integers)
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