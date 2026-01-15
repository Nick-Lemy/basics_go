package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	fmt.Println("Printint increment 1")
    counter++
}
func increment2() {
	fmt.Println("Printing increment 2")
	counter++
}

func main() {
    go increment()
    go increment2()

    time.Sleep(1 * time.Second)
    fmt.Println("Counter:", counter)
}

/*
Race condition explanation:

Both goroutines try to update counter at the same time. Since counter++ is not atomic,
they might read, increment, 
and write at the same time, causing one increment to be lost. 
The result can be 1 instead of 2, depending on timing.
*/