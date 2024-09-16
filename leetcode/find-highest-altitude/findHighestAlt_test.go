package findhighestaltitude

import (
	"fmt"
	"testing"
)

func TestFindHightestAlt(t *testing.T) {
	var tcases = []struct {
		inp  []int
		want int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			out := largestAltitude(tt.inp)
			if out != tt.want {
				t.Errorf("Got %v, Want %v\n", out, tt.want)
			}
		})
	}
}
