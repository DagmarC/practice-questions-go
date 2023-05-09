package reverse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseInt(t *testing.T) {
	var tc = []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2}, []int{2, 1}},
		{arr: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{arr: []int{1}, want: []int{1}},
	}

	for _, test := range tc {
		Reverse(&test.arr)
		if !cmp.Equal(test.arr, test.want) {
			t.Errorf("want %v, got %v", test.want, test.arr)
		}
	}
}
