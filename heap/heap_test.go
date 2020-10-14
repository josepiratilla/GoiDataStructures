package heap

import (
	"math/rand"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	vs := []struct {
		push        []int
		expectedPop []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{2, 1},
			[]int{2, 1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{1, 2, 3, 7, 0, -1, 6},
			[]int{7, 6, 3, 2, 1, 0, -1},
		},
		{
			[]int{1, 2, 3, 7, 0, -1, 6, 2, 2},
			[]int{7, 6, 3, 2, 2, 2, 1, 0, -1},
		},
	}
	for _, v := range vs {
		h := new(Heap)
		for _, push := range v.push {
			err := h.Push(push)
			if err != nil {
				t.Errorf("Unexpected error when pushing %v\n", v.push)
			}
		}
		for i, expectedPop := range v.expectedPop {
			actualPop, err := h.Pop()
			if err != nil {
				t.Errorf("Unexpected error when pop %v\n", v.push)
			}
			if actualPop != expectedPop {
				t.Errorf("%dth value of %v.\nExpected Pop: %d\nActual Pop: %d\n", i+1, v.push, expectedPop, actualPop)
			}
		}
	}

}

func TestSize(t *testing.T) {
	vs := []struct {
		push     []int
		expected int
	}{
		{
			[]int{1},
			1,
		},
	}
	for _, v := range vs {
		h := new(Heap)
		for _, push := range v.push {
			err := h.Push(push)
			if err != nil {
				t.Errorf("Unexpected error when pushing %v\n", v.push)
			}
		}
		actual := h.Size()
		if actual != v.expected {
			t.Errorf("With Input %v\nExpected: %d\nActual: %d\n", v.push, v.expected, actual)
		}
	}
}

func TestEmptyError(t *testing.T) {
	rand.Seed(1)
	for i := 0; i < 100; i++ {
		h := new(Heap)
		queueSize := 1 + rand.Intn(100)
		for j := 0; j < queueSize; j++ {
			h.Push(rand.Intn(10000) - 5000)
		}
		for j := 0; j < queueSize; j++ {
			_, err := h.Pop()
			if err != nil {
				t.Errorf("Unexpected error when pop the %dth element after push %d elements\n", j, queueSize)
			}

		}
		_, err := h.Pop()
		if err != errEmptyQueue {
			t.Errorf("Add %d elements and removed %d elements without the Empty Queue error", queueSize, queueSize+1)
		}

	}
}
