package utils

import "fmt"

func ParityCounter(numbers []int) {
	var odds, evens []int
	oddCount, evenCount := 0, 0
	for _, number := range numbers {
		if number%2 == 0 {
			evens = append(evens, number)
			evenCount++
		} else {
			odds = append(odds, number)
			oddCount++
		}
		fmt.Printf("Processing value=%d, oddCount=%d, evenCount=%d\n", number, oddCount, evenCount)
	}

	fmt.Printf("Number of Odds: %d, Number of Evens: %d\n", oddCount, evenCount)
}
