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
	return nil
}

// Pop gets and removes the highest element in the heap
func (h *Heap) Pop() (int, error) {
	return h.data[0], nil
}
