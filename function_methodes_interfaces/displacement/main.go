package main

import (
	"fmt"
	"math"
)

/*
Formula of deplacement in our case
s = (1/2)at^2 + Vot + So
*/

func GenDisplaceFn(a, Vo, So float64) func(t float64) float64 {
	return func (t float64) float64 {
		s := 0.5*a*(math.Pow(t, 2)) + Vo*t + So
		return  s
	}
}
func main(){
	var acc, ivel, idis, t float64
	fmt.Println("Enter acceleration: ")
	fmt.Scan(&acc)

	fmt.Println("Enter initial velocity: ")
	fmt.Scan(&ivel)

	fmt.Println("Enter initial displacement: ")
	fmt.Scan(&idis)

	ComputeDisplace := GenDisplaceFn(acc, ivel, idis)
	fmt.Println()
	fmt.Println("Enter Time: ")
	fmt.Scan(&t)


	fmt.Println()
	fmt.Println("Final result: ", ComputeDisplace(t))
}
