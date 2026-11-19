package racer

import "time"

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

// Run runs every lane, one after another, and returns one Result per lane,
// in the order the lanes were added.
func (r *Race) Run() []Result {
	// "var" declares a variable without giving it a value yet.
	// An empty slice is ready to be appended to.
	var results []Result

	// "range" walks over a slice. Each time round it gives the position and
	// the item. We do not need the position here, so it goes to "_".
	for _, l := range r.lanes {
		results = append(results, r.runLane(l))
	}
	return results
}

// runLane runs one lane for the full number of laps and times it.
// Lowercase name, so it is a private helper.
func (r *Race) runLane(l lane) Result {
	start := time.Now()
	for i := 0; i < r.laps; i++ {
		l.fn()
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
