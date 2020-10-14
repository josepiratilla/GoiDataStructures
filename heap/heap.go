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
	h.orderUpFromLast()

	return nil
}

// Pop gets and removes the highest element in the heap
func (h *Heap) Pop() (int, error) {
	out := h.data[0]
	h.swap(0, h.last())
	h.data = h.data[:h.last()]
	h.orderDown(0)
	return out, nil
}

func (h *Heap) last() int {
	return len(h.data) - 1
}

func (h *Heap) orderUpFromLast() {
	h.orderUp(h.last())
}
func (h *Heap) orderUp(i int) {
	a := parent(i)
	b := i
	for h.order(a, b) {
		if a == 0 {
			break
		}
		b = a
		a = parent(a)
	}
}

func (h *Heap) order(a, b int) bool {
	if h.data[a] < h.data[b] {
		h.swap(a, b)
		return true
	}
	return false
}

func (h *Heap) orderDown(i int) {
	change := true

	for change {
		i, change = h.orderThree(i)
	}
}

func (h *Heap) orderThree(a int) (int, bool) {
	l := h.last()
	b := leftChild(a)
	c := rightChild(a)

	max := a
	if b <= l {
		if h.data[b] > h.data[max] {
			max = b
		}
	}
	if c <= l {
		if h.data[c] > h.data[max] {
			max = c
		}
	}
	if max > a {
		h.swap(a, max)
		return max, true
	}
	return a, false

}

func (h *Heap) swap(a, b int) {
	temp := h.data[b]
	h.data[b] = h.data[a]
	h.data[a] = temp
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2

}
