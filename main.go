package main

import (
	"fmt"
	heap "golibcustom/priorityqueue"
)

func main() {
	func() {
		h := heap.NewHeap(20, func(a, b *int) bool { return b == nil || a != nil && *a < *b })
		h.Push(12)
		h.Push(23)
		h.Push(17)
		h.Push(15)
		h.Push(18)

		pI, err := h.Peek()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("peek item %d \n", pI)

		po, err := h.Pop()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("pop item %d \n", po)

		pinext, err := h.Peek()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("peek item %d \n", pinext)
	}()
}
