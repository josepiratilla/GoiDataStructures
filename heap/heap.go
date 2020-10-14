package heap

// Heap represent a sorted heap to implement a priority queue
// when pushin it will add an element to the queue, and sort it.
// when poping it will get the highest element.
type Heap struct {
	data []int
}

// Push adds an element to the heap
func (h *Heap) Push(i int) error {
	h.data = append(h.data, i)
	n := h.last()
	if h.data[n] > h.data[parent(n)] {
		h.swap(n, parent(n))
	}
	return nil
}

// Pop gets and removes the highest element in the heap
func (h *Heap) Pop() (int, error) {
	out := h.data[0]
	h.swap(0, h.last())
	return out, nil
}

func (h *Heap) last() int {
	return len(h.data) - 1
}

func (h *Heap) swap(a, b int) {
	temp := h.data[b]
	h.data[b] = h.data[a]
	h.data[a] = temp
}

func parent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}
