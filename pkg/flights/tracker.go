package flights

import (
	"sort"
)

type history struct {
	ledger *datastructures.Graph[Airport]
}

func newFlightHistory() *history {
	return &history{
		ledger: datastructures.NewGraph[Airport](),
	}
}

func (h *history) addFlightRoute(start Airport, end Airport) {
	h.ledger.AddEdge(start, end)
}

func (h *history) getLongestFlightPath() ([]Airport, error) {
	return h.ledger.LongestPath()
}

type Tracker struct {
	flightHistory *history
}

func NewTracker() *Tracker {
	return &Tracker{
		flightHistory: newFlightHistory(),
	}
}

func (t *Tracker) AddFlightRoute(start Airport, end Airport) {
	t.flightHistory.addFlightRoute(start, end)
}

func (t *Tracker) Trace() ([]Airport, error) {
	longestPath, err := t.flightHistory.getLongestFlightPath()
	if err != nil {
		return nil, err
	}

	fromAirport := longestPath[len(longestPath)-1]
	toAirport := longestPath[0]

	return []Airport{
		fromAirport,
		toAirport,
	}, nil
}

func reverse[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
