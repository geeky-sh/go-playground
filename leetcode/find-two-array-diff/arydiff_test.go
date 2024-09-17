package findtwoarraydiff

import (
	"fmt"
	"testing"
)

func TestAryDiff(t *testing.T) {
	var tcases = []struct {
		inp1 []int
		inp2 []int
		want [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{[]int{1, 3}, []int{4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{[]int{3}, []int{}}},
	}

	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			out := findDifference(tt.inp1, tt.inp2)
			if len(out) != len(tt.want) {
				t.Errorf("Got %v, Want %v", out, tt.want)
			}
			if len(out) > 0 {
				for i, ay := range out {
					if len(ay) != len(tt.want[i]) {
						t.Errorf("Got %v, Want %v", out, tt.want)
					}
					for j, aay := range ay {
						if aay != tt.want[i][j] {
							t.Errorf("Got %v, Want %v", out, tt.want)
						}
					}
				}
			}
		})
	}
}
