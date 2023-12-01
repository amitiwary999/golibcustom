package priorityqueue

func parentIndex(pos int) int {
	return pos / 2
}

func leftChildIndex(pos int) int {
	return 2 * pos
}

func rightChildIndex(pos int) int {
	return 2*pos + 1
}

/*
if for the pos the left child position more than the number of items then
it means we have no item for the leaf and this node is leaf
*/
func (h *Heap[T]) isLeaf(pos int) bool {
	return leftChildIndex(pos) > h.size
}

func (h *Heap[T]) swap(a int, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}
