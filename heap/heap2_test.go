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
	fmt.Printf("TestValidateHeap: %+v\n", heap.validateMinHeap())
	//           1
	//      2       5
	//   4     7   9  3
	// 10 6  8
}

func TestValidateHeap(t *testing.T) {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	heap := NewHeap2(s)
	fmt.Printf("TestValidateHeap: %+v\n", heap)
	fmt.Printf("TestValidateHeap: %+v\n", heap.validateMinHeap())
	//        1
	//      2    5
	//   4     7  9  3
	// 10 6  8
}

func TestHeapQuestions(t *testing.T) {
	s := [][]int{
		{5, 1, 4, 2, 3},
		{4, 4, 3, 1, 9, 2},
		{100},
		{3,3,3,3},
	}

	for _, v := range s {
		heap := NewHeap2(v)
		fmt.Printf("TestNewHeap2: %+v\n", heap)
		heap.build()
		fmt.Printf("TestNewHeap2: %+v\n", heap)
		fmt.Printf("TestValidateHeap: %+v\n", heap.validateMinHeap())
	}
}

func BenchmarkValidateHeap(b *testing.B) {
	s := []int{1,2,3,4,5,6,7,8,9,10}
    for i := 0; i < b.N; i++ {
		validateMinHeapf(s)
    }
}
