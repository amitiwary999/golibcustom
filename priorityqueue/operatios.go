package priorityqueue

import "fmt"

func (h *Heap[T]) Peek() (res T, err error) {
	if h.size < 1 {
		return res, fmt.Errorf("queue is empty")
	}
	res = h.data[0]
	return res, nil
}

func (h *Heap[T]) Push(item T) error {
	if h.size >= len(h.data) {
		return fmt.Errorf("container is full, can't add more data")
	}
	cur := h.size
	h.data[cur] = item
	for h.comperator(&h.data[cur], &h.data[parentIndex(cur)]) {
		h.swap(cur, parentIndex(cur))
		cur = parentIndex(cur)
	}
	h.size++
	return nil
}

func (h *Heap[T]) Pop() (res T, err error) {
	if h.size < 1 {
		return res, fmt.Errorf("queue is empty")
	}
	res = h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.fixQueueItemOrder(0)
	return res, nil
}

func (h *Heap[T]) fixQueueItemOrder(pos int) {
	if h.isLeaf(pos) {
		return
	}
	var leftIndex = leftChildIndex(pos)
	var rightIndex = rightChildIndex(pos)
	var cur *T = &h.data[pos]
	var left *T = &h.data[leftIndex]
	var right *T
	if rightIndex < h.size {
		right = &h.data[rightIndex]
	}

	if h.comperator(left, cur) || h.comperator(right, cur) {
		if h.comperator(left, right) {
			h.swap(pos, leftIndex)
			h.fixQueueItemOrder(leftIndex)
		} else {
			h.swap(pos, rightIndex)
			h.fixQueueItemOrder(rightIndex)
		}
	}
}
