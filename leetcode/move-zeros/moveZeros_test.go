package movezeros

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	var tcases = []struct {
		inp  []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			moveZeroes(tt.inp)
			for ii := range tt.inp {
				if tt.inp[ii] != tt.want[ii] {
					t.Errorf("Incorrect: Got %v, Want %v", tt.inp, tt.want)
				}
			}
		})
	}
}
