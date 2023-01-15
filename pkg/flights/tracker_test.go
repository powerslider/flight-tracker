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
		input    [][]flights.Airport
		expected string
	}{
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "A"},
			},
			expected: "A -> B -> A",
		},
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "B"},
			},
			expected: "B -> C -> B",
		},
		{
			input: [][]flights.Airport{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			expected: "A -> B -> C -> A",
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "SFO"},
			},
			expected: "SFO -> EWR -> SFO",
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "EWR"},
			},
			expected: "EWR -> ATL -> EWR",
		},
		{
			input: [][]flights.Airport{
				{"SFO", "EWR"},
				{"EWR", "ATL"},
				{"ATL", "SFO"},
			},
			expected: "SFO -> EWR -> ATL -> SFO",
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

		if !strings.Contains(err.Error(), test.expected) {
			t.Errorf("Error does not print cycle: %q", err)
		}
	}
}
