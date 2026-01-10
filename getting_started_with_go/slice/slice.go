package main
import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	slice := make([]int, 0, 3)

	var userInput string

	for {
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&userInput)
		if userInput == "X" {
			fmt.Print("\nProgramme Terminated...\n\n")
			break
		}
		num, err := strconv.Atoi(userInput)
		if err!=nil {
			continue
		}
		slice = append(slice, num)
		sort.Slice(slice, func(i, j int) bool {return slice[i] < slice[j]})
		fmt.Println(slice)
	}

}
