package main

import (
	"fmt"
	"slices"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	CheckName(string) bool
}

// Cow Struct
type Cow struct {
	name, food, locomotion, sound string
}

func (snake Cow) Eat(){
	fmt.Printf("%v eats: %v", snake.name, snake.food)
	fmt.Println()
}
func (snake Cow) Move(){
	fmt.Printf("The locomotion method of %v is called: %v", snake.name, snake.locomotion)
	fmt.Println()
}
func (snake Cow) Speak(){
	fmt.Printf("The Sound made by %v is called: %v", snake.name, snake.sound)
	fmt.Println()
}
func (snake Cow) CheckName(name string) bool{
	return strings.Compare(snake.name, name) == 0
}

// Bird Struct
type Bird struct {
	name, food, locomotion, sound string
}

func (bird Bird) Eat(){
	fmt.Printf("%v eats: %v", bird.name, bird.food)
	fmt.Println()
}
func (bird Bird) Move(){
	fmt.Printf("The locomotion method of %v is called: %v", bird.name, bird.locomotion)
	fmt.Println()
}
func (bird Bird) Speak(){
	fmt.Printf("The Sound made by %v is called: %v", bird.name, bird.sound)
	fmt.Println()
}
func (bird Bird) CheckName(name string) bool{
	return strings.Compare(bird.name, name) == 0
}

// Snake Struct
type Snake struct {
	name, food, locomotion, sound string
}

func (snake Snake) Eat(){
	fmt.Printf("%v eats: %v", snake.name, snake.food)
	fmt.Println()
}
func (snake Snake) Move(){
	fmt.Printf("The locomotion method of %v is called: %v", snake.name, snake.locomotion)
	fmt.Println()
}
func (snake Snake) Speak(){
	fmt.Printf("The Sound made by %v is called: %v", snake.name, snake.sound)
	fmt.Println()
}
func (snake Snake) CheckName(name string) bool{
	return strings.Compare(snake.name, name) == 0
}



var animals []Animal = []Animal{}
func main(){
	var userInput string
	var allInputs []string
	
	fmt.Println()
	fmt.Println("Program starting...")
	fmt.Print("> ")
	for len(allInputs) < 3 {
		fmt.Scan(&userInput)
		if strings.Compare(userInput, "quit") == 0 {
			fmt.Println("Program Terminated...")
			return
		}
		
		allInputs = append(allInputs, userInput)
	}
	if strings.Compare(allInputs[0], "newanimal") != 0 && strings.Compare(allInputs[0], "query") != 0 {
		main()
	}
	switch allInputs[0] {
	case "newanimal":
		switch allInputs[2] {
		case "cow":
			animals = append(animals, Cow{
				name: allInputs[1],
				food: "grass",
				locomotion: "walk",
				sound: "moo",
			})

		case "bird":
			animals = append(animals, Bird{
				name: allInputs[1],
				food: "worms",
				locomotion: "fly",
				sound: "peep",
			})
		case "snake":
			fmt.Println("okkkk...")
			animals = append(animals, Snake{
				name: allInputs[1],
				food: "mice",
				locomotion: "slither",
				sound: "hsss",
			})
		default:
			fmt.Println("Wrong input, try again...")
			main()
			
		}
		fmt.Println("Created it!")
	case "query":
		fmt.Println(animals)
		if len(animals) == 0{
			fmt.Println("No animals created yet... Try again")
			main()
		}
		idx := slices.IndexFunc(animals, func(animal Animal) bool {
			return animal.CheckName(allInputs[1])
		})
		if idx == -1 {
			fmt.Println("This animal does not exist... Try again")
			main()
		}
		
		selectedAnimal := animals[idx]
		switch allInputs[2] {
		case "eat":
			selectedAnimal.Eat()
			fmt.Println()
		case "move":
			selectedAnimal.Move()
			fmt.Println()
		case "speak":
			selectedAnimal.Speak()
			fmt.Println()
		default:
			fmt.Println("Wrong input, try again...")
			main()
		}
	}
	fmt.Println(animals[len(animals)-1])
	// fmt.Println(allInputs)
	main()
}
