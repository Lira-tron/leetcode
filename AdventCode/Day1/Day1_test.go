package leetcode

import (
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Example 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve()
		})
	}
}
