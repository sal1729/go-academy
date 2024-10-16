package main

import (
	"fmt"
	"sync"
)

var x int

func incrementOdd() {
	if (x+1)%2 == 0 {
		x += 2
	} else {
		x++
	}
}

func incrementEven() {
	if (x+1)%2 == 0 {
		x++
	} else {
		x += 2
	}
}

func main() {
	var mu sync.Mutex
	odds := make(chan bool)
	evens := make(chan bool)
	done := make(chan bool)

	// Set up the two goroutines for incrementing
	go func() {
		for {
			<-odds
			mu.Lock()
			incrementOdd()
			fmt.Printf("Odd value: %d\n", x)
			mu.Unlock()
			evens <- true
		}
	}()

	go func() {
		for {
			<-evens
			mu.Lock()
			incrementEven()
			fmt.Printf("Even value: %d\n", x)
			mu.Unlock()
			if x >= 20 {
				done <- true
			} else {
				odds <- true
			}
		}
	}()

	// Trigger by sending signal to odds
	odds <- true

	// Wait for done signal, then tidy up
	<-done
	fmt.Printf("Final value: %d\n", x)
	close(odds)
	close(evens)
	close(done)

}
