package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	/*
		fmt.Println("Halo, Nick!")

		//strings
		var nameOe string = "Nick"
		var nameTwo = "Lemy"
		var nameThree string
		nameThree = "Name 3"
		nameFour := "yoshi"
		fmt.Printf("%v, %v, %v, %v", nameOe, nameTwo, nameThree, nameFour)

		//ints
		var ageOne int = 20
		var ageTwo = 30
		ageThree := 40

		fmt.Print(ageOne, ageTwo, ageThree)

		// bits & memory
		var numOne int8 = 25
		var numTwo int8 = -128
		var numThree uint16 = 256

		// floats
		var scoreOne float32 = 10.345
		var scoreTwo float64 = 98888882342342.234

		scoreThree := 1.2
	*/

	/*

		age := 20
		name := "Nick"
		// Print
		fmt.Print("Hello, \n")
		fmt.Print("World! \n")

		// Println
		fmt.Println(`He there~ this is Nick`)
		fmt.Println("I'm thrilled to be here")

		// Printf
		fmt.Printf("my age is %v and my name is %v \n", age, name)
		fmt.Printf("my age is %_ and my name is %_ \n", age, name)
		fmt.Printf("my age is %q and my name is %q \n", age, name)
		fmt.Printf("age is of type %T \n", age)
		fmt.Printf("You scored %f \n", 22.23)
		fmt.Printf("You scored %0.2f points \n", 1902.123)

		// Sprintf (save formatted strings)
		var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
		fmt.Println("the saved string is:", str)
	*/

	/*

		var ages [3]int = [3]int{1, 2, 3}
		// var agestwo = [3]int{20, 17, 27}
		names := [4]string{"yoshi", "mario", "peach", "bowser"}

		fmt.Println(names, len(names))
		fmt.Println(ages, len(ages))

		// slices (uses arrays under the hood)
		scores := []int{100, 50, 60}
		scores[0] = 43
		scores = append(scores, 61)
		fmt.Println(scores, len(scores))

		// slices raanges
		rangeOne := names[1:3]
		rangetwo := names[2:]
		rangeThree := names[:3]
		fmt.Println(rangeOne)
		fmt.Println(rangetwo)
		fmt.Println(rangeThree)
	*/

	/*
		// strings package
		greeting := "hello there friends!"
		fmt.Println(strings.Contains(greeting, "hello"))
		fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
		fmt.Println(strings.ToUpper(greeting))
		fmt.Println(strings.Index(greeting, "ll"))
		splittedString := strings.Split(greeting, " ")
		fmt.Println(splittedString)

		// sorting
		ages := []int{23, 45, 12, 67, 34, 89}

		sort.Ints(ages)
		fmt.Println(ages)
		index := sort.SearchInts(ages, 1203)
		fmt.Println(index, len(ages))

		names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

		sort.Strings(names)
		fmt.Println(names)

		index2 := sort.SearchStrings(names, "peach")
		fmt.Println(index2)
	*/

	/*
		// loops
		x := 0
		for x < 5 {
			fmt.Println("value of x is:", x)
			x++
		}

		for i := 0; i < 5; i++ {
			fmt.Println("value of i is:", i)
		}

		for i := range 5 {
			fmt.Println("value of i is:", i)
		}

		names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

		for i := range len(names) {
			fmt.Println(names[i])
		}

		for index, value := range names {
			fmt.Println(index, value)
		}

		for _, value := range names {
			fmt.Println(value)
		}
	*/

	/* Booleans and Conditionals
	age := 92834
	fmt.Println(age <= 50)
	fmt.Println(age >= 10)
	fmt.Println(age == 19)

	if age >= 21 {
		fmt.Println("age is greater than 21")
	} else if age < 21 {
		fmt.Println("age is less than 21")
	} else {
		fmt.Println("age is not 21 or less")
	}

	names := []string{"nick", "christ", "boris", "barthel"}?

	for _, value := range names {
		if value != "boris" {
			fmt.Printf("This is %v\n", value)
			continue
		}
		fmt.Println("This is Boris, please show some Respect")
		break
	}
	*/
	names := []string{"nick", "christ", "boris", "barthel"}

	// Using functions
	sayGreeting("Nick-Lemy")
	sayBye("Christ!")
	cycleNames(names, sayGreeting)
	fmt.Printf("the Area of the circle is: %0.3f\n\n\n", cycleArea(5.25))

	initial1, initial2 := getInitials("Nick Lemy")
	ini1, init2 := getInitials("kayiranga")
	fmt.Println(initial1, initial2)

	fmt.Println(ini1, init2)

}
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func sayGreeting(n string) {
	fmt.Printf("Hello %v, I hope you're doing well.\nHow's your day going?\n\n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r
}
