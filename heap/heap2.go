package heap

import "fmt"

func init() {
	fmt.Println("heap2 init")
	fmt.Println(9 / 2)
}

type Heap2 struct {
	max   int
	array []int
	lastIndex int
}

func NewHeap2(a []int) Heap2 {
	return Heap2{
		max:   256,
		array: a,
		lastIndex: len(a) - 1,
	}
}

func (h *Heap2) build() {
	halfIndex := (h.lastIndex / 2) - 1

	for i := halfIndex; i >= 0; i-- {
		h.BubbleDown(i)
	}
}

func (h *Heap2) Poll(i int) int {
	h.arraySwap(0, h.lastIndex)
	return 0
}

func (h *Heap2) BubbleDown(idx int) {
	c1 := (idx * 2) + 1
	c2 := (idx * 2) + 2

	var c, cIndex int
	if c1 > h.lastIndex {
		return
	} else if c2 > h.lastIndex {
		c = h.array[c1]
		cIndex = c1
	} else {
		if h.array[c1] > h.array[c2] {
			c = h.array[c2]
			cIndex = c2
		} else {
			c = h.array[c1]
			cIndex = c1
		}
	}

	if c < h.array[idx] {
		h.arraySwap(cIndex, idx)
		h.BubbleDown(cIndex)
	}
}

func (h *Heap2) arraySwap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *Heap2) ValidateHeap() bool {
	return h.validateHeap(0)
}

func (h *Heap2) validateHeap(i int) bool {
	c1 := i * 2 + 1
	c2 := i * 2 + 2

	return h.compare(i, c1) && h.compare(i, c2)
}

func (h *Heap2) validateMinHeap() bool {
	for i := h.lastIndex; i > 0; i-- {
		value := h.array[i]
		parent := (i-1) / 2
		parentValue := h.array[parent]
		if parentValue > value {
			return false
		}
	}
	return true
}

func validateMinHeapf(a []int) bool {
	for i := len(a) - 1; i > 0; i-- {
		value := a[i]
		parent := (i-1) / 2
		parentValue := a[parent]
		if parentValue > value {
			return false
		}
	}
	return true
}

func (h *Heap2) compare(i, j int) bool {
	if j > h.lastIndex {
		return true
	}

	if h.array[i] > h.array[j] {
		return false
	}

	return h.validateHeap(j)
}
