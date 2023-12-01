package priorityqueue

type Heap[T any] struct {
	data       []T
	size       int
	comperator Comparator[T]
}

type Comparator[T any] func(a *T, b *T) bool

func NewHeap[T any](capacity int, comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{
		comperator: comparator,
		size:       0,
		data:       make([]T, capacity, capacity),
	}
}
