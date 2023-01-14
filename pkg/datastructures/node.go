package datastructures

type Node[T comparable] map[T]bool

func (n Node[T]) addEdge(key T) {
	n[key] = true
}

func (n Node[T]) edges() []T {
	var keys []T
	for k := range n {
		keys = append(keys, k)
	}

	return keys
}
