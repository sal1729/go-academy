package utils

import (
	"errors"
)

type Units []int
type Tens []int
type Hundreds []int

func isInRange(queryValue int, minVal int, maxVal int) bool {
	return queryValue >= minVal && queryValue <= maxVal
}

func sumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func NewUnits(numbers []int) (Units, error) {
	units := make(Units, len(numbers))
	for idx, number := range numbers {
		if isInRange(number, 0, 9) {
			units[idx] = number
		} else {
			return nil, errors.New("input to NewUnits must be numbers between 0 and 9")
		}
	}
	return units, nil
}
func (units Units) Sum() int {
	return sumArray(units)
}

func NewTens(numbers []int) (Tens, error) {
	tens := make(Tens, len(numbers))
	for idx, number := range numbers {
		if isInRange(number, 10, 99) {
			tens[idx] = number
		} else {
			return nil, errors.New("input to NewTens must be numbers between 10 and 99")
		}
	}
	return tens, nil
}
func (tens Tens) Sum() int {
	return sumArray(tens)
}

func NewHundreds(numbers []int) (Hundreds, error) {
	hundreds := make(Hundreds, len(numbers))
	for idx, number := range numbers {
		if isInRange(number, 100, 999) {
			hundreds[idx] = number
		} else {
			return nil, errors.New("input to NewHundreds must be numbers between 100 and 999")
		}
	}
	return hundreds, nil
}
func (hundreds Hundreds) Sum() int {
	return sumArray(hundreds)
}
