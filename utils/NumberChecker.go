package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func NumberChecker(input string, minVal int, maxVal int) (string, int, error) {
	integer := regexp.MustCompile(`^-?\d+$`)
	if integer.MatchString(input) {
		number, stringConvErr := strconv.Atoi(input)
		if stringConvErr != nil {
			return "", number, stringConvErr
		}

		if number < minVal {
			return fmt.Sprintf("Your number is less than %d", minVal), number, nil
		}
		if number > maxVal {
			return fmt.Sprintf("Your number is greater than %d", maxVal), number, nil
		}
		return fmt.Sprint("Your number is valid"), number, nil
	}

	// TODO we're being wildly inconsistent about where printing out happens
	// Do it all inside or all outside, but not a random mixture of both!
	// TODO What should we return?
	return "", 0, errors.New("number must be an integer, in numeric form")
}
