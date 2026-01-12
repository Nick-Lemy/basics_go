package main

import (
	"fmt"
	"os"
)

type Animal struct {name, food, locomotion, noise string}

func (n *Animal) Eat() {
	fmt.Print(n.food)
}

func (n *Animal) Move(){
	fmt.Print(n.locomotion)
}

func (n *Animal) Speak(){
	fmt.Print(n.noise)
}

func main()  {
	var prompt1, prompt2 string
	var selectedAnimal Animal
	cow := Animal { 
		name: "cow", 
		food: "grass", 
		locomotion: "walk", 
		noise: "moo",
	}
	bird := Animal { 
		name: "bird", 
		food: "worms", 
		locomotion: "fly", 
		noise: "peep",
	}
	snake := Animal { 
		name: "snake", 
		food: "mice", 
		locomotion: "slither", 
		noise: "hsss",
	}

	fmt.Println("Starting program...")
	fmt.Println()
	println("Enter animal name then info name (in 1 line, eg: cow move)")
	fmt.Print(">")
	fmt.Scan(&prompt1)
	fmt.Scan(&prompt2)

	fmt.Println()
	switch prompt1 {
	case "cow":
		selectedAnimal = cow
	case "bird":
		selectedAnimal = bird
	case "snake":
		selectedAnimal = snake
	default:
		fmt.Println("Program Terminated...")
		os.Exit(0)
	}

	switch prompt2 {
	case "eat":
		fmt.Printf("%v eats: ", prompt1)
		selectedAnimal.Eat()
	case "move":
		fmt.Printf("%v's locomotion method is: ", prompt1)
		selectedAnimal.Move()
	case "speak":
		fmt.Printf("%v's sound is: ", prompt1)
		selectedAnimal.Speak()
	default:
		fmt.Println("Program Terminated...")
		os.Exit(0)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Program Terminated...")
}