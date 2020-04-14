package heap

type Heap2 struct {
	max   int
	array []int
}

func NewHeap2(a []int) Heap2 {
	return Heap2{
		max:   256,
		array: a,
	}
}

func (h *Heap2) lastIndex() int {
	return len(h.array) - 1
}

func (h *Heap2) build() {
	halfIndex := (h.lastIndex() / 2) - 1

	for i := halfIndex; i <= halfIndex*2; i++ {
		h.heapify(halfIndex)
	}
}

func (h *Heap2) heapify(i int) {
	c1 := i*2 + 1
	c2 := i*2 + 2

	if h.array[c1] < h.array[i] || h.array[c2] < h.array[i] {
		if h.array[c2] > h.array[c2] {
			h.array[i], h.array[c2] = h.array[c2], h.array[i]
		} else {
			h.array[i], h.array[c1] = h.array[c1], h.array[i]
		}
	}

	if i > 0 {
		parent := i / 2
		h.heapify(parent)
	}
}

// [5, 1, 4, 2, 3, 9, 7]
// half := 1
//
// 		   5
// 	   1       4
//  2   3    9   7
//
