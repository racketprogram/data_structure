package heap

import (
)

type Heap struct {
	maxSize      int
	currentIndex int
	array        []int
}

func NewHeap(heapSize int) Heap {
	return Heap{
		maxSize:      heapSize,
		currentIndex: 0,
		array:        make([]int, heapSize, heapSize),
	}
}

func (h *Heap) Insert(v int) {
	if h.currentIndex >= h.maxSize {
		return
	}

	h.array[h.currentIndex] = v
	h.Max(h.currentIndex)
	h.currentIndex++
}

func (h *Heap) Max(idx int) {
	if idx == 0 {
		return
	}
	parentIdx := (idx / 2)
	if h.array[parentIdx] < h.array[idx] {
		tmp := h.array[parentIdx]
		h.array[parentIdx] = h.array[idx]
		h.array[idx] = tmp
		h.Max(parentIdx)
	}
}
