package racer

// Several imports go inside one pair of brackets.
import (
	"sync"        // standard library: tools for code that runs at the same time
	"sync/atomic" // numbers that are safe to share between goroutines
	"time"
)

// A "struct" groups several values under one name.
//
// lane is one competitor in a race: a label, and the function to run.
// The lowercase first letter makes it private to this package. People using
// the library never touch a lane directly, so they do not need to see it.
type lane struct {
	name string
	fn   func()
}

// Result is what one lane achieved. It starts with a capital letter because
// people using the library need to read it.
type Result struct {
	Name   string
	Total  time.Duration // time for all the laps together
	PerLap time.Duration // average time for one lap
}

// Race holds everything about one race: who is in it and how long it is.
type Race struct {
	// "[]lane" is a slice of lanes. A slice is a list that can grow.
	lanes []lane
	laps  int
}

// New creates an empty race.
//
// It returns "*Race", a pointer to a Race. A pointer is the address of a
// value rather than a copy of it. Everyone holding the pointer sees the same
// Race, so adding a lane through it changes the real thing.
// "&Race{...}" means "make a Race and give me its address".
func New() *Race {
	return &Race{laps: 1000} // 1000 laps unless Laps is called
}

// Lane adds a competitor to the race.
//
// The "(r *Race)" before the name makes this a method: a function that
// belongs to Race. It is called with a dot, like race.Lane("name", fn).
//
// It returns the same race again. That is what lets calls be chained:
//
//	racer.New().Lane("a", fnA).Lane("b", fnB).Run()
func (r *Race) Lane(name string, fn func()) *Race {
	// append adds one item to the end of a slice and returns the new slice.
	r.lanes = append(r.lanes, lane{name: name, fn: fn})
	return r
}

// Laps sets how many times each function is run.
//
// One run of a fast function is over too quickly to measure well, so we
// run it many times and divide.
func (r *Race) Laps(n int) *Race {
	if n < 1 {
		n = 1 // zero laps would mean dividing by zero later
	}
	r.laps = n
	return r
}

// Run starts every lane at the same moment and waits for all of them to
// finish. It returns one Result per lane, in the order the lanes were added.
func (r *Race) Run() []Result {
	// make creates a slice with a fixed number of empty slots, one per lane.
	// Each lane writes into its own slot, so no two lanes ever touch the
	// same piece of memory. That is what makes this safe.
	results := make([]Result, len(r.lanes))

	// A WaitGroup is a counter of jobs still running.
	// "var" declares a variable without giving it a value; a WaitGroup is
	// ready to use as it is.
	var wg sync.WaitGroup

	// One lap counter per lane, so the progress bars know how far each lane
	// has got. The lanes write these numbers while the bar-drawing code
	// reads them. Two goroutines using one ordinary variable at the same
	// time is a bug called a data race. atomic.Int64 is a number that is
	// safe to share.
	progress := make([]atomic.Int64, len(r.lanes))

	// "range" walks over a slice. Each time round it gives the position (i)
	// and the item (l).
	for i, l := range r.lanes {
		wg.Add(1) // one more job is running

		// "go" starts a goroutine: the function after it begins running in
		// the background, and this loop carries on immediately without
		// waiting. Goroutines are how Go does many things at once.
		go func() {
			// "defer" means "do this when the function ends, however it
			// ends". Here: tell the WaitGroup this job is finished.
			defer wg.Done()

			// "&" gives the address of this lane's counter, so runLane
			// updates the real counter and not a copy of it.
			results[i] = r.runLane(l, &progress[i])
		}() // the () calls the function we just wrote
	}

	// Draw the progress bars while the lanes run. This returns once every
	// lane has completed its last lap.
	r.watch(progress)

	// Without this line Run could return before a lane had saved its
	// result. Wait blocks until the counter is back to zero.
	wg.Wait()
	return results
}

// runLane runs one lane for the full number of laps and times it.
// Lowercase name, so it is a private helper.
//
// done is this lane's lap counter. It goes up by one after every lap.
func (r *Race) runLane(l lane, done *atomic.Int64) Result {
	start := time.Now()
	for i := 0; i < r.laps; i++ {
		l.fn()
		// This costs a few nanoseconds per lap, and it is counted in the
		// lane's time. For very tiny functions that is noticeable.
		done.Add(1)
	}
	total := time.Since(start)

	return Result{
		Name:  l.name,
		Total: total,
		// A Duration is a whole number of nanoseconds, so it can be divided.
		// r.laps is a plain int, and Go never mixes number types silently:
		// time.Duration(r.laps) converts it first.
		PerLap: total / time.Duration(r.laps),
	}
}
