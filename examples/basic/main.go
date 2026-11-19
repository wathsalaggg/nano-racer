// This is a small program that uses the racer library.
//
// "package main" is special: it means "this is a program you can run",
// not a library. Run it from the top folder of the project with:
//
//	go run ./examples/basic
package main

import (
	"fmt" // standard library: printing text

	// Our own library. The word "racer" in front is the name we use for it
	// in this file.
	racer "github.com/wathsalaggg/nano-racer"
)

// countTo counts from 0 up to n and does nothing else.
// It exists only to give us something slow enough to time.
func countTo(n int) {
	total := 0
	// A "for" loop: start i at 0, keep going while i < n, add 1 each time.
	for i := 0; i < n; i++ {
		total += i
	}
	// Go refuses to compile a variable that is never used.
	// Assigning it to "_" tells Go we are ignoring it on purpose.
	_ = total
}

// Every runnable Go program starts at func main().
func main() {
	// Time wants a function with no inputs, but countTo needs a number.
	// So we wrap the call inside a small function with no name,
	// written right here: func() { ... }
	took := racer.Time(func() {
		countTo(10_000_000) // underscores in numbers are only for readability
	})

	fmt.Println("counting to 10 million took", took)
}
