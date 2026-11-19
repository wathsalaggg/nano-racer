// Example: three ways to sort the same 500 numbers.
//
//	go run ./examples/sorting
//
// New to Go? Read examples/basic first. It explains the basics that this
// file does not repeat.
package main

import (
	"fmt"
	"math/rand" // standard library: random numbers
	"sort"      // standard library: sorting

	racer "github.com/wathsalaggg/nano-racer"
)

// bubbleSort is the simplest sort to describe. Walk along the slice and
// swap any two neighbours that are in the wrong order. Do that again and
// again until a whole walk needs no swaps.
//
// It changes the slice it is given. A slice passed to a function is not
// copied; the function works on the caller's numbers.
func bubbleSort(nums []int) {
	// "for" followed by a condition keeps looping while it is true.
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				// Go can swap two values in one line, with no helper
				// variable.
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
	}
}

// insertionSort works the way many people sort a hand of playing cards.
// Take the next card and slide it left until it sits in the right place
// among the cards already sorted.
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		// "&&" means "and". Keep sliding left while there is room to the
		// left and the neighbour on the left is bigger.
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func main() {
	// 500 numbers in a shuffled order.
	//
	// A random generator started from the same "seed" (here 1) produces
	// the same numbers every time. That keeps the race fair from one run
	// to the next.
	// Perm(500) gives the numbers 0 to 499 in a random order.
	shuffled := rand.New(rand.NewSource(1)).Perm(500)

	// lap builds the function for one lane.
	//
	// Every lap has to start from the same shuffled numbers, or the second
	// lap would be sorting numbers that are already sorted. So each lap
	// first copies the shuffled numbers onto the lane's own desk, then
	// sorts the copy.
	//
	// Each lane gets its own desk because the lanes run at the same time,
	// and two lanes sorting the same slice would wreck each other's work.
	//
	// This is a function that takes a function and returns a function.
	lap := func(sortIt func([]int)) func() {
		desk := make([]int, len(shuffled))
		return func() {
			copy(desk, shuffled)
			sortIt(desk)
		}
	}

	fmt.Println("Sorting 500 shuffled numbers, 3000 times each.")

	results := racer.New().
		Lane("bubble sort", lap(bubbleSort)).
		Lane("insertion sort", lap(insertionSort)).
		Lane("sort.Ints", lap(sort.Ints)). // the one that comes with Go
		Laps(3000).
		Run()

	fmt.Println()
	racer.Print(results)
}
