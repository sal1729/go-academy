package main

import (
	"fmt"
	"go-academy/utils"
	"math/rand"
	"sort"
)

// TODO tests? No branching
func main() {
	// First, create a randomly ordered array containing the numbers 1 to 10
	// We do this by creating a slice
	numberSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(numberSlice), func(i, j int) {
		numberSlice[i], numberSlice[j] = numberSlice[j], numberSlice[i]
	})
	fmt.Println("Numbers: ", numberSlice)
	// Then save the slice to an array
	var numberArray [10]int
	for idx, val := range numberSlice {
		numberArray[idx] = val
	}
	// And tidy up the slice
	numberSlice = nil
	fmt.Println("NumberSlice, empty: ", numberSlice)

	// Display numberArray sorted ascending
	// For the assignment we create a copy of the underlying array, as a slice
	numberSliceAsc := make([]int, len(numberArray))
	copy(numberSliceAsc, numberArray[:])
	fmt.Println("Numbers to sort: ", numberSliceAsc)
	sort.Ints(numberSliceAsc)
	fmt.Println("Numbers, ascending: ", numberSliceAsc)

	// Display numberArray sorted descending
	// We use a new slice, so we can use it later too
	numberSliceDesc := make([]int, len(numberArray))
	copy(numberSliceDesc, numberArray[:])
	fmt.Println("Numbers to sort: ", numberSliceDesc)
	sort.Slice(numberSliceDesc, func(i, j int) bool {
		return numberSliceDesc[i] > numberSliceDesc[j]
	})
	fmt.Println("Numbers, descending: ", numberSliceDesc)

	// Count odds and evens, ascending
	utils.ParityCounter(numberSliceAsc)

	// Count odds and evens, descending
	utils.ParityCounter(numberSliceDesc)
}
