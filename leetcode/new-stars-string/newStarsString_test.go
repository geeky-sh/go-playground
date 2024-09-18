package newstarsstring

import (
	"fmt"
	"testing"
)

func TestNewStarsString(t *testing.T) {
	var tcases = []struct {
		inp  string
		want string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}

	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			out := removeStars(tt.inp)
			if out != tt.want {
				t.Errorf("Got %v Want %v\n", out, tt.want)
			}
		})
	}
}
