// Package racer times Go functions and races them against each other.
//
// Every Go file starts with a "package" line. All files in the same folder
// share one package name. Other programs use this code by importing it.
package racer

// "import" brings in code from other packages.
// "time" is part of Go's standard library, so nothing has to be installed.
import "time"

// Time runs fn once and returns how long it took.
//
// Reading the first line of the function, piece by piece:
//
//	func            starts a function
//	Time            its name. The capital T makes it public ("exported"),
//	                so code outside this package is allowed to call it.
//	fn func()       it takes one input named fn. The type of fn is func(),
//	                meaning "a function with no inputs and no outputs".
//	                In Go, functions can be passed around like any other value.
//	time.Duration   what it gives back: a length of time, stored in nanoseconds.
func Time(fn func()) time.Duration {
	// ":=" creates a new variable and gives it a value in one step.
	start := time.Now() // look at the clock

	fn() // run the function we were given

	// time.Since(start) is "the clock now, minus start".
	return time.Since(start)
}
