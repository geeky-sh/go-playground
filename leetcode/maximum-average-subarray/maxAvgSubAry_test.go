package maximumaveragesubarray

import (
	"fmt"
	"testing"
)

func TestMaxArgSubAry(t *testing.T) {
	var tcases = []struct {
		inp  []int
		inp2 int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00000},
	}

	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			out := findMaxAverage(tt.inp, tt.inp2)
			if out != tt.want {
				t.Errorf("Got %v Want %v\n", out, tt.want)
			}
		})
	}
}
