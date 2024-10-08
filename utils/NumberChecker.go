package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func NumberChecker(input string, minVal int, maxVal int) (string, error) {
	integer := regexp.MustCompile(`^-?\d+$`)
	if integer.MatchString(input) {
		number, stringConvErr := strconv.Atoi(input)
		if stringConvErr != nil {
			return "", stringConvErr
		}

		if number < minVal {
			return fmt.Sprintf("Your number is less than %d", minVal), nil
		}
		if number > maxVal {
			return fmt.Sprintf("Your number is greater than %d", maxVal), nil
		}
		return fmt.Sprint("Your number is valid"), nil
	}

	return "", errors.New("number must be an integer, in numeric form")
}
