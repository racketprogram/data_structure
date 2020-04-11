package heap

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	NewHeap(10)
}

func TestNewHeapInsert(t *testing.T) {
	s := []int{2, 10, 5, 1, 7, 9, 3, 4, 6, 8}
	heap := NewHeap(len(s))

	for _, v := range s {
		heap.Insert(v)
	}
	
	fmt.Printf("%+v\n", heap)
}
