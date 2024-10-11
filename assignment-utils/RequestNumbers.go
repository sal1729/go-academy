package utils

import (
	"bufio"
	"fmt"
)

func RequestNumbers(numberOfRequests int, minVal int, maxVal int, reader *bufio.Reader) ([]int, error) {
	request := fmt.Sprintf("Enter a number between %d and %d: ", minVal, maxVal)
	inputs := make([]int, numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		input, _ := RequestInput(request, reader)
		_, number, numCheckError := NumberChecker(input, minVal, maxVal)
		if numCheckError != nil {
			// Note NumberChecker only throws if input is not a number, not if input is outside of range
			return nil, numCheckError
		}
		inputs[i] = number
	}
	return inputs, nil
}
