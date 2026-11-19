package racer

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

// "const" is a value that never changes.
// barWidth is how many characters wide a full progress bar is.
const barWidth = 30

// watch draws the progress bars again and again until every lane is done.
//
// progress holds one lap counter per lane. The lanes add to the counters,
// and this function only reads them.
func (r *Race) watch(progress []atomic.Int64) {
	// The bars work by moving the cursor around the screen. If the output is
	// going into a file instead, that would fill the file with junk, so we
	// draw nothing at all.
	if !isTerminal() {
		return
	}

	// Print one empty line per lane. This makes room for the bars, so the
	// first drawing has something to draw over.
	fmt.Print(strings.Repeat("\n", len(r.lanes)))

	// "for" with nothing after it loops forever, until a return or break.
	for {
		stillRunning := r.drawBars(progress)
		if !stillRunning {
			return // that was the final drawing, with every bar full
		}
		// Pause before drawing again: 20 drawings per second looks smooth
		// and costs almost nothing.
		time.Sleep(50 * time.Millisecond)
	}
}

// drawBars draws one bar per lane, on top of the previous drawing.
// It reports whether any lane still has laps left to run.
func (r *Race) drawBars(progress []atomic.Int64) bool {
	// Find the longest lane name, so all the bars start in the same column.
	nameWidth := 0
	for _, l := range r.lanes {
		if len(l.name) > nameWidth {
			nameWidth = len(l.name)
		}
	}

	// Terminals obey special text called escape codes. They start with the
	// invisible "escape" character, written \x1b in Go.
	// "\x1b[3A" means "move the cursor up 3 lines". We move up to the first
	// bar so the new bars replace the old ones. This is what makes them
	// look animated.
	fmt.Printf("\x1b[%dA", len(r.lanes))

	stillRunning := false
	for i, l := range r.lanes {
		// Load reads the counter safely while a lane may be changing it.
		// It gives an int64, and int(...) converts that to a plain int.
		done := int(progress[i].Load())
		if done < r.laps {
			stillRunning = true
		}

		// How much of the bar to fill. Dividing two ints in Go throws away
		// the remainder, so this is always a whole number of characters.
		filled := done * barWidth / r.laps
		bar := strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled)

		// %-*s  text padded on the right; the * takes the width from the
		//       next value (nameWidth) instead of a fixed number
		// %3d   a whole number, at least 3 characters wide
		// %%    a real percent sign
		fmt.Printf("%-*s [%s] %3d%%\n", nameWidth, l.name, bar, done*100/r.laps)
	}
	return stillRunning
}

// isTerminal reports whether the program's output is a real terminal
// window, as opposed to a file or another program.
func isTerminal() bool {
	// Many Go functions return two things: a result and an error.
	// The error is nil (meaning "nothing") when all went well.
	info, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	// A terminal is what the operating system calls a "character device".
	// The & picks that one yes/no flag out of the file's mode.
	return info.Mode()&os.ModeCharDevice != 0
}
