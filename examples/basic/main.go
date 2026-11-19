// This is a small program that uses the racer library.
//
// "package main" is special: it means "this is a program you can run",
// not a library. Run it from the top folder of the project with:
//
//	go run ./examples/basic
package main

import (
	"fmt"     // standard library: printing text
	"strings" // standard library: working with text

	// Our own library. The word "racer" in front is the name we use for it
	// in this file.
	racer "github.com/wathsalaggg/nano-racer"
)

// Two ways to build the same piece of text: the letter "x", 500 times.
// They give the same answer. The race shows which one gets there faster.

// withPlus glues the text together with +=.
// Every += builds a brand new string and copies the old one into it,
// so the work grows as the string gets longer.
func withPlus() {
	s := ""
	// A "for" loop: start i at 0, keep going while i < 500, add 1 each time.
	for i := 0; i < 500; i++ {
		s += "x"
	}
	// Go refuses to compile a variable that is never used.
	// Assigning it to "_" tells Go we are ignoring it on purpose.
	_ = s
}

// withBuilder uses strings.Builder, a tool made for this job.
// It keeps spare room at the end, so adding a letter rarely needs a copy.
func withBuilder() {
	var b strings.Builder
	for i := 0; i < 500; i++ {
		b.WriteString("x")
	}
	_ = b.String()
}

// Every runnable Go program starts at func main().
func main() {
	// Each line ends with a dot so the chain can continue on the next line.
	results := racer.New().
		Lane("plus", withPlus).
		Lane("builder", withBuilder).
		Laps(20_000). // underscores in numbers are only for readability
		Run()

	fmt.Println() // an empty line between the bars and the table
	racer.Print(results)
}
