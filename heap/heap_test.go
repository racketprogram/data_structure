package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	fmt.Printf("TestNewHeapInsert: %+v\n", heap)
}

func TestNewHeapPoll(t *testing.T) {
	s := []int{2, 10, 5, 1, 7, 9, 3, 4, 6, 8}
	heap := NewHeap(len(s))
	for _, v := range s {
		heap.Insert(v)
	}
	fmt.Printf("TestNewHeapPoll %+v\n", heap)
	var sorted []int
	for i := 0; i < 10; i++ {
		sorted = append(sorted, heap.Poll())
		fmt.Printf("TestNewHeapPoll %+v\n", heap)
	}
	fmt.Printf("TestNewHeapPoll sorted %+v\n", sorted)
	assert.Equal(t, sorted, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}
