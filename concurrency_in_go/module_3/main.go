package main

import (
	"fmt"
	"slices"
	"sync"
)

// Sorts a slice in place
func sliceSorter(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting subarray:", arr)
	slices.Sort(arr)
}

func merge(a, b []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main() {
	array := make([]int, 0, 12)
	var userInput int

	fmt.Println("Program Starting...")
	fmt.Println("Enter 12 integers in one line separated by spaces:")
	fmt.Print("> ")

	for i := 0; i < 12; i++ {
		fmt.Scan(&userInput)
		array = append(array, userInput)
	}

	// Partition array
	arr1 := array[:3]
	arr2 := array[3:6]
	arr3 := array[6:9]
	arr4 := array[9:]

	var wg sync.WaitGroup
	wg.Add(4)

	go sliceSorter(arr1, &wg)
	go sliceSorter(arr2, &wg)
	go sliceSorter(arr3, &wg)
	go sliceSorter(arr4, &wg)

	wg.Wait()

	// Merge sorted subarrays
	merged1 := merge(arr1, arr2)
	merged2 := merge(arr3, arr4)
	final := merge(merged1, merged2)

	fmt.Println("Final sorted array:", final)
}
