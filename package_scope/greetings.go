package main

import "fmt"

var points = []int{23, 41, 12, 9}

func sayGreeting(n string) {
	fmt.Printf("Hello %v, I hope you're doing well.\nHow's your day going?\n\n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye %v\n", n)
}
