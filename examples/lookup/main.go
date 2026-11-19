// Example: is a number in a collection? A slice against a map.
//
//	go run ./examples/lookup
//
// New to Go? Read examples/basic first. It explains the basics that this
// file does not repeat.
package main

import (
	"fmt"

	racer "github.com/wathsalaggg/nano-racer"
)

// How many numbers are in the collection.
const size = 1000

// inSlice looks for target the only way a slice allows: check the first
// item, then the second, and so on until it is found or the slice ends.
//
// "bool" is a true/false value.
func inSlice(numbers []int, target int) bool {
	for _, n := range numbers {
		if n == target { // "==" asks "are these equal?"
			return true
		}
	}
	return false
}

func main() {
	// A slice holding the numbers 0, 1, 2 ... 999.
	// With only one name before ":=", range gives just the position.
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	// A map stores pairs: a key, and the value kept under that key.
	// "map[int]bool" is a map whose keys are ints and whose values are
	// bools. Here every number is a key, and its value is true.
	//
	// What makes a map special is that it does not search. It works out
	// where a key lives directly from the key itself, so looking something
	// up takes about the same time whether the map holds ten keys or ten
	// million.
	inMap := make(map[int]bool)
	for _, n := range numbers {
		inMap[n] = true
	}

	// The last number is the worst case for the slice: it has to walk past
	// all the others first.
	target := size - 1

	// Each lane counts how many times it found the target. Using the answer
	// stops Go from skipping the work as pointless, and lets us check
	// afterwards that both lanes really did the same job.
	sliceFound := 0
	mapFound := 0

	fmt.Printf("Is %d among %d numbers? Asked 2 million times.\n", target, size)

	// The two functions below are written in place, with no name. They can
	// see the variables of main (numbers, inMap, target and the counters).
	// A function that uses variables from around it is called a closure.
	results := racer.New().
		Lane("slice", func() {
			if inSlice(numbers, target) {
				sliceFound++
			}
		}).
		Lane("map", func() {
			// Asking a map for a key looks like using a position in a slice.
			if inMap[target] {
				mapFound++
			}
		}).
		Laps(2_000_000).
		Run()

	fmt.Println()
	racer.Print(results)

	fmt.Println()
	fmt.Printf("slice found it %d times, map found it %d times\n", sliceFound, mapFound)
}
