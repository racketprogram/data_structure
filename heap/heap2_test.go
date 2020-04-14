package heap

import (
	"fmt"
	"testing"
)

func TestNewHeap2(t *testing.T) {
	s := []int{2, 10, 5, 1, 7, 9, 3, 4, 6, 8}
	heap := NewHeap2(s)
	fmt.Printf("TestNewHeap2: %+v\n", heap)
	heap.build()
	fmt.Printf("TestNewHeap2: %+v\n", heap)

	//        1
	//      2    5
	//   4     7  9  3
	// 10 6  8
}
