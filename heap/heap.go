package heap

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
		swap(&h.array[parentIdx], &h.array[idx])
		h.Max(parentIdx)
	}
}

func (h *Heap) Poll() int {
	pop := h.array[0]
	swap(&h.array[0], &h.array[h.currentIndex-1])
	h.array[h.currentIndex-1] = 0
	h.currentIndex--
	h.Min(0)
	return pop
}

func (h *Heap) Min(idx int) {
	child1 := (idx * 2) + 1
	child2 := (idx * 2) + 2

	c1, c2 := 0, 0
	if child1 > h.currentIndex {
		c1 = -9999999
	} else {
		c1 = h.array[child1]
	}
	if child2 > h.currentIndex {
		c2 = -9999999
	} else {
		c2 = h.array[child2]
	}

	c := 0
	cIndex := 0
	if c1 < c2 {
		c = c2
		cIndex = child2
	} else {
		c = c1
		cIndex = child1
	}

	if c > h.array[idx] {
		swap(&h.array[cIndex], &h.array[idx])
		h.Min(cIndex)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
