// Example: two ways to work out a Fibonacci number.
//
//	go run ./examples/fibonacci
//
// The Fibonacci numbers are 0, 1, 1, 2, 3, 5, 8, 13, 21 ...
// Each one is the two before it added together.
//
// New to Go? Read examples/basic first. It explains the basics that this
// file does not repeat.
package main

import (
	"fmt"

	racer "github.com/wathsalaggg/nano-racer"
)

// fibRecursive follows the definition word for word: the answer is the two
// earlier answers added together.
//
// A function that calls itself is called recursive. It reads nicely, but
// look at the cost: each call makes two more calls, and each of those makes
// two more. Working out number 20 this way takes 21,891 calls, most of them
// repeating work that was already done somewhere else.
func fibRecursive(n int) int {
	if n < 2 {
		return n // the first two numbers are 0 and 1; this stops the calls
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// fibLoop walks up the list once, remembering only the last two numbers.
// Number 20 takes 20 small steps.
func fibLoop(n int) int {
	// Two variables can be created in one line.
	a, b := 0, 1
	for i := 0; i < n; i++ {
		// Both sides are worked out before anything is stored, so this
		// moves one step along the list: a becomes b, b becomes a+b.
		a, b = b, a+b
	}
	return a
}

func main() {
	const which = 20

	// Each lane keeps its last answer. Using the answer stops Go from
	// skipping the work as pointless, and lets us check afterwards that
	// both lanes agree.
	recursiveAnswer := 0
	loopAnswer := 0

	fmt.Printf("Working out Fibonacci number %d, 20,000 times each.\n", which)

	results := racer.New().
		Lane("recursive", func() { recursiveAnswer = fibRecursive(which) }).
		Lane("loop", func() { loopAnswer = fibLoop(which) }).
		Laps(20_000).
		Run()

	fmt.Println()
	racer.Print(results)

	fmt.Println()
	fmt.Printf("recursive says %d, loop says %d\n", recursiveAnswer, loopAnswer)
}
