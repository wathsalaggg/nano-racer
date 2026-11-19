package racer

import (
	"fmt"
	"sort"
)

// Print writes the results as a table, fastest lane first, and marks the
// slowest lane.
func Print(results []Result) {
	if len(results) == 0 {
		return // nothing to print
	}

	// Sorting rearranges a slice. The slice we were given belongs to the
	// caller, and they may still want it in its original order. So we make
	// our own copy and sort that.
	sorted := make([]Result, len(results))
	copy(sorted, results)

	// sort.Slice needs to be told what "comes first" means. We hand it a
	// small function that answers: should item a come before item b?
	// Here: yes, if a took less time.
	sort.Slice(sorted, func(a, b int) bool {
		return sorted[a].Total < sorted[b].Total
	})

	// After sorting, the fastest is first and the slowest is last.
	// Positions in a slice start at 0, so the last one is at len - 1.
	fastest := sorted[0]
	last := len(sorted) - 1

	nameWidth := 0
	for _, r := range sorted {
		if len(r.Name) > nameWidth {
			nameWidth = len(r.Name)
		}
	}

	for i, r := range sorted {
		// Work out the note at the end of the line.
		note := ""
		if i == 0 {
			note = "fastest"
		} else if fastest.Total > 0 { // never divide by zero
			// Dividing two Durations as they are would throw away everything
			// after the decimal point, so 2.9 would become 2. float64 is a
			// number type that keeps decimals.
			times := float64(r.Total) / float64(fastest.Total)
			// Sprintf works like Printf, but gives the text back instead of
			// printing it. %.1f is a decimal number with 1 digit after the
			// point.
			note = fmt.Sprintf("%.1fx slower", times)
		}

		// "&&" means "and": only mark a slowest lane when there is more
		// than one lane in the race.
		if i == last && len(sorted) > 1 {
			note += "  <- slowest"
		}

		line := fmt.Sprintf("%d. %-*s %10v per lap   %s", i+1, nameWidth, r.Name, r.PerLap, note)

		// Colour is another escape code: "\x1b[31m" switches the text to
		// red, and "\x1b[0m" switches back to normal.
		if i == last && len(sorted) > 1 && isTerminal() {
			line = "\x1b[31m" + line + "\x1b[0m"
		}
		fmt.Println(line)
	}
}
