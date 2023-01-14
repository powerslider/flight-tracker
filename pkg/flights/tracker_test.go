package flights_test

import (
	"fmt"
	"github.com/powerslider/flight-tracker/pkg/flights"
	"testing"
)

func TestTopSort(t *testing.T) {
	tracker := flights.NewTracker()

	// a -> b -> c
	tracker.AddFlightRoute("a", "b")
	tracker.AddFlightRoute("b", "c")

	results, err := tracker.Trace()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(results)
	//if results[0] != "c" || results[1] != "b" || results[2] != "a" {
	//	t.Errorf("Wrong sort order: %v", results)
	//}
}

func TestTopSort2(t *testing.T) {
	tracker := flights.NewTracker()

	// a -> b -> c
	tracker.AddFlightRoute("a", "b")
	tracker.AddFlightRoute("c", "b")
	tracker.AddFlightRoute("a", "c")

	results, err := tracker.Trace()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(results)
}

func TestTopSort3(t *testing.T) {
	tracker := flights.NewTracker()

	// a -> b -> c
	tracker.AddFlightRoute("d", "b")
	tracker.AddFlightRoute("a", "c")
	tracker.AddFlightRoute("e", "d")
	tracker.AddFlightRoute("c", "e")

	results, err := tracker.Trace()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(results)
}
