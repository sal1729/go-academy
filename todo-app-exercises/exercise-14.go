package main

import "fmt"

var n int

func incrementToNextOdd() {
	if (n+1)%2 == 0 {
		n += 2
	} else {
		n++
	}
}

func incrementToNextEven() {
	if (n+1)%2 == 0 {
		n++
	} else {
		n += 2
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go incrementToNextOdd()
		go incrementToNextEven()
		fmt.Printf("value: %d\n", n)
	}
}
