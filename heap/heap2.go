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
	halfIndex := (h.lastIndex / 2)

	for i := halfIndex; i >= 0; i-- {
		h.heapify(halfIndex)
	}
}

func (h *Heap2) heapify(i int) {
	c1 := i*2 + 1
	c2 := i*2 + 2

	if c2 > h.lastIndex {
		if h.array[c1] < h.array[i] {
			h.arraySwap(i, c1)
		}
	} else if h.array[c1] < h.array[i] || h.array[c2] < h.array[i] {
		if h.array[c1] > h.array[c2] {
			h.arraySwap(i, c2)
		} else {
			h.arraySwap(i, c1)
		}
	}

	// FIXME
	if i > 0 {
		parent := i / 2
		h.heapify(parent)
	}
}

func (h *Heap2) Poll(i int) int {
	h.arraySwap(0, h.lastIndex)
	return 0
}

func (h *Heap2) bubbleDown(i int) {

}

func (h *Heap2) arraySwap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}