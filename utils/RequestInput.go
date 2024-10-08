package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StringReader interface {
	ReadString()
}

func RequestInput(request string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(request)

	input, readErr := reader.ReadString('\n')
	if readErr != nil {
		return "", readErr
	}

	return strings.TrimSpace(input), nil
}

func RequestInputs(requests []string) ([]string, error) {
	var inputs []string

	for _, request := range requests {
		input, readErr := RequestInput(request)
		if readErr != nil {
			return nil, readErr
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}
