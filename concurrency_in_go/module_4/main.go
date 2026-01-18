package main

import (
	"fmt"
	"sync"
)

const numPhilosophers = 5
const eatCount = 3

type Chopstick struct{ sync.Mutex }

func philosopher(id int, left, right *Chopstick, host chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < eatCount; i++ {
		// Request permission from host
		host <- struct{}{} // blocks if host limit reached

		// Pick up chopsticks (any order)
		left.Lock()
		right.Lock()

		fmt.Printf("starting to eat %d\n", id)
		fmt.Printf("finishing eating %d\n", id)

		// Put down chopsticks
		left.Unlock()
		right.Unlock()

		// Release host permission
		<-host
	}
}

func main() {
	// Initialize chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Host channel allows max 2 philosophers to eat concurrently
	host := make(chan struct{}, 2)

	var wg sync.WaitGroup
	wg.Add(numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		left := chopsticks[i]
		right := chopsticks[(i+1)%numPhilosophers]

		// Start philosopher goroutine
		go philosopher(i+1, left, right, host, &wg)
	}

	wg.Wait()
}
