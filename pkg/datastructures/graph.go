package datastructures

import (
	"fmt"
	"strings"
)

type Graph[T comparable] struct {
	adjacencyList map[T]Node[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjacencyList: make(map[T]Node[T]),
	}
}

func (g *Graph[T]) NodeKeys() []T {
	keys := make([]T, 0, len(g.adjacencyList))
	for k := range g.adjacencyList {
		keys = append(keys, k)
	}

	return keys
}

func (g *Graph[T]) AddNodeIfAbsent(node T) Node[T] {
	// Check if node is already added.
	n, ok := g.adjacencyList[node]

	// If not add the new node
	if !ok {
		n = make(Node[T])
		g.adjacencyList[node] = n
	}

	// Return existing or newly added node.
	return n
}

func (g *Graph[T]) AddEdge(from T, to T) {
	// Add from node.
	fromNode := g.AddNodeIfAbsent(from)
	// Add to node.
	_ = g.AddNodeIfAbsent(to)

	// Connect from and to adjacencyList via an edge.
	fromNode.addEdge(to)
}

func (g *Graph[T]) ContainsNode(key T) bool {
	_, ok := g.adjacencyList[key]

	return ok
}

func (g *Graph[T]) ResolveDependencies(key T) ([]T, error) {
	results := NewOrderedSet[T]()

	// Run topological sorting algorithm to achieve dependency resolution.
	err := g.topSort(key, results, nil)
	if err != nil {
		return nil, err
	}

	return results.items, nil
}

func (g *Graph[T]) LongestPath() ([]T, error) {
	longestPath := make([]T, 0)

	for _, k := range g.NodeKeys() {
		currentPath, err := g.ResolveDependencies(k)
		if err != nil {
			return nil, err
		}

		if len(currentPath) > len(longestPath) {
			longestPath = currentPath
		}
	}

	return longestPath, nil
}

func (g *Graph[T]) topSort(key T, results *OrderedSet[T], visited *OrderedSet[T]) error {
	if visited == nil {
		visited = NewOrderedSet[T]()
	}

	added := visited.add(key)
	if !added {
		return detectCyclicDependency(key, visited)
	}

	n := g.adjacencyList[key]
	for _, edge := range n.edges() {
		err := g.topSort(edge, results, visited.copy())
		if err != nil {
			return err
		}
	}

	results.add(key)

	return nil
}

func detectCyclicDependency[T comparable](key T, visited *OrderedSet[T]) error {
	index := visited.indexOf(key)

	cycle := append(visited.items[index:], key)
	cycleOutput := make([]string, len(cycle))

	for i, k := range cycle {
		cycleOutput[i] = fmt.Sprintf("%v", k)
	}

	return fmt.Errorf("cyclic dependency error: %s", strings.Join(cycleOutput, " -> "))
}
