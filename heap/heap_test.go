package heap

import "testing"

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
