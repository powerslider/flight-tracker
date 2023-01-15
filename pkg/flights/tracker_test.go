package flights_test

import (
	"github.com/powerslider/flight-tracker/pkg/flights"
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

func TestFlightTracker(t *testing.T) {
	tests := []struct {
		input    [][]flights.Airport
		expected []flights.Airport
	}{
		{
			input: [][]flights.Airport{
				{"A", "B"},
			},
			expected: []flights.Airport{"A", "B"},
		},
		{
			input: [][]flights.Airport{
				{"C", "B"},
				{"A", "C"},
			},
			expected: []flights.Airport{"A", "B"},
		},
		{
			input: [][]flights.Airport{
				{"D", "B"},
				{"A", "C"},
				{"E", "D"},
				{"C", "E"},
			},
			expected: []flights.Airport{"A", "B"},
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
			},
			expected: []flights.Airport{"SFO", "EWR"},
		},
		{
			input: [][]flights.Airport{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
			expected: []flights.Airport{"SFO", "EWR"},
		},
		{
			input: [][]flights.Airport{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			expected: []flights.Airport{"SFO", "EWR"},
		},
	}

	for _, test := range tests {
		tracker := flights.NewTracker()

		for _, row := range test.input {
			tracker.AddFlightRoute(row[0], row[1])
		}

		actual, err := tracker.Trace()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, actual, test.expected)
	}
}

func TestFlightTrackerCyclicDependencies(t *testing.T) {
	tests := []struct {
		input                  [][]flights.Airport
		expectedPossibleCycles []string
	}{
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "A"},
			},
			expectedPossibleCycles: []string{
				"A -> B -> A",
				"B -> A -> B",
			},
		},
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "B"},
			},
			expectedPossibleCycles: []string{
				"B -> C -> B",
				"C -> B -> C",
				"B -> A -> B",
			},
		},
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			expectedPossibleCycles: []string{
				"A -> B -> C -> A",
				"B -> C -> A -> B",
				"C -> A -> B -> C",
			},
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "SFO"},
			},
			expectedPossibleCycles: []string{
				"SFO -> EWR -> SFO",
				"EWR -> SFO -> EWR",
			},
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "EWR"},
			},
			expectedPossibleCycles: []string{
				"EWR -> ATL -> EWR",
				"EWR -> SFO -> EWR",
				"ATL -> EWR -> ATL",
			},
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "SFO"},
			},
			expectedPossibleCycles: []string{
				"SFO -> EWR -> ATL -> SFO",
				"EWR -> ATL -> SFO -> EWR",
				"ATL -> SFO -> EWR -> ATL",
			},
		},
	}

	for _, test := range tests {
		tracker := flights.NewTracker()

		for _, row := range test.input {
			tracker.AddFlightRoute(row[0], row[1])
		}

		_, err := tracker.Trace()
		if err == nil {
			t.Error("Expected cycle error did not occur!")
		}

		if !errorContainsAnyOf(test.expectedPossibleCycles, err) {
			t.Errorf("Error does not print cycle: %q", err)
		}
	}
}

func errorContainsAnyOf(s []string, err error) bool {
	for _, v := range s {
		if strings.Contains(err.Error(), v) {
			return true
		}
	}

	return false
}
