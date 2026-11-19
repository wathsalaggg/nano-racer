package racer

// A file whose name ends in _test.go holds tests. Tests are small programs
// that check the real code does what it should. Go only reads these files
// when you run:
//
//	go test
import "testing"

// A test is a function whose name starts with Test and which takes
// "t *testing.T". Go finds and runs every such function by itself.
// t is how the test reports a problem.
func TestRunRunsEveryLapOfEveryLane(t *testing.T) {
	// Each lane counts how many times it was called.
	// A function written inside another function can use and change the
	// variables around it, which is how these two reach the counters.
	aCalls := 0
	bCalls := 0

	results := New().
		Lane("a", func() { aCalls++ }).
		Lane("b", func() { bCalls++ }).
		Laps(50).
		Run()

	// "!=" means "is not equal to".
	if aCalls != 50 {
		// Errorf marks the test as failed and says why. The usual style is
		// to say what we got, then what we wanted.
		t.Errorf("lane a ran %d times, want 50", aCalls)
	}
	if bCalls != 50 {
		t.Errorf("lane b ran %d times, want 50", bCalls)
	}

	// Fatalf also fails the test, and stops it right here. That matters
	// because the lines below would crash if there were fewer than
	// 2 results.
	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}
	if results[0].Name != "a" || results[1].Name != "b" { // "||" means "or"
		t.Errorf("results are in the wrong order: %q then %q", results[0].Name, results[1].Name)
	}
}

func TestLapsBelowOneBecomesOne(t *testing.T) {
	calls := 0

	New().Lane("only", func() { calls++ }).Laps(0).Run()

	if calls != 1 {
		t.Errorf("function ran %d times, want 1", calls)
	}
}
