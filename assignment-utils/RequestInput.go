package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func RequestInput(request string, reader *bufio.Reader) (string, error) {
	fmt.Printf(request)

	input, readErr := reader.ReadString('\n')
	if readErr != nil {
		return "", readErr
	}

	return strings.TrimSpace(input), nil
}

func RequestInputs(requests []string, reader *bufio.Reader) ([]string, error) {
	var inputs []string

	for _, request := range requests {
		input, readErr := RequestInput(request, reader)
		if readErr != nil {
			return nil, readErr
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}
