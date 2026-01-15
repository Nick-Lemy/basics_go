package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
    counter++
}

func main() {
    go increment()
    go increment()

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