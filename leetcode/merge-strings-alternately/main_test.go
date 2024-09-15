package main

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	var tcases = []struct {
		inp1 string
		inp2 string
		want string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
		{"", "", ""},
		{"ab", "", "ab"},
	}

	for i, tt := range tcases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			got := mergeAlternately(tt.inp1, tt.inp2)
			if got != tt.want {
				t.Errorf("Got %s Want %s", got, tt.want)
			}
		})
	}
}
